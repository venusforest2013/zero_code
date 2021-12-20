package main

import (
	"fmt"
	"github.com/venusforest2013/zero_code/link_list"
)

func main() {

	link_list.Init()
	link_list.PrintLinkList(link_list.SampleListNode)
	result := link_list.RemoveSpecificNode(link_list.SampleListNode, link_list.NodeDel)
	fmt.Println("哈哈")
	link_list.PrintLinkList(result)
}
