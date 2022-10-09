package divide2conquer

import "fmt"

func Divide2ConquerExample() {
	arr := []int{5, 4, 3, 2, 1}
	var count int
	ret := MergeSort(arr, &count)
	fmt.Println(ret)
	fmt.Println(count)
}

func merge(l, r []int, count *int) []int {
	lLen := len(l)
	rLen := len(r)
	lIndex := 0
	rIndex := 0
	ret := make([]int, 0)
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] <= r[rIndex] {
			ret = append(ret, l[lIndex])
			lIndex++
		} else {
			ret = append(ret, r[rIndex])
			*count = *count + (lLen - lIndex)
			rIndex++
		}
	}

	if lIndex < lLen {
		ret = append(ret, l[lIndex:]...)
	}
	if rIndex < rLen {
		ret = append(ret, r[rIndex:]...)
	}
	return ret
}

func MergeSort(arr []int, count *int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	l := MergeSort(arr[:n/2], count)
	r := MergeSort(arr[n/2:], count)

	ret := merge(l, r, count)
	return ret
}
