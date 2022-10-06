package stack

import (
	"errors"
	"fmt"
)

type StackNode struct {
	data int
	prev *StackNode
}

type LinkStack struct {
	top  *StackNode
	size int
}

func (z *LinkStack) Push(element int) {
	node := &StackNode{
		data: element,
	}
	if z.size == 0 {
		z.top = node
	} else {
		node.prev = z.top
		z.top = node
	}
	z.size++
}

func (z *LinkStack) Pop() (int, error) {
	if z.size == 0 {
		return 0, errors.New("stack empty")
	}

	ret := z.top.data
	z.top = z.top.prev
	z.size--

	return ret, nil
}

func (z *LinkStack) PrintStack() {
	cur := z.top
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.prev
	}
}
func (z *LinkStack) IsEmpty() bool {
	return z.size == 0

}
func (z *LinkStack) Peek() (int, error) {
	if z.size == 0 {
		return 0, errors.New("stack empty")
	}
	return z.top.data, nil

}
