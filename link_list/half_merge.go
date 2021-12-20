package link_list

func HalfMerge(node *ListNode) *ListNode {
	//边界判断
	if node == nil || node.Next == nil {
		return node
	}
	//找到中点
	node1 := GetHalfPointOfLinkList(node)
	node2 := node1.Next
	node1.Next = nil
	//后半段反转
	node2Reverse := ReverseLinkListRecursion(node2)
	//两端合并
	result := MergeTwoLinkList(node, node2Reverse)
	return result
}

//找到单链表中点
func GetHalfPointOfLinkList(node *ListNode) *ListNode {
	slow := node
	fast := node.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func MergeTwoLinkList(node1 *ListNode, node2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	for node1 != nil && node2 != nil {
		n1 := &ListNode{
			Data: node1.Data,
		}
		n2 := &ListNode{
			Data: node2.Data,
		}
		cur.Next = n1
		n1.Next = n2
		cur = n2
		node1 = node1.Next
		node2 = node2.Next
	}
	if node1 != nil {
		cur.Next = node1
	}
	if node2 != nil {
		cur.Next = node2
	}
	return dump.Next
}
