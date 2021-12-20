package link_list

func GetKthFromLast(node *ListNode, k int) *ListNode {
	fast := node

	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	slow := node
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
