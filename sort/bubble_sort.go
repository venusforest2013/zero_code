package sort

func BubbleSort(arr []int) {
	//边界判断
	n := len(arr)
	if n <= 1 {
		return
	}
	//循环沉底
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
