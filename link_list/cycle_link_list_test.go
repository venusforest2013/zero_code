package link_list

import (
	"fmt"
	"testing"
)

func TestIsCycleLinkList(t *testing.T) {
	result := IsCycleLinkList(SampleListNode3)
	fmt.Println(result)
}
