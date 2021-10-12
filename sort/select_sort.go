package sort

func SelectSort(arr []int) {
	//边界判断
	n := len(arr)
	if n <= 1 {
		return
	}
	//循环交换最小值
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
