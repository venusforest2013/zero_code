package tree

import "fmt"

func DeleteNodeExample() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := SortedArrayToBST(arr)
	DeleteNode(root, 8)
	ret := BST2SortedArray(root)
	fmt.Println(ret)
}

func DeleteNode(root *TreeNode, data int) {
	if root == nil {
		return
	}

	//尝试找到删除节点
	p := root
	var pp *TreeNode
	for p != nil && p.Val != data {
		if data < p.Val {
			pp = p
			p = p.Left
		} else {
			pp = p
			p = p.Right
		}
	}
	//未找到
	if p == nil {
		return
	}

	//删除节点左右都有子节点,找到删除节点
	if p.Left != nil && p.Right != nil {
		minP := p.Right
		minPP := p
		for minP != nil {

			minPP = minP
			minP = minP.Left
		}
		p.Val = minP.Val
		p = minP
		pp = minPP

	}
	//找到补位的孩子节点
	var child *TreeNode
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	//用孩子节点补位
	if pp == nil {
		root = child
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}

}
