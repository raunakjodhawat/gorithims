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
			for i:=0; i<slice.Len(); i++ {
				l.Add(elementsSlice[i], startIndex)
				startIndex += 1
			}
		} else {
			for i := 0; i < l.Length; i++ {
				if i == startIndex-1 {
					currCopy := curr.Next
					for j:=0; j<slice.Len(); j++ {
						innerNode := &node{Val: elementsSlice[j]}
						curr.Next = innerNode
						curr = curr.Next
						l.Length += 1
					}
					curr.Next = currCopy
					break
				}
				curr = curr.Next
			}
		}
	}
	return nil
}

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

//func (l *ListNode) AddAll(elementsInterface interface{}, startIndexSlice ...int) error {
//	slice := reflect.ValueOf(elementsInterface)
//	if slice.Kind() != reflect.Slice {
//		return fmt.Errorf("A slice input is required for this function expected: Slice, got: %v", slice.Kind())
//	}
//
//	startIndex, err := l.startIndexChecker(startIndexSlice...)
//	if err != nil || startIndex != -1 {
//		return err
//	}
//	listSize := l.Size()
//
//	// convert the interface using reflect to slice of interface
//	elementsSlice := make([]interface{}, slice.Len())
//
//	for i := 0; i < slice.Len(); i++ {
//		elementsSlice[i] = slice.Index(i).Interface()
//	}
//
//	currentNodeAtIndex, err := l.Get(startIndex)
//	if err != nil {
//		return err
//	}
//
//	if startIndex <= listSize  && startIndex != 0 {
//		copyNext := currentNodeAtIndex
//		var currentNode *node
//		if startIndex == listSize {
//			currentNode = currentNodeAtIndex
//		} else {
//			currentNode = currentNodeAtIndex.Prev
//		}
//		for _, element := range elementsSlice {
//			n := &node{Val: element}
//			currentNode.Next = n
//			currentNode.Prev = currentNode
//			currentNode = currentNode.Next
//		}
//		if startIndex != listSize {
//			currentNode.Next = copyNext
//		}
//	} else if startIndex == 0 {
//		if l.Head == nil {
//			for _, element := range elementsSlice {
//				l.Add(element)
//			}
//		} else {
//			headCopy := l.Head
//			for i := len(elementsSlice) - 1; i >=0 ; i-- {
//				n := &node{Val: elementsSlice[i]}
//				headCopy.Prev = n
//				cpy := headCopy
//				headCopy = headCopy.Prev
//				headCopy.Next = cpy
//			}
//			l.Head = headCopy
//		}
//
//	} else {
//		return fmt.Errorf("starting index can not be greater than the list size, expected a starting index less than %v, received %v", listSize, startIndexSlice)
//	}
//	return nil
//}

func (l *ListNode) Clear() {
	l.Head = nil
}

func (l *ListNode) Get(index int) (*node, error) {
	if index <= l.Size() {
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
	size, _, _ := l.iterateList(false, nil, false, false)
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
	if len(searchIndex) != 0 {
		if searchIndex[0] == size {
			return searchIndex[0], nil, l.Tail
		}
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

// equals
