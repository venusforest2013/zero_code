package main

import (
	"fmt"

	"github.com/venusforest2013/zero_code/search"
)

func main() {
	arr := []int{1, 2, 3, 4, 4, 5}
	ret := search.BinSearchLeft(arr, 60)
	fmt.Println(ret)
}
