package search

func BinSearchRight(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || (right >= 0 && arr[right] != target) {
		return -1
	}
	return right
}
