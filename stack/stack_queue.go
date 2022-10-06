package stack

import (
	"fmt"
	"time"
)

//用栈实现队列，add,poll,peak

func StackQueueExample() {
	sq := NewStackQueue()
	sq.Add(1)
	sq.Add(2)
	sq.Add(3)
	sq.Add(4)
	sq.Add(5)
	sq.Add(6)
	sq.pushStack.PrintStack()
	fmt.Println("push stack print finish")
	sq.popStack.PrintStack()
	fmt.Println("pop stack print finish")
	time.Sleep(time.Second)
	ret, err := sq.Poll()
	fmt.Println(ret)
	fmt.Println(err)
	fmt.Println("=====")
	sq.pushStack.PrintStack()
	fmt.Println("push stack print finish=====")
	sq.popStack.PrintStack()
	fmt.Println("pop stack print finish")
}

type StackQueue struct {
	pushStack *SliceStack
	popStack  *LinkStack
}

func NewStackQueue() *StackQueue {
	return &StackQueue{
		pushStack: NewSliceStack(),
		popStack:  NewLinkStack(),
	}
}

func (z *StackQueue) Add(data int) {
	z.pushStack.Push(data)
	z.pushToPop()

}

func (z *StackQueue) Poll() (int, error) {
	ret, err := z.popStack.Pop()
	z.pushToPop()
	return ret, err

}

func (z *StackQueue) Peek() (int, error) {
	return z.popStack.Peek()

}

func (z *StackQueue) pushToPop() {
	if !z.popStack.IsEmpty() {
		return
	}
	for !z.pushStack.IsEmpty() {
		temp, _ := z.pushStack.Pop()
		z.popStack.Push(temp)
	}

}
