package link_list

func RemoveDuplicate(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	var pre *ListNode
	cur := node
	mp := make(map[int]int, 0)
	for cur != nil {
		if mp[cur.Data] == 0 {
			mp[cur.Data] = 1
			pre = cur
			cur = cur.Next
		} else {
			cur = cur.Next
			pre.Next = cur
		}
	}
	return node
}
