package contexttest

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

var Canceled = errors.New("context canceled")

type deadlineExceededError struct {}

func (deadlineExceededError) Error() string {return "context deadline exceeded"}
func (deadlineExceededError) Timeout() bool {return true}
func (deadlineExceededError) Temporary() bool {return true}

var DeadlineExceeded = deadlineExceededError{}

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknow emtpy Context"
}

var (
	background = new(emptyCtx)
	todo = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}



/******************************/

var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type canceler interface {
	cancel(removeFromParent bool,err error)
	Done() <- chan struct{}
}

type cancelCtx struct {
	Context
	mu sync.Mutex
	done chan struct{}
	children map[canceler]struct{}
	err error
}

func (c *cancelCtx) Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done =make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func (c *cancelCtx) String() string {
	return fmt.Sprintf("%v.WithCancel",c.Context)
}

func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context:internal error:missing cancel error")
	}
	c.mu.Lock()

	if c.err != nil {
		c.mu.Unlock()
		return //already canceled
	}

	c.err = err

	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}

	for child := range c.children {
		child.cancel(false,err)
	}

	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		removeChild(c.Context,c)
	}
 }


type timerCtx struct {
	cancelCtx
	timer *time.Timer
	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *timerCtx) String() string {
	return fmt.Sprintf("%v.WithDeadline(%s [%s])",c.cancelCtx.Context,c.deadline,time.Until(c.deadline))
}

func (c *timerCtx) cancel(removeFromParent bool,err error) {
	c.cancelCtx.cancel(false,err)
	if removeFromParent {
		removeChild(c.cancelCtx.Context,c)
	}

	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}

type valueCtx struct {
	Context
	key,val interface{}
}

func (c *valueCtx) String() string {
	return fmt.Sprintf("%v.WithValue(%#v,%#v)",c.Context,c.key,c.val)
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}


/****************************************/
//所以传的时候是指针类型 也就是3个ctx内部的context是指针类型
func parentCancelCtx(parent Context) (*cancelCtx,bool) {
	for {
		switch c := parent.(type) {
		case *cancelCtx:
			return c,true
		case *timerCtx:
			return &c.cancelCtx,true
		case *valueCtx:
			parent = c.Context
		default:
			return nil,false
		}
	}
}

func removeChild(parent Context, child canceler) {
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children,child)
	}
	p.mu.Unlock()
}

//parent挂的时候自己也挂 自己挂的如果想挂了自己的child 可以去调用cancel方法 但是不移除child 如果想移除自己在parent中的就去cancel的参数为true
//把自己放到parent的map中
//仅适用与cancel和deadline和timeout
func propagateCancel(parent Context, child canceler) {
	if parent.Done() == nil {
		return //parent is never canceled
	}

	if p, ok := parentCancelCtx(parent);ok {
		p.mu.Lock()
		if p.err != nil {
			//parent has already been canceled
			child.cancel(false,p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false,parent.Err())
			case  <-child.Done():
			}
		}()
	}
}

/*********************************/

type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancle CancelFunc) {
	c := cancelCtx{Context:parent}
	propagateCancel(parent,&c)
	return &c, func() {c.cancel(true,Canceled)}
}

func WithDeadline(parent Context, d time.Time) (Context,CancelFunc) {
	if cur,ok := parent.Deadline();ok && cur.Before(d) {
		return WithCancel(parent)
	}

	c :=&timerCtx{
		cancelCtx:cancelCtx{Context:parent},
		deadline:d,
	}

	propagateCancel(parent,c)
	dur := time.Until(d)
	if dur <= 0 {
		c.cancel(true,DeadlineExceeded)
		return c,func() {c.cancel(false,Canceled)}
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err == nil {
		c.timer = time.AfterFunc(dur,func() {
			c.cancel(true,DeadlineExceeded)
		})
	}
	return c,func() {c.cancel(true,Canceled)}
}

func WithValue(parent Context,key,val interface{}) Context {
	if key == nil {
		panic("nil key")
	}

	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	return &valueCtx{parent,key,val}
}