// linkedlist package contains all functions pertaining to working with linked list in Go. As per java SE 14 (https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/LinkedList.html)
package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

type node struct {
	Val  interface{}
	Next *node
	Prev *node
}

/**
ListNode type, stores the head and Tail. It represents a list of address with each list instantiated having a default nil Head and Tail
*/
type ListNode struct {
	Head   *node
	Tail   *node
	Length int
}

func (l *ListNode) Add(element interface{}, startIndexSlice ...int) error {

	startIndex, err := l.getStartingIndex(startIndexSlice...)
	if err != nil {
		return err
	}

	slice := reflect.ValueOf(element)
	if slice.Kind() != reflect.Slice {
		n := &node{Val: element}
		if l.Head == nil {
			l.Head = n
			l.Tail = n
		} else {
			if startIndex == 0 {
				headCopy := l.Head
				l.Head = n
				n.Next = headCopy
			} else if startIndex != l.Length {
				curr := l.Head
				for i := 0; i < l.Length; i++ {
					if i == startIndex-1 {
						copyNext := curr.Next
						curr.Next = n
						n.Next = copyNext
						break
					}
					curr = curr.Next
				}
			} else {
				l.Tail.Next = n
				l.Tail = n
			}
		}
		l.Length += 1
	} else {
		elementsSlice := make([]interface{}, slice.Len())

		for i := 0; i < slice.Len(); i++ {
			elementsSlice[i] = slice.Index(i).Interface()
		}

		curr := l.Head
		if startIndex == 0 {
			for i := 0; i < slice.Len(); i++ {
				l.Add(elementsSlice[i], startIndex)
				startIndex += 1
			}
		} else {
			for i := 0; i < l.Length; i++ {
				if i == startIndex-1 {
					currCopy := curr.Next
					for j := 0; j < slice.Len(); j++ {
						innerNode := &node{Val: elementsSlice[j]}
						curr.Next = innerNode
						curr = curr.Next
						l.Length += 1
					}
					curr.Next = currCopy
					if i-1+startIndex == l.Length {
						l.Tail = curr
					}
					break
				}
				curr = curr.Next
			}
		}
	}
	return nil
}

func (l *ListNode) AddAll(elementsInterface interface{}, startIndexSlice ...int) error {
	slice := reflect.ValueOf(elementsInterface)
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf("A slice input is required for this function expected: Slice, got: %v", slice.Kind())
	}

	return l.Add(elementsInterface, startIndexSlice...)
}

func (l *ListNode) AddFirst(elementsInterface interface{}) error {
	return l.Add(elementsInterface, 0)
}

func (l *ListNode) AddLast(elementsInterface interface{}) error {
	return l.Add(elementsInterface, l.Length)
}

func (l *ListNode) Clear() {
	l.Head = nil
}

func (l *ListNode) Clone() ListNode {
	return *l
}

func (l *ListNode) Contains(element interface{}) bool {
	matchIndex, _ := l.iterateList(false, element, false, false)
	if matchIndex != -1 {
		return true
	}
	return false
}

func (l *ListNode) Element() *node {
	return l.Head
}

func (l *ListNode) Get(index int) (interface{}, error) {
	if index > l.Length {
		return nil, fmt.Errorf("received %v, which is greater than the list size: %v", index, l.Length)
	}
	_, nodeElement := l.iterateList(false, index, false, true)
	return nodeElement.Val, nil
}

func (l *ListNode) GetFirst() (interface{}, error) {
	if l.Head != nil {
		return l.Head.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

func (l *ListNode) GetLast() (interface{}, error) {
	if l.Head != nil {
		return l.Tail.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

// Extra functions in addition to ones described in java documentation
func (l *ListNode) Print(debug ...bool) {
	var shouldDebug bool
	if len(debug) != 0 {
		if debug[0] {
			shouldDebug = false
		}
	}
	l.iterateList(true, nil, shouldDebug, false)
}

// unexported and utility functions used by above functions
func (l *ListNode) getStartingIndex(startIndexSlice ...int) (int, error) {
	if len(startIndexSlice) > 0 {
		if len(startIndexSlice) > 2 {
			return -1, fmt.Errorf("only one argument expected as start index, received %v", startIndexSlice)
		} else {
			return startIndexSlice[0], nil
		}
	}
	return l.Length, nil
}

func (l *ListNode) iterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool, searchByIndex bool) (int, *node) {
	curr := l.Head
	for i := 0; i < l.Length; i++ {
		if shouldPrint {
			fmt.Print((*curr).Val)
			if shouldDebug {
				fmt.Printf("\t Prev: %p \t current: %p \t Next: %p", curr.Prev, curr, curr.Next)
			}
			fmt.Println()
		}
		if searchKey != nil && searchKey == curr.Val {
			return i, nil
		}
		if searchByIndex {
			if i == searchKey {
				return -1, curr
			}
		}
		curr = curr.Next
	}
	return -1, nil
}

// Functions need to looked at
//func (l *ListNode) Get(index int) (*node, error) {
//	if index <= l.Size() {
//		_, _, n := l.iterateList(false, nil, false, false, index)
//		return n, nil
//	}
//	return nil, errors.New("search Index is greater than the list Size")
//}
//
//
//func (l *ListNode) PrintListNode() {
//	l.iterateList(true, nil, false, false)
//}
//
//func (l *ListNode) PrintReverseList() {
//	l.reverseIterateList(true)
//}
//
//func (l *ListNode) Size() int {
//	size, _, _ := l.iterateList(false, nil, false, false)
//	return size
//}
//
//func (l *ListNode) GetFirstMatchIndex(val interface{}) int {
//	matchIndex, _, _ := l.iterateList(false, val, false, false)
//	return matchIndex + 1
//}
//
//func (l *ListNode) isEmpty() bool {
//	return l.Head == nil
//}
//func (l *ListNode) RemoveFirst() (*ListNode, error) {
//	if !l.isEmpty() || l.Size() >= 1 {
//		l.Head = l.Head.Next
//		return l, nil
//	}
//	return l, errors.New("list is empty, or there's only one element")
//}
//
//func (l *ListNode) DebugPrintList() {
//	l.iterateList(true, nil, true, false)
//}
//
//func (l *ListNode) ToSlice() []interface{} {
//	_, slice, _ := l.iterateList(false, nil, false, true)
//	return slice
//}

//func (l *ListNode) Set(index int, element interface{}) []interface{} {
//	newNode := &node{Val: element}
//	if l.Size() < index {
//		if index == 0 {
//
//		} else if index == l.Size() - 1 {
//
//		} else {
//
//		}
//	} else {
//		//return nil
//	}
//}
// RemoveLastOccurrence
// RemoveLast
// RemoveFirstOccurrence

func (l *ListNode) reverseIterateList(shouldPrint bool) {
	curr := l.Tail
	for curr != nil {
		if shouldPrint {
			fmt.Println((*curr).Val)
		}
		curr = curr.Prev
	}
}

// equals
