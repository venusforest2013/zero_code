package sort

func QuickSort(arr []int, left, right int) {
	//边界
	if right <= left {
		return
	}
	//选基准
	mid := (left + right) / 2
	key := arr[mid]
	i := left
	j := right
	//左右分堆
	for {
		for i <= j && arr[i] < key {
			i++
		}
		for i <= j && arr[j] > key {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	//递归
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}
