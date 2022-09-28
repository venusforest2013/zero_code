package search

func BinSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > target {
			right = mid - 1

		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
