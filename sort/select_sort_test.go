package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{8, 6, 5, 4, 3, 2, 1}
	SelectSort(arr)
	fmt.Println(arr)
}
