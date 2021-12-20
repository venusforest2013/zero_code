package link_list

func ReverseNeighbour(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	dump := &ListNode{}
	pre := dump
	cur := node
	pre.Next = cur
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next, pre.Next, next.Next = next.Next, next, cur
		pre = cur
		cur = cur.Next
	}
	return dump.Next

}
