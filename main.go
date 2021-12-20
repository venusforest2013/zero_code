package main

import (
	"github.com/venusforest2013/zero_code/link_list"
)

func main() {
	link_list.PrintLinkList(link_list.SampleListNode)
	result := link_list.RemoveDuplicate(link_list.SampleListNode)
	link_list.PrintLinkList(result)
}
