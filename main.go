package main

import (
	"fmt"

	"github.com/venusforest2013/zero_code/stack"
)

func main() {

	s := &stack.LinkStack{}
	s.Push(1)
	s.Push(2)
	s.PrintStack()
	ret := s.Pop()
	fmt.Println(ret)
	fmt.Println("====")
	s.PrintStack()
}
