package main

import (
	"fmt"

	"github.com/venusforest2013/zero_code/sub_string"
)

func main() {
	str := "iabca"
	res := sub_string.GetLongestNoRepeatSubString(str)
	fmt.Println(res)

}
