package stack

import "fmt"

//栈逆序

func ReverseStackExample() {
	s := NewSliceStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.PrintStack()
	fmt.Println("----")
	ReverseStack(s)
	s.PrintStack()
}

func Get2RemoveLastElement(s stacker) (int, error) {
	ret, err := s.Pop()
	if err != nil {
		return 0, err
	}
	if s.IsEmpty() {
		return ret, nil
	} else {
		last, err := Get2RemoveLastElement(s)
		if err != nil {
			return 0, err
		}
		s.Push(ret)
		return last, nil
	}
}

func ReverseStack(s stacker) error {
	if s.IsEmpty() {
		return nil
	}
	i, err := Get2RemoveLastElement(s)
	if err != nil {
		return err
	}
	ReverseStack(s)
	s.Push(i)
	return nil

}
