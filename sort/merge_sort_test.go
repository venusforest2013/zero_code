package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	res := mergeSort(arr)
	fmt.Println(res)
}
