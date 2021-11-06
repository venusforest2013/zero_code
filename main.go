package main

import (
	"fmt"

	"github.com/venusforest2013/zero_code/tree"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	root := tree.CreateBST(arr)
	fmt.Println(root.Left.Val)
	res := tree.PreOrder(root)
	fmt.Println(res)

}
