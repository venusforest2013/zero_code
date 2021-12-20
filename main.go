package main

func main() {
}

func ProcessPrime(n int) []int {
	//边界判断
	if n < 2 {
		return nil
	}
	//中间变量
	seq := make(chan int)
	end := make(chan int)
	res := make([]int, 0)
	go Process(seq, end, &res)
	//输出第一层序列
	for i := 2; i <= n; i++ {
	}
	//关闭chan
	//等待
}

func Process(seq chan int, end chan int, res *[]int) {
	//取第一个素数
	//中间变量
	//输出下一个序列
}
