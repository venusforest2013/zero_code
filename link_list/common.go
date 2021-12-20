package link_list

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

var SampleListNode *ListNode
var SampleListNode2 *ListNode

func init() {
	node1 := &ListNode{
		Data: 1,
	}
	node2 := &ListNode{
		Data: 2,
	}
	node3 := &ListNode{
		Data: 3,
	}
	node4 := &ListNode{
		Data: 4,
	}
	node5 := &ListNode{
		Data: 5,
	}
	node6 := &ListNode{
		Data: 6,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	SampleListNode = node1

	node11 := &ListNode{
		Data: 11,
	}
	node12 := &ListNode{
		Data: 12,
	}
	node13 := &ListNode{
		Data: 13,
	}
	node14 := &ListNode{
		Data: 14,
	}
	node15 := &ListNode{
		Data: 15,
	}
	node11.Next = node12
	node12.Next = node13
	node13.Next = node14
	node14.Next = node15
	SampleListNode2 = node11
}

func PrintLinkList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}
