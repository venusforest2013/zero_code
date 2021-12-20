package link_list

func SumTwoLinkList(node1 *ListNode, node2 *ListNode) *ListNode {
	//边界
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	//for 循环
	p1 := node1
	p2 := node2
	carry := 0
	sum := 0
	dump := &ListNode{}
	cur := dump
	for p1 != nil && p2 != nil {
		sum = p1.Data + p2.Data + carry
		d := sum % 10
		node := &ListNode{
			Data: d,
		}
		cur.Next = node
		cur = node
		carry = sum / 10
		p1 = p1.Next
		p2 = p2.Next
	}
	//剩余
	for p1 != nil {
		sum = p1.Data + carry
		d := sum % 10
		node := &ListNode{
			Data: d,
		}
		cur.Next = node
		cur = node
		p1 = p1.Next
	}
	for p2 != nil {
		sum = p2.Data + carry
		d := sum % 10
		node := &ListNode{
			Data: d,
		}
		cur.Next = node
		cur = node
		p2 = p2.Next
	}
	if carry > 0 {
		node := &ListNode{
			Data: carry,
		}
		cur.Next = node
		cur = node
	}
	//返回
	return dump.Next

}
