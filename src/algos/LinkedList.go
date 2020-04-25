package algos

import (
	"errors"
	"fmt"
)

type node struct {
	Val  interface{}
	Next *node
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

func (l *ListNode) Push(val int) {
	n := &node{Val: val}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func (l *ListNode) PrintListNode() {
	l.iterateList(true, nil)
}

func (l *ListNode) Length() int {
	return l.iterateList(false, nil)
}

func (l *ListNode) GetFirstMatchIndex(val interface{}) int {
	return l.iterateList(false, val)
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

func (l *ListNode) iterateList(shouldPrint bool, searchKey interface{}) int {
	len := 1
	curr := l.Head
	for curr != nil {
		if shouldPrint {
			fmt.Println((*curr).Val)
		}
		if searchKey != nil && searchKey == curr.Val {
			return len - 1
		}
		curr = curr.Next
		len += 1
	}
	return len
}