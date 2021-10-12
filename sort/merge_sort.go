package sort

//合并有序数组
func merge(l, r []int) []int {
	res := make([]int, 0)
	lIndex, rIndex := 0, 0
	lLen := len(l)
	rLen := len(r)
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] < r[rIndex] {
			res = append(res, l[lIndex])
			lIndex++
		} else {
			res = append(res, r[rIndex])
			rIndex++
		}
	}
	if lIndex < lLen {
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}
	return res
}

func mergeSort(arr []int) []int {
	//边界判断
	n := len(arr)
	if n <= 1 {
		return arr
	}
	//选基准
	mid := n / 2
	//分治
	l := mergeSort(arr[:mid])
	r := mergeSort(arr[mid:])
	//合并
	return merge(l, r)
}
