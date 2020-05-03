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

type ListNode struct {
	Head *node
	Tail *node
}

func (l *ListNode) Add(element interface{}) {
	n := &node{Val: element}
	if l.Head == nil {
		l.Head = n
	} else {
		n.Prev = l.Tail
		l.Tail.Next = n
	}
	l.Tail = n
}

func (l *ListNode) AddAll(elementsInterface interface{}, startIndexSlice ...int) error {
	slice := reflect.ValueOf(elementsInterface)
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf("A slice input is required for this function expected: Slice, got: %v", slice.Kind())
	}

	var startIndex int
	listSize := l.Size()
	if len(startIndexSlice) > 0 {
		startIndex = startIndexSlice[0]
		fmt.Println(startIndex, "Start index", listSize)
	} else if len(startIndexSlice) > 1 {
		return fmt.Errorf("start index can only have one element, got: %v", startIndexSlice)
	} else {
		startIndex = listSize - 1
	}

	if startIndex < listSize {
		// convert the interface using reflect to slice of interface
		elementsSlice := make([]interface{}, slice.Len())

		for i := 0; i < slice.Len(); i++ {
			elementsSlice[i] = slice.Index(i).Interface()
		}

		currentNodeAtIndex, err := l.Get(startIndex)
		if err != nil {
			return err
		}
		fmt.Println(currentNodeAtIndex.Val)
		for _, element := range elementsSlice {
			n := &node{Val: element}
			if l.Head == nil {
				l.Head = n
			} else {
				n.Prev = l.Tail
				l.Tail.Next = n
			}
			l.Tail = n
		}
	} else {
		return fmt.Errorf("starting index can not be greater than the list size, expected a starting index less than %v, received %v", listSize - 1, startIndexSlice)
	}
	return nil
}

func (l * ListNode) Get(index int) (*node, error) {
	if index < l.Size() {
		_, _, n := l.iterateList(false, nil, false, false, index)
		return n, nil
	}
	return nil, errors.New("search Index is greater than the list Size")
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

func (l *ListNode) PrintListNode() {
	l.iterateList(true, nil, false, false)
}

func (l *ListNode) PrintReverseList() {
	l.reverseIterateList(true)
}

func (l *ListNode) Size() int {
	size, _ , _:= l.iterateList(false, nil, false, false)
	return size
}

func (l *ListNode) GetFirstMatchIndex(val interface{}) int {
	matchIndex, _, _ := l.iterateList(false, val, false, false)
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

func (l *ListNode) ToSlice() []interface{} {
	_, slice, _ := l.iterateList(false, nil, false, true)
	return slice
}

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

func (l *ListNode) iterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool, returnSlice bool, searchIndex ...int) (int, []interface{}, *node) {
	size := 0
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
			return size - 1, nil, nil
		}
		if len(searchIndex) != 0 {
			if searchIndex[0] == size {
				return searchIndex[0], nil, curr
			}
		}
		curr = curr.Next
		size++
	}
	return size, listToSlice, nil
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
