package main

import (
	"fmt"
)

func main() {
	a := 3
	ret := onesCount(a)
	fmt.Println(ret)
}
func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}
