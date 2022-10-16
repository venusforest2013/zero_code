package dynamic

import "fmt"

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
