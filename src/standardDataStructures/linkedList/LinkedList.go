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
ListNode type, stores the head and tail. It represents a list of address with each list instantiated having a default nil.head and tail
*/
type ListNode struct {
	head   *node
	tail   *node
	length int
}

func (l *ListNode) Add(element interface{}, startIndexSlice ...int) error {

	startIndex, err := l.getStartingIndex(startIndexSlice...)
	if err != nil {
		return err
	}

	slice := reflect.ValueOf(element)
	if slice.Kind() != reflect.Slice {
		n := &node{Val: element}
		if l.head == nil {
			l.head = n
			l.tail = n
		} else {
			if startIndex == 0 {
				headCopy := l.head
				l.head = n
				n.Next = headCopy
				headCopy.Prev = l.head
			} else if startIndex != l.length {
				curr := l.head
				for i := 0; i < l.length; i++ {
					if i == startIndex-1 {
						copyCurrent := curr
						copyNext := curr.Next
						curr.Next = n
						n.Prev = copyCurrent
						n.Next = copyNext
						n.Next.Prev = n
						break
					}
					curr = curr.Next
				}
			} else if startIndex == l.length {
				prevCopy := l.tail
				prevCopy.Next = n
				n.Prev = prevCopy
				l.tail = n
			} else {
				l.tail.Prev = n
			}
		}
		l.length += 1
	} else {
		elementsSlice := make([]interface{}, slice.Len())

		for i := 0; i < slice.Len(); i++ {
			elementsSlice[i] = slice.Index(i).Interface()
		}

		curr := l.head
		if startIndex == 0 {
			for i := 0; i < slice.Len(); i++ {
				l.Add(elementsSlice[i], startIndex)
				startIndex += 1
			}
		} else if startIndex != l.length {
			for i := 0; i < l.length; i++ {
				if i == startIndex-1 {
					nextCopy := curr.Next
					for j := 0; j < slice.Len(); j++ {
						innerNode := &node{Val: elementsSlice[j]}
						curr.Next = innerNode
						currCopy := curr
						curr = curr.Next
						curr.Prev = currCopy
						l.length += 1
					}
					curr.Next = nextCopy
					if nextCopy != nil {
						nextCopy.Prev = curr
					}
					if i-1+startIndex == l.length {
						l.tail = curr
					}
					break
				}
				curr = curr.Next
			}
		} else {
			currTail := l.tail
			for i := 0; i < slice.Len(); i++ {
				n := &node{Val: elementsSlice[i]}
				currTail.Next = n
				currTail = currTail.Next
				currTail.Prev = l.tail
				l.tail = l.tail.Next
				l.length += 1
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
	return l.Add(elementsInterface, l.length)
}

func (l *ListNode) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
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
	return l.head
}

func (l *ListNode) Get(index int) (interface{}, error) {
	if index >= l.length {
		return nil, fmt.Errorf("received %v, which is greater than the list size: %v", index, l.length)
	}
	_, nodeElement := l.iterateList(false, index, false, true)
	return nodeElement.Val, nil
}

func (l *ListNode) GetFirst() (interface{}, error) {
	if l.head != nil {
		return l.head.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

func (l *ListNode) GetLast() (interface{}, error) {
	if l.head != nil {
		return l.tail.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

func (l *ListNode) IndexOf(searchKey interface{}) int {
	index, _ := l.iterateList(false, searchKey, false, false)
	return index
}

func (l *ListNode) LastIndexOf(searchKey interface{}) int {
	index, _ := l.reverseIterateList(false, searchKey, false, false)
	return index
}

func (l *ListNode) ListIterator(index int) (*node, error) {
	if index >= l.length {
		return nil, fmt.Errorf("received %v, which is greater than the list size: %v", index, l.length)
	}
	_, itrPointer := l.iterateList(false, index, false, true)
	return itrPointer, nil
}

func (l *ListNode) Offer(element interface{}) (bool, error) {
	return l.offerHelper(element)
}

func (l *ListNode) OfferFirst(element interface{}) (bool, error) {
	return l.offerHelper(element, 0)
}

func (l *ListNode) OfferLast(element interface{}) (bool, error) {
	return l.offerHelper(element)
}

func (l *ListNode) Peek() (*node, error) {
	return l.peekHelper(l.head)
}

func (l *ListNode) PeekFirst() (*node, error) {
	return l.peekHelper(l.head)
}

func (l *ListNode) PeekLast() (*node, error) {
	return l.peekHelper(l.tail)
}

func (l *ListNode) Poll() (*node, error) {
	if l.head != nil {
		if l.head.Next == nil {
			return nil, nil
		} else {
			headCopy := l.head
			l.head = l.head.Next
			l.head.Prev = nil
			l.length -= 1
			return headCopy, nil
		}
	}
	return nil, errors.New("list is not yet initialized")
}

func (l *ListNode) PollFirst() (*node, error) {
	return l.Poll()
}

func (l *ListNode) PollLast() (*node, error) {
	if l.head != nil {
		if l.tail.Prev == nil {
			return nil, nil
		} else {
			tailCopy := l.tail
			l.tail = l.tail.Prev
			l.tail.Next = nil
			l.length -= 1
			return tailCopy, nil
		}
	}
	return nil, errors.New("list is not yet initialized")
}

func (l *ListNode) Pop() (interface{}, error) {
	currentNode, err := l.Peek()
	if err != nil {
		return nil, err
	}
	l.Poll()
	return currentNode.Val, nil
}

func (l *ListNode) Push(element interface{}) error {
	return l.Add(element, 0)
}

func (l *ListNode) Remove(startIndexSlice ...int) (*node, error) {
	if l.head == nil {
		return nil, errors.New("list is not initialized")
	}
	if len(startIndexSlice) > 0 {
		if startIndexSlice[0] < l.length {
			if startIndexSlice[0] == 0 {
				return l.Poll()
			} else {
				_, element := l.iterateList(false, startIndexSlice[0], false, true)
				element.Prev.Next = element.Next
				l.length -= 1
				return element, nil
			}
		} else {
			return nil, fmt.Errorf("index %v can not be removed, because its greater than list length", startIndexSlice[0])
		}
	}
	return l.Poll()
}

func (l *ListNode) RemoveFirst() (*node, error) {
	return l.Remove()
}

func (l *ListNode) RemoveLast() (*node, error) {
	return l.Remove(l.length - 1)
}

func (l *ListNode) RemoveLFirstOccurrence(searchKey interface{}) (*node, error) {
	index, _ := l.iterateList(false, searchKey, false, false)
	if index != -1 {
		return l.Remove(index)
	}
	return nil, fmt.Errorf("%v, is not present in the list", searchKey)
}

func (l *ListNode) RemoveLastOccurrence(searchKey interface{}) (*node, error) {
	index, _ := l.reverseIterateList(false, searchKey, false, false)
	if index != -1 {
		return l.Remove(index)
	}
	return nil, fmt.Errorf("%v, is not present in the list", searchKey)
}

func (l *ListNode) Set(index int, element interface{}) error {
	if index < l.length {
		_, oldElement := l.iterateList(false, index, false, true)
		oldElement.Val = element
		return nil
	}
	return errors.New("index is greater than the list length")
}

func (l *ListNode) Size() int {
	return l.length
}

func (l *ListNode) ToArray() []interface{} {
	listToArray := make([]interface{}, l.length)
	var i int
	curr := l.head
	for curr != nil {
		listToArray[i] = curr.Val
		curr = curr.Next
		i += 1
	}
	return listToArray
}

// Extra functions in addition to ones described in java documentation
func (l *ListNode) Print(debug ...bool) {
	var shouldDebug bool
	if len(debug) != 0 {
		if debug[0] {
			shouldDebug = true
		}
	}
	l.iterateList(true, nil, shouldDebug, false)
}

// unexported and utility functions used by above functions
func (l *ListNode) getStartingIndex(startIndexSlice ...int) (int, error) {
	if len(startIndexSlice) > 0 {
		if len(startIndexSlice) > 2 {
			return -1, fmt.Errorf("only one argument expected as start index, received %v", startIndexSlice)
		} else if startIndexSlice[0] > l.length {
			return startIndexSlice[0], fmt.Errorf("starting index cant be greater than length of list, expected a number less than %v, got %v", l.length+1, startIndexSlice[0])
		} else {
			return startIndexSlice[0], nil
		}
	}
	return l.length, nil
}

func (l *ListNode) iterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool, searchByIndex bool) (int, *node) {
	curr := l.head
	var i int
	for curr != nil {
		if shouldPrint {
			fmt.Print((*curr).Val)
			if shouldDebug {
				fmt.Printf("\t Prev: %p \t current: %p \t Next: %p", curr.Prev, curr, curr.Next)
			}
			fmt.Println()
		}
		if searchKey != nil && searchKey == curr.Val && !searchByIndex {
			return i, nil
		}
		if searchByIndex {
			if i == searchKey {
				return -1, curr
			}
		}
		curr = curr.Next
		i++
	}
	return -1, nil
}

// TODO
func (l *ListNode) Sort() *node {
	return l.head
}

func (l *ListNode) reverseIterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool, searchByIndex bool) (int, *node) {
	curr := l.tail
	counter := l.length - 1
	for curr != nil {
		if shouldPrint {
			fmt.Print((*curr).Val)
			if shouldDebug {
				fmt.Printf("\t Prev: %p \t current: %p \t Next: %p", curr.Prev, curr, curr.Next)
			}
			fmt.Println()
		}
		if searchKey != nil && searchKey == curr.Val {
			return counter, nil
		}
		if searchByIndex {
			if counter == searchKey {
				return -1, curr
			}
		}
		curr = curr.Prev
		counter -= 1
	}
	return -1, nil
}

func (l *ListNode) offerHelper(element interface{}, startIndex ...int) (bool, error) {
	err := l.Add(element, startIndex...)
	if err != nil {
		return false, err
	}
	return true, err
}

func (l *ListNode) peekHelper(searchNode *node) (*node, error) {
	if searchNode != nil {
		return searchNode, nil
	} else {
		return nil, errors.New("list is not yet initialized")
	}
}
