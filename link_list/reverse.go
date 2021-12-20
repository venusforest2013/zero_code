package link_list

func ReverseLinkListIteration(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	var pre *ListNode
	cur := node

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

func ReverseLinkListRecursion(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	childHeader := ReverseLinkListRecursion(node.Next)
	node.Next.Next = node
	node.Next = nil
	return childHeader
}
