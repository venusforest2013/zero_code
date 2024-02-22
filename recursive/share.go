package recursive

import (
	"fmt"
)

func testShare(arr []int, n int, goal int, m int, groupId int, groupSum int, aux []int) bool {
	if goal < 0 {
		return false
	}
	if goal == 0 {

		groupId++
		goal = groupSum
		if groupId == m+1 {
			return true
		}
	}
	for i := 0; i < n; i++ {
		if aux[i] != 0 {
			continue
		}
		aux[i] = groupId
		if testShare(arr, n, goal-arr[i], m, groupId, groupSum, aux) {
			return true
		}
		aux[i] = 0
	}
	return false
}

func maxShare(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
	}

	for m := n; m >= 2; m-- {
		//每个分组对m取余，应该无余数
		if sum%m != 0 {
			continue
		}
		subSum := sum / m
		aux := make([]int, n)
		if testShare(arr, n, subSum, m, 1, subSum, aux) {
			fmt.Println(aux)
			return m
		}

	}
	return 0
}
