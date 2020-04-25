package algos

import (
	"fmt"
	"testing"
)

func TestListNode_Push(t *testing.T) {
	list := ListNode{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(6)
	list.Push(7)
	list.Push(8)
	fmt.Println(list.Length())
}

func TestListNode_Length(t *testing.T) {
	fmt.Println()
}
