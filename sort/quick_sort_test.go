package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{4, 3, 2, 1}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
