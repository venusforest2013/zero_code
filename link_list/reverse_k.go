package link_list

func ReverseK(node *ListNode, k int) *ListNode {
	//定义结果
	var result *ListNode
	var findResult bool


	//上k个节点的尾部节点
	preTail := &ListNode{}
	preTail.Next = node
	//游标节点
	curTail := preTail
	//下k个节点的头部节点
	var nextHeader *ListNode
	//for 循环
	for curTail != nil {
		//找出当前k个节点
		for i := 0; i < k; i++ {
			curTail = curTail.Next
			//凑不够k个直接返回
			if curTail == nil {
				return result
			}
		}

		//拿出下k个节点的头节点
		nextHeader = curTail.Next
		curTail.Next = nil
		newHeader := ReverseLinkListRecursion(preTail.Next)
		if !findResult {
			result = newHeader
			findResult = true
		}
		newCurTail := preTail.Next
		preTail.Next.Next = nextHeader
		preTail.Next = newHeader
		preTail = newCurTail
		curTail = newCurTail

	}
	return result

}
