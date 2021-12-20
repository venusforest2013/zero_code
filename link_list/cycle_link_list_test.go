package link_list

import (
	"fmt"
	"testing"
)

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
	node6.Next = node3
	SampleListNode3 = node1

}

var SampleListNode3 *ListNode

func TestIsCycleLinkList(t *testing.T) {
	result := IsCycleLinkList(SampleListNode3)
	fmt.Println(result)
}
