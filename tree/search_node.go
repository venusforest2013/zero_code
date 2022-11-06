package tree

import "fmt"

func TreeSearchExample() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	root := SortedArrayToBST(arr)
	fmt.Println(root)
	ret := LevelOrder(root)
	fmt.Println(ret)
	ret1 := TreeSearch(root, 8)
	fmt.Println(ret1)
}

func TreeSearch(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return nil
	}

	node := root
	for node != nil {
		if data == node.Val {
			return node
		} else if data < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil

}
