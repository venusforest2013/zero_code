package link_list

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

var SampleListNode *ListNode

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
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	SampleListNode = node1
}

func PrintLinkList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}
