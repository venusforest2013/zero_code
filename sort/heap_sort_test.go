package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{6,5,4,3}
	HeapSort(arr)
	fmt.Println(arr)
}
