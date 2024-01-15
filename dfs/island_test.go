package dfs

import (
	"fmt"
	"testing"
)

func TestNumIsland(t *testing.T) {
	row0 := []byte{'1', '1', '0', '0', '0'}
	row1 := []byte{'1', '1', '0', '0', '0'}
	row2 := []byte{'0', '0', '1', '0', '0'}
	row3 := []byte{'0', '0', '0', '1', '1'}
	arr := [][]byte{
		row0,
		row1,
		row2,
		row3,
	}
	ret := GetIslandCount(arr)
	fmt.Println(ret)
}
