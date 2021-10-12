package prime

func process(seq, end chan int, res *[]int) {
	//取第一个prime
	prime, ok := <-seq
	if !ok {
		close(end)
		return
	}
	*res = append(*res, prime)
	//定义数据流
	subSeq := make(chan int)
	//开启goroutine
	go process(subSeq, end, res)
	//输入数据流
	for num := range seq {
		if num%prime != 0 {
			subSeq <- num
		}
	}
	//关闭数据流
	close(subSeq)
}

func ProcessPrime(n int) []int {
	//边界判断
	if n < 2 {
		return nil
	}
	//定义结果集合
	res := make([]int, 0)
	//定义结束标识
	end := make(chan int)
	//定义数据流
	seq := make(chan int)
	//开启goroutine
	go process(seq, end, &res)
	//输入数据流
	for i := 2; i <= n; i++ {
		seq <- i
	}
	//close
	close(seq)
	//等待
	<-end
	return res
}
