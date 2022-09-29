package search

func BinSearchLeft(arr []int, target int) int {
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
			right = mid - 1
		}
	}

	if left >= n || (left < n && arr[left] != target) {
		return -1
	}
	return left

}
