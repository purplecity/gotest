package Snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	nodeBits  uint8 = 10
	stepBits  uint8 = 12
	nodeMax   int64 = -1 ^ (-1 << nodeBits) // -1 用二进制表示为1的补码   1023
	stepMax   int64 = -1 ^ (-1 << stepBits)
	timeShift uint8 = nodeBits + stepBits
	nodeShift uint8 = stepBits
)

var Epoch int64 = 1288834974657
var nodeStartMutex = sync.Mutex{}
var Nodestart = int64(0)

type ID int64

type Node struct {
	mu        sync.Mutex //互斥锁
	timestamp int64      //时间戳
	node      int64      //节点ID
	step      int64      //序列号ID
}

func NewNode(node int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}
	return &Node{timestamp: 0, node: node, step: 0}, nil
}

func (n *Node) Generate() ID {
	n.mu.Lock()
	defer n.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if n.timestamp == now {
		n.step++

		if n.step > stepMax {
			for now > n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		n.step = 0
	}
	n.timestamp = now
	result := ID((now-Epoch)<<timeShift | (n.node << nodeShift) | (n.step))
	return result
}

func GenID() string {
	nodeStartMutex.Lock()
	defer nodeStartMutex.Unlock()
	if Nodestart == 1023 {
		Nodestart = 1023 - Nodestart
	} else {
		Nodestart += 1
	}
	node, _ := NewNode(Nodestart)
	id := node.Generate()
	return fmt.Sprint(id)
}