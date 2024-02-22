package recursive

import (
	"fmt"
	"testing"
)

func TestMaxShare(t *testing.T) {
	arr := []int{3, 2, 4, 3, 6}
	ret := maxShare(arr)
	fmt.Println(ret)
}
