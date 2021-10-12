package sort

//单节点堆化
func heapify(arr []int, pos, count int) {
	for {
		child := pos * 2
		if child > count {
			break
		}
		if child < count && arr[child+1] > arr[child] {
			child = child + 1
		}
		if arr[pos] > arr[child] {
			break
		}
		arr[pos], arr[child] = arr[child], arr[pos]
		pos = child
	}
}

func HeapSort(arr []int) {
	//边界判断
	count := len(arr) - 1
	if count <= 1 {
		return
	}
	//初始化
	for i := count / 2; i >= 1; i-- {
		heapify(arr, i, count)
	}
	//头节点循环堆化
	for count >= 1 {
		arr[1], arr[count] = arr[count], arr[1]
		count--
		heapify(arr, 1, count)
	}
}
