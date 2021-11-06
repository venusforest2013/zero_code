package tree

func CreateBST(arr []int) *TreeNode {
	//边界判断
	n := len(arr)
	if n == 0 {
		return nil
	}
	//初始化node数据
	nodeArr := make([]*TreeNode, n)
	for i := 0; i < n; i++ {
		nodeArr[i] = &TreeNode{
			Val: arr[i],
		}
	}
	//建立node之间的连接
	for i := 0; i < (n-1)/2; i++ {
		nodeArr[i].Left = nodeArr[2*i+1]
		nodeArr[i].Right = nodeArr[2*i+2]

	}
	//特殊奇偶关系处理
	if n%2 == 0 {
		nodeArr[(n-1)/2].Left = nodeArr[n-1]
	}
	return nodeArr[0]
}
