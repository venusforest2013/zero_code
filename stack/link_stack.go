package stack

import "fmt"

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

func (z *LinkStack) Pop() int {
	if z.size == 0 {
		return 0
	}

	ret := z.top.data
	z.top = z.top.prev
	z.size--

	return ret
}

func (z *LinkStack) PrintStack() {
	cur := z.top
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.prev
	}
}
