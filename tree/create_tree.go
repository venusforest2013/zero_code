package tree

import "fmt"

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

func CreateBST2Examaple() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	root := CreateBST2(arr)
	ret := LevelOrder(root)
	fmt.Println(ret)
}

//数组第一个元素不用
func CreateBST2(arr []int) *TreeNode {
	n := len(arr)
	if n <= 1 {
		return nil
	}

	nodeArr := make([]*TreeNode, n)
	for i := 0; i < n; i++ {
		nodeArr[i] = &TreeNode{
			Val: arr[i],
		}
	}

	for i := 1; i < n/2; i++ {
		nodeArr[i].Left = nodeArr[2*i]
		nodeArr[i].Right = nodeArr[2*i+1]

	}
	if n%2 == 1 {
		nodeArr[n/2+1].Left = nodeArr[n-1]
	}

	return nodeArr[1]

}
