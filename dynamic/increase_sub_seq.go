package dynamic

import "fmt"

// class Solution {
// public int lengthOfLIS(int[] nums) {
// if (nums.length == 0) {
// return 0;
// }
// int[] dp = new int[nums.length];
// dp[0] = 1;
// int maxans = 1;
// for (int i = 1; i < nums.length; i++) {
// dp[i] = 1;
// for (int j = 0; j < i; j++) {
// if (nums[i] > nums[j]) {
// dp[i] = Math.max(dp[i], dp[j] + 1);
// }
// }
// maxans = Math.max(maxans, dp[i]);
// }
// return maxans;
// }
// }

func GetMaxIncreasedSubSeqLenExample() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	ret := getMaxIncreasedSubSeqLen(arr)
	fmt.Println(ret)
}

func getMaxIncreasedSubSeqLen(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ret := 0
	for _, v := range dp {
		if v > ret {
			ret = v
		}
	}
	fmt.Println(dp)
	return ret

}

func max(a, b int) int {
	ret := a
	if a < b {
		ret = b
	}
	return ret
}
