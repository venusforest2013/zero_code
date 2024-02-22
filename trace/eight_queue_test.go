package trace

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	ret := solveNQueens(4)
	fmt.Println(ret)
	fmt.Println(len(ret))
}
