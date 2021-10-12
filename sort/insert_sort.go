package sort

func InsertSort(arr []int) {
	//边界判断
	n := len(arr)
	if n <= 1 {
		return
	}
	//循环向右移动
	for i := 1; i < n; i++ {
		temp := arr[i]
		key := i - 1
		for key >= 0 && arr[key] > temp {
			arr[key+1] = arr[key]
			key--
		}
		arr[key+1] = temp
	}

}
