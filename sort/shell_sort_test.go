package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	ShellSort(arr)
	fmt.Println(arr)
}
