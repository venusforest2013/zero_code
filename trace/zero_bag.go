package trace

import "fmt"

// public int maxW = Integer.MIN_VALUE; //存储背包中物品总重量的最大值
// // cw表示当前已经装进去的物品的重量和；i表示考察到哪个物品了；
// // w背包重量；items表示每个物品的重量；n表示物品个数
// // 假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：
// // f(0, 0, a, 10, 100)
// public void f(int i, int cw, int[] items, int n, int w) {
// if (cw == w || i == n) { // cw==w表示装满了;i==n表示已经考察完所有的物品
// if (cw > maxW) maxW = cw;
// return;
// }
// f(i+1, cw, items, n, w);
// if (cw + items[i] <= w) {// 已经超过可以背包承受的重量的时候，就不要再装了
// f(i+1,cw + items[i], items, n, w);
// }
// }

func ZeroOneBagExample() {
	arr := []int{2, 3, 5}
	w := 6
	ret := ZeroOneBag(arr, w)
	fmt.Println(ret)
}
func ZeroOneBag(arr []int, w int) int {
	maxW := 0
	backTrace(0, len(arr), arr, 0, w, &maxW)
	return maxW
}

func backTrace(i int, n int, arr []int, cw int, w int, maxW *int) {
	if cw >= w || i == n {
		if cw > *maxW {
			*maxW = cw
		}
		return
	}

	backTrace(i+1, n, arr, cw, w, maxW)
	if cw+arr[i] <= w {
		backTrace(i+1, n, arr, cw+arr[i], w, maxW)
	}
}
