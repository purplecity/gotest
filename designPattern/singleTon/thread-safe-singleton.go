package singleTon

import "sync"

//a counter that holds the number of times it has been called during program execution


// ---------------------------------------mux----------------------------------------
type singleton struct {
	count int
	mux sync.RWMutex
}

/*
var instance singleton

func GetInstance() *singleton {
	return &instance
*/

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.count++
	return  s.count
}

func (s *singleton) GetCount() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.count
}

// ---------------------------------------chan----------------------------------------
var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)
var count int

func init() {
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

var newInstance singleton

func GetNewInstance() *singleton {
	return &newInstance
}

//3个触发器
func (s *singleton) NewAddOne() {
	addCh <- true
}

func (s *singleton) NewGetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}