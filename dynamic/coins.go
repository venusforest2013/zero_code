package dynamic

import (
	"fmt"
	"math"
)

func CoinsExample() {
	arr := []int{2, 3, 5}
	ret := MinCoinsCount(arr, 11)
	fmt.Println(ret)
}

//凑硬币
func MinCoinsCount(arr []int, target int) int {
	dp := make([]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = math.MaxInt
	}

	for _, v := range arr {
		dp[v] = 1
	}

	for i := 1; i <= target; i++ {
		for _, j := range arr {
			if i-j >= 0 {
				temp := dp[i-j] + 1
				if temp < dp[i] && temp > 0 {
					dp[i] = temp
				}
			}
		}
	}

	fmt.Println(dp)
	return dp[target]
}
