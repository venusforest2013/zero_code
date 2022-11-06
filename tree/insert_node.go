package tree

import "fmt"

func TreeInsertExample() {
	arr := []int{1, 2, 3, 4, 5, 7, 8}
	root := SortedArrayToBST(arr)
	TreeInsert(root, 6)
	ret := BST2SortedArray(root)
	fmt.Println(ret)

}

func TreeInsert(node *TreeNode, data int) {

	if node == nil {

		node = &TreeNode{
			Val: data,
		}
		return
	}

	for node != nil {
		if data < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{
					Val: data,
				}
				return
			} else {
				node = node.Left
			}

		} else {
			if node.Right == nil {
				node.Right = &TreeNode{
					Val: data,
				}
				return
			} else {
				node = node.Right
			}

		}
	}
}
