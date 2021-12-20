package main

import (
	"github.com/venusforest2013/zero_code/link_list"
)

func main() {
	result := link_list.GetKthFromLast(link_list.SampleListNode, 3)
	link_list.PrintLinkList(result)
}
