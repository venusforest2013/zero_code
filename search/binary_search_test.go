package search

import (
	"fmt"
	"testing"
)

func TestBinSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	ret := BinSearch(arr, 3)
	fmt.Println(ret)
}
