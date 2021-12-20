package link_list

func IsCycleLinkList(node *ListNode) bool {
	if node == nil || node.Next == nil {
		return false
	}

	slow := node
	fast := node.Next

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
