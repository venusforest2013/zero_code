package main

import (
	"github.com/venusforest2013/zero_code/link_list"
)

func main() {

	result := link_list.MergeSortedLinkList(link_list.SampleListNode, link_list.SampleListNode2)
	link_list.PrintLinkList(result)
}
