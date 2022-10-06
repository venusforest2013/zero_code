package stack

import (
	"errors"
	"fmt"
)

type SliceStack struct {
	arr []int
}

func (z *SliceStack) Push(element int) {
	z.arr = append(z.arr, element)

}

func (z *SliceStack) IsEmpty() bool {
	return len(z.arr) == 0
}

func (z *SliceStack) Pop() (int, error) {

	l := len(z.arr)
	if l == 0 {
		return 0, errors.New("stack empty")
	}
	ret := z.arr[l-1]
	z.arr = z.arr[:l-1]
	return ret, nil
}

func (z *SliceStack) PrintStack() {

	for _, v := range z.arr {
		fmt.Println(v)
	}
}
func (z *SliceStack) Peek() (int, error) {

	size := len(z.arr)
	if size == 0 {
		return 0, errors.New("stacke emptye")
	}
	return z.arr[0], nil
}
