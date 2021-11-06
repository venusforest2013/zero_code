package tree

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil

	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		nextQueue := make([]*TreeNode, 0)
		resRow := make([]int, 0)
		for i := 0; i < len(queue); i++ {
			resRow = append(resRow, queue[i].Val)
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)

			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)

			}

		}
		queue = nextQueue
		res = append(res, resRow)
	}
	return res
}

func LeverOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil

	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		resRow := make([]int, 0)
		nextQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			resRow = append(resRow, queue[i].Val)
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		res = append(res, resRow)
		queue = nextQueue
	}
	return res
}
