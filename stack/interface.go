package stack

type stacker interface {
	IsEmpty() bool
	Push(int)
	Peek() (int, error)
	Pop() (int, error)
	PrintStack()
}
