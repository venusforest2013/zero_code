package sort

func ShellSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for step := n / 2; step > 0; step = step / 2 {

		for i := step; i < n; i = i + step {
			temp := arr[i]
			key := i - step
			for key >= 0 && arr[key] > temp {
				arr[key+step] = arr[key]
				key = key - step
			}
			arr[key+step] = temp
		}
	}
}
