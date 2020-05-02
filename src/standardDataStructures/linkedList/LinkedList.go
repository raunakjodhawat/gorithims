
package linkedList

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
	l.iterateList(true, nil, false)
}

func (l *ListNode) PrintReverseList(){
	l.reverseIterateList(true)
}

func (l *ListNode) Length() int {
	return l.iterateList(false, nil, false)
}

func (l *ListNode) GetFirstMatchIndex(val interface{}) int {
	return l.iterateList(false, val, false) + 1
}

func (l *ListNode) isEmpty() bool {
	return l.Head == nil
}
func (l *ListNode) RemoveFirst() (*ListNode, error) {
	if !l.isEmpty() || l.Length() >= 1 {
		l.Head = l.Head.Next
		return l, nil
	}
	return l, errors.New("list is empty, or there's only one element")
}

func (l *ListNode)DebugPrintList(){
	l.iterateList(true, nil, true)
}

func (l *ListNode) iterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool) int {
	len := 0
	curr := l.Head
	for curr != nil {
		if shouldPrint {
			fmt.Print((*curr).Val)
			if shouldDebug {
				fmt.Printf("\t Prev: %p \t current: %p \t Next: %p", curr.Prev, curr , curr.Next)
			}
			fmt.Println()
		}
		if searchKey != nil && searchKey == curr.Val {
			return len - 1
		}
		curr = curr.Next
		len += 1
	}
	return len
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
