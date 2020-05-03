package linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	Val  interface{}
	Next *node
	Prev *node
}

type ListNode struct {
	Head *node
	Tail *node
}

func (l *ListNode) First() (interface{}, error) {
	if l.Head != nil {
		return l.Head.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

func (l *ListNode) Last() (interface{}, error) {
	if l.Tail != nil {
		return l.Tail.Val, nil
	}
	return -1, errors.New("tail is not yet initialized")
}

func (l *ListNode) Push(val interface{}) {
	n := &node{Val: val}
	if l.Head == nil {
		l.Head = n
	} else {
		n.Prev = l.Tail
		l.Tail.Next = n
	}
	l.Tail = n
}

func (l *ListNode) PrintListNode() {
	l.iterateList(true, nil, false, false)
}

func (l *ListNode) PrintReverseList() {
	l.reverseIterateList(true)
}

func (l *ListNode) Size() int {
	size, _ := l.iterateList(false, nil, false, false)
	return size
}

func (l *ListNode) GetFirstMatchIndex(val interface{}) int {
	matchIndex, _ := l.iterateList(false, val, false, false)
	return matchIndex + 1
}

func (l *ListNode) isEmpty() bool {
	return l.Head == nil
}
func (l *ListNode) RemoveFirst() (*ListNode, error) {
	if !l.isEmpty() || l.Size() >= 1 {
		l.Head = l.Head.Next
		return l, nil
	}
	return l, errors.New("list is empty, or there's only one element")
}

func (l *ListNode) DebugPrintList() {
	l.iterateList(true, nil, true, false)
}

func (l * ListNode) ToSlice() []interface{} {
	_, slice := l.iterateList(false, nil, false, true)
	return slice
}

func (l *ListNode) iterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool, returnSlice bool) (int, []interface{}) {
	len := 0
	curr := l.Head
	var listToSlice []interface{}
	for curr != nil {
		if shouldPrint {
			fmt.Print((*curr).Val)
			if shouldDebug {
				fmt.Printf("\t Prev: %p \t current: %p \t Next: %p", curr.Prev, curr, curr.Next)
			}
			fmt.Println()
		}
		if returnSlice {
			listToSlice = append(listToSlice, (*curr).Val)
		}
		if searchKey != nil && searchKey == curr.Val {
			return len - 1, nil
		}
		curr = curr.Next
		len++
	}
	return len, listToSlice
}

func (l *ListNode) reverseIterateList(shouldPrint bool) {
	curr := l.Tail
	for curr != nil {
		if shouldPrint {
			fmt.Println((*curr).Val)
		}
		curr = curr.Prev
	}
}
