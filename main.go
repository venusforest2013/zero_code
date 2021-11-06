package main

import (
	"fmt"

	"github.com/venusforest2013/zero_code/spiral_order"
)

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := spiral_order.SpiralOrder(arr)
	fmt.Println(res)

}
