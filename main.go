package main

import (
	"fmt"
	"github.com/venusforest2013/zero_code/link_list"
)

func main() {

	link_list.Init()
	result := link_list.GetIntersectionNode(link_list.SampleListNode, link_list.SampleListNode2)
	fmt.Println("哈哈")
	link_list.PrintLinkList(result)
}
