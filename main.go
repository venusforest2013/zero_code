package main

import (
	"fmt"

	"github.com/venusforest2013/zero_code/tree"
)

func main() {
	tree.DeleteNodeExample()
}

// weight:物品重量，n:物品个数，w:背包可承载重量
// public int knapsack(int[] weight, int n, int w) {
// boolean[][] states = new boolean[n][w+1]; // 默认值false
// states[0][0] = true;  // 第一行的数据要特殊处理，可以利用哨兵优化
// if (weight[0] <= w) {
// states[0][weight[0]] = true;
// }
// for (int i = 1; i < n; ++i) { // 动态规划状态转移
// for (int j = 0; j <= w; ++j) {// 不把第i个物品放入背包
// if (states[i-1][j] == true) states[i][j] = states[i-1][j];
// }
// for (int j = 0; j <= w-weight[i]; ++j) {//把第i个物品放入背包
// if (states[i-1][j]==true) states[i][j+weight[i]] = true;
// }
// }
// for (int i = w; i >= 0; --i) { // 输出结果
// if (states[n-1][i] == true) return i;
// }
// return 0;
// }

func Knapsack(items []int, w int) int {
	n := len(items)
	states := make([][]bool, n)
	for i := 0; i < n; i++ {
		temp := make([]bool, w+1)
		states[i] = temp
	}
	states[0][0] = true
	if items[0] <= w {
		states[0][items[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			if states[i-1][j] {
				states[i][j] = true
			}
		}
		for j := w - items[i]; j >= 0; j-- {
			if states[i-1][j] {
				states[i][j+items[i]] = true
			}
		}
	}
	fmt.Println(states)
	for j := w; j >= 0; j-- {
		if states[n-1][j] {
			return j
		}
	}
	return 0

}
