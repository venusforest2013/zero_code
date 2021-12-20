package link_list

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

var SampleListNode *ListNode
var SampleListNode2 *ListNode

var NodeDel = &ListNode{
	Data: 100,
}

func Init() {

	node1 := &ListNode{
		Data: 1,
	}
	node2 := &ListNode{
		Data: 2,
	}
	node3 := &ListNode{
		Data: 3,
	}

	node1.Next = node2
	node2.Next = NodeDel
	NodeDel.Next = node3
	SampleListNode = node1

	node11 := &ListNode{
		Data: 11,
	}
	node12 := &ListNode{
		Data: 12,
	}

	node11.Next = node12
	node12.Next = NodeDel
	SampleListNode2 = node11
}

func PrintLinkList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}
