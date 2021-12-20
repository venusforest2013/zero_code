package link_list

import "testing"

func TestSumTwoLinkList(t *testing.T) {
	node11 := &ListNode{
		Data: 9,
	}
	node12 := &ListNode{
		Data: 1,
	}
	node21 := &ListNode{
		Data: 5,
	}
	node22 := &ListNode{
		Data: 1,
	}
	node11.Next = node12
	node21.Next = node22
	PrintLinkList(node11)

	result := SumTwoLinkList(node11, node21)

	PrintLinkList(result)
}
