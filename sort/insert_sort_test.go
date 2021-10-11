package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort( t *testing.T){
	arr :=[]int{6,5,4,3,2,1}
	InsertSort(arr)
	fmt.Println(arr)
}
