package stack

import "fmt"

type SliceStack struct {
	arr []int
}

func NewSliceStatck() *SliceStack {
	arr := make([]int, 0)
	s := &SliceStack{
		arr: arr,
	}
	return s
}

func (z *SliceStack) Push(element int) {
	z.arr = append(z.arr, element)

}

func (z *SliceStack) Pop() int {
	l := len(z.arr)
	ret := z.arr[l-1]
	z.arr = z.arr[:l-1]
	return ret
}

func (z *SliceStack) PrintStack() {

	for _, v := range z.arr {
		fmt.Println(v)
	}
}
