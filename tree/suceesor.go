package tree

//中序列后继节点，比节点大的最小值的节点
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}
	node := root
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}
