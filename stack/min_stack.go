package stack

import "fmt"

func MinStackExample() {
	ms := NewMinStack()
	ms.Push(3)
	ms.Push(4)
	ms.Push(5)
	ms.Push(1)
	ms.Push(2)
	ms.Push(1)
	// ms.min.Print()
	ms.min.PrintStack()
	fmt.Println("====")
	ms.normal.PrintStack()
	fmt.Println("====")
	ret, _ := ms.GetMin()
	fmt.Println(ret)
	// ms.normal.Print()
	// min, _ := ms.GetMin()
	// fmt.Println(min)
	// fmt.Println("---")

	// ms.Pop()
	// min, _ = ms.GetMin()
	// fmt.Println(min)
	// fmt.Println("---")
	// ms.Pop()
	// min, _ = ms.GetMin()
	// fmt.Println(min)
	// fmt.Println("---")
	// ms.Pop()
	// min, _ = ms.GetMin()
	// fmt.Println(min)
}

type MinStack struct {
	normal *SliceStack
	min    *LinkStack
}

func NewMinStack() *MinStack {
	return &MinStack{
		normal: NewSliceStack(),
		min:    NewLinkStack(),
	}
}
func NewLinkStack() *LinkStack {
	return &LinkStack{}
}
func NewSliceStack() *SliceStack {
	return &SliceStack{
		arr: make([]int, 0),
	}
}
func (z *MinStack) Push(data int) {
	z.normal.Push(data)
	if z.min.IsEmpty() {
		z.min.Push(data)
	} else {
		minData := data
		currentMinData, _ := z.min.Peek()
		if currentMinData < minData {
			minData = currentMinData
		}
		z.min.Push(minData)
	}
}
func (z *MinStack) Pop() (int, error) {
	ret, err := z.normal.Pop()
	if err != nil {
		return 0, err
	}
	z.min.Pop()
	return ret, err

}

func (z *MinStack) GetMin() (int, error) {
	return z.min.Peek()

}
