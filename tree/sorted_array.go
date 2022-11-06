package tree

import "fmt"

func BST2SortedArrayExample() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	node := SortedArrayToBST(arr)
	fmt.Println(node)
	ret := BST2SortedArray(node)
	fmt.Println(ret)
}

//有序数组转二叉树
func SortedArrayToBST(nums []int) *TreeNode {
	return helpers(nums, 0, len(nums)-1)
}

func helpers(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helpers(nums, left, mid-1)
	root.Right = helpers(nums, mid+1, right)
	return root
}

func BST2SortedArray(node *TreeNode) []int {
	ret := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ret = append(ret, node.Val)
		dfs(node.Right)
	}
	dfs(node)
	return ret
}
