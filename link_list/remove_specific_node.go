package link_list

func RemoveSpecificNode(node *ListNode, delNode *ListNode) *ListNode {
	if node == nil {
		return node
	}

	cur := node
	for cur != nil {
		if cur == delNode {
			if cur.Next == nil {
				cur = nil
			} else {
				cur.Data = cur.Next.Data
				cur.Next = cur.Next.Next
			}
			break
		}
		cur = cur.Next
	}
	return node
}
