package tree

import "math"



func IsValid(root *TreeNode) bool {
	return IsValidHelper(root, math.MinInt32, math.MaxInt32)
}

func IsValidHelper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && (root.Left.Val > root.Val || root.Val < min) {
		return false
	}

	if root.Right != nil && (root.Right.Val < root.Val || root.Val > max) {
		return false
	}
	return true
}
