package main

import (
	"fmt"

	"github.com/venusforest2013/zero_code/search"
)

func main() {
	arr := []int{1, 2, 3, 4, 4, 5}
	ret := search.BinSearchRight(arr, 4)
	fmt.Println(ret)
}

func heapify(arr []int, pos, count int) {
	for {
		child := pos * 2
		if child > count {
			break
		}
		if child < count && arr[child+1] > arr[child] {
			child = child + 1
		}
		if arr[pos] > arr[child] {
			break
		}
		arr[pos], arr[child] = arr[child], arr[pos]
		pos = child
	}

}

func heapSort(arr []int) {
	count := len(arr) - 1

	if count <= 1 {
		return
	}

	for i := count / 2; i >= 1; i-- {
		heapify(arr, i, count)
	}

	for count >= 1 {
		arr[1], arr[count] = arr[count], arr[1]
		count--
		heapify(arr, 1, count)
	}
}
