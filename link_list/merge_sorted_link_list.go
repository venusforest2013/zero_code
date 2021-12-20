package link_list

func MergeSortedLinkList(node1, node2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	for node1 != nil && node2 != nil {

		if node1.Data < node2.Data {
			node := &ListNode{
				Data: node1.Data,
			}
			cur.Next = node
			node1 = node1.Next
		} else {
			node := &ListNode{
				Data: node2.Data,
			}
			cur.Next = node
			node2 = node2.Next
		}
		cur = cur.Next
	}
	if node1 == nil {
		cur.Next = node2
	}
	if node2 == nil {
		cur.Next = node1
	}
	return dump.Next
}
