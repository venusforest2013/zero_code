package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{6, 4, 3, 8, 5, 4, 3, 0}
	BubbleSort(arr)
	fmt.Println(arr)
}
