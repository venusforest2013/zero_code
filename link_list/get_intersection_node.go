package link_list

func GetIntersectionNode(node1, node2 *ListNode) *ListNode {
	if node1 == nil || node2 == nil {
		return nil
	}

	n1 := 0
	p1 := node1
	for p1 != nil {
		n1++
		p1 = p1.Next
	}
	n2 := 0
	p2 := node2
	for p2 != nil {
		n2++
		p2 = p2.Next
	}

	if p1 != p2 {
		return nil
	}

	p1 = node1
	p2 = node2
	if n1 > n2 {
		diff:=n1-n2

		for i := 1; i < diff; i++ {
			p1 = p1.Next
		}
	} else {

		diff:=n2-n1
		for i := 1; i < diff; i++ {
			p2 = p2.Next
		}
	}
	for {
		if p1 == p2 {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
