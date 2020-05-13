// Package linkedlist package contains all functions pertaining to working with linked list in Go. As per java SE 14 (https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/LinkedList.html)
package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

// Node holds the data element and pointer to next and previous node in the linked list
type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

// ListNode type, stores the head and tail. It represents a list of address with each list instantiated having a default nil.head and tail
type ListNode struct {
	head   *Node
	tail   *Node
	length int
}

// Add Inserts the specified element at the specified position in this list. By default add to the end of list
func (l *ListNode) Add(element interface{}, startIndexSlice ...int) error {

	startIndex, err := l.getStartingIndex(startIndexSlice...)
	if err != nil {
		return err
	}

	slice := reflect.ValueOf(element)
	if slice.Kind() != reflect.Slice {
		n := &Node{Val: element}
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
		l.length++
	} else {
		elementsSlice := make([]interface{}, slice.Len())

		for i := 0; i < slice.Len(); i++ {
			elementsSlice[i] = slice.Index(i).Interface()
		}

		curr := l.head
		if startIndex == 0 {
			for i := 0; i < slice.Len(); i++ {
				err := l.Add(elementsSlice[i], startIndex)
				if err != nil {
					return err
				}
				startIndex++
			}
		} else if startIndex != l.length {
			for i := 0; i < l.length; i++ {
				if i == startIndex-1 {
					nextCopy := curr.Next
					for j := 0; j < slice.Len(); j++ {
						innerNode := &Node{Val: elementsSlice[j]}
						curr.Next = innerNode
						currCopy := curr
						curr = curr.Next
						curr.Prev = currCopy
						l.length++
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
				n := &Node{Val: elementsSlice[i]}
				currTail.Next = n
				currTail = currTail.Next
				currTail.Prev = l.tail
				l.tail = l.tail.Next
				l.length++
			}
		}
	}
	return nil
}

// AddAll Inserts all of the elements in the specified collection into this list, starting at the specified position. or by default at the end
func (l *ListNode) AddAll(elementsInterface interface{}, startIndexSlice ...int) error {
	slice := reflect.ValueOf(elementsInterface)
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf("A slice input is required for this function expected: Slice, got: %v", slice.Kind())
	}
	return l.Add(elementsInterface, startIndexSlice...)
}

// AddFirst Inserts the specified element at the beginning of this list.
func (l *ListNode) AddFirst(elementsInterface interface{}) error {
	return l.Add(elementsInterface, 0)
}

// AddLast Appends the specified element to the end of this list.
func (l *ListNode) AddLast(elementsInterface interface{}) error {
	return l.Add(elementsInterface, l.length)
}

// Clear Removes all of the elements from this list.
func (l *ListNode) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Clone Returns a shallow copy of this LinkedList.
func (l *ListNode) Clone() ListNode {
	return *l
}

// Contains Returns true if this list contains the specified element.
func (l *ListNode) Contains(element interface{}) bool {
	matchIndex, _ := l.iterateList(false, element, false, false)
	return matchIndex != -1
}

// Element Retrieves, but does not remove, the head (first element) of this list.
func (l *ListNode) Element() *Node {
	return l.head
}

// Get Returns the element at the specified position in this list.
func (l *ListNode) Get(index int) (interface{}, error) {
	if index >= l.length {
		return nil, fmt.Errorf("received %v, which is greater than the list size: %v", index, l.length)
	}
	_, NodeElement := l.iterateList(false, index, false, true)
	return NodeElement.Val, nil
}

// GetFirst Returns the first element in this list.
func (l *ListNode) GetFirst() (interface{}, error) {
	if l.head != nil {
		return l.head.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

// GetLast Returns the last element in this list.
func (l *ListNode) GetLast() (interface{}, error) {
	if l.head != nil {
		return l.tail.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

// IndexOf Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (l *ListNode) IndexOf(searchKey interface{}) int {
	index, _ := l.iterateList(false, searchKey, false, false)
	return index
}

// LastIndexOf Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (l *ListNode) LastIndexOf(searchKey interface{}) int {
	index, _ := l.reverseIterateList(false, searchKey, false, false)
	return index
}

// ListIterator Returns a list-iterator of the elements in this list (in proper sequence), starting at the specified position in the list.
func (l *ListNode) ListIterator(index int) (*Node, error) {
	if index >= l.length {
		return nil, fmt.Errorf("received %v, which is greater than the list size: %v", index, l.length)
	}
	_, itrPointer := l.iterateList(false, index, false, true)
	return itrPointer, nil
}

// Offer Adds the specified element as the tail (last element) of this list.
func (l *ListNode) Offer(element interface{}) (bool, error) {
	return l.offerHelper(element)
}

// OfferFirst Inserts the specified element at the front of this list.
func (l *ListNode) OfferFirst(element interface{}) (bool, error) {
	return l.offerHelper(element, 0)
}

// OfferLast Inserts the specified element at the end of this list.
func (l *ListNode) OfferLast(element interface{}) (bool, error) {
	return l.offerHelper(element)
}

// Peek Retrieves, but does not remove, the head (first element) of this list.
func (l *ListNode) Peek() (*Node, error) {
	return l.peekHelper(l.head)
}

// PeekFirst Retrieves, but does not remove, the first element of this list, or returns null if this list is empty.
func (l *ListNode) PeekFirst() (*Node, error) {
	return l.peekHelper(l.head)
}

// PeekLast Retrieves, but does not remove, the last element of this list, or returns null if this list is empty.
func (l *ListNode) PeekLast() (*Node, error) {
	return l.peekHelper(l.tail)
}

// Poll Retrieves and removes the head (first element) of this list.
func (l *ListNode) Poll() (*Node, error) {
	if l.head != nil {
		if l.head.Next == nil {
			return l.head, nil
		}
		headCopy := l.head
		l.head = l.head.Next
		l.head.Prev = nil
		l.length--
		return headCopy, nil
	}
	return nil, errors.New("list is not yet initialized")
}

// PollFirst Retrieves and removes the first element of this list, or returns null if this list is empty.
func (l *ListNode) PollFirst() (*Node, error) {
	return l.Poll()
}

// PollLast Retrieves and removes the last element of this list, or returns null if this list is empty.
func (l *ListNode) PollLast() (*Node, error) {
	if l.head != nil {
		if l.tail.Prev == nil {
			return l.tail, nil
		}
		tailCopy := l.tail
		l.tail = l.tail.Prev
		l.tail.Next = nil
		l.length--
		return tailCopy, nil
	}
	return nil, errors.New("list is not yet initialized")
}

// Pop Pops an element from the stack represented by this list.
func (l *ListNode) Pop() (interface{}, error) {
	currentNode, err := l.Peek()
	if err != nil {
		return nil, err
	}
	_, err = l.Poll()
	if err != nil {
		return nil, err
	}
	return currentNode.Val, nil
}

// Push Pushes an element onto the stack represented by this list.
func (l *ListNode) Push(element interface{}) error {
	return l.Add(element, 0)
}

// Remove Retrieves and removes the head (first element) of this list.
func (l *ListNode) Remove(startIndexSlice ...int) (*Node, error) {
	if l.head == nil {
		return nil, errors.New("list is not initialized")
	}
	if len(startIndexSlice) > 0 {
		if startIndexSlice[0] < l.length {
			if startIndexSlice[0] == 0 {
				return l.Poll()
			}
			if l.length-1 == startIndexSlice[0] {
				return l.PollLast()
			}
			_, element := l.iterateList(false, startIndexSlice[0], false, true)
			prevCopy := element.Prev
			prevCopy.Next = element.Next
			element.Next.Prev = prevCopy
			l.length--
			return element, nil
		}
		return nil, fmt.Errorf("index %v can not be removed, because its greater than list length", startIndexSlice[0])
	}
	return l.Poll()
}

// RemoveFirst Removes and returns the first element from this list.
func (l *ListNode) RemoveFirst() (*Node, error) {
	return l.Remove()
}

// RemoveLast Removes and returns the last element from this list.
func (l *ListNode) RemoveLast() (*Node, error) {
	return l.Remove(l.length - 1)
}

// RemoveFirstOccurrence Removes the first occurrence of the specified element in this list (when traversing the list from head to tail).
func (l *ListNode) RemoveFirstOccurrence(searchKey interface{}) (*Node, error) {
	index, _ := l.iterateList(false, searchKey, false, false)
	if index != -1 {
		return l.Remove(index)
	}
	return nil, fmt.Errorf("%v, is not present in the list", searchKey)
}

// RemoveLastOccurrence Removes the last occurrence of the specified element in this list (when traversing the list from head to tail).
func (l *ListNode) RemoveLastOccurrence(searchKey interface{}) (*Node, error) {
	index, _ := l.reverseIterateList(false, searchKey, false, false)
	if index != -1 {
		return l.Remove(index)
	}
	return nil, fmt.Errorf("%v, is not present in the list", searchKey)
}

// Set Replaces the element at the specified position in this list with the specified element.
func (l *ListNode) Set(index int, element interface{}) error {
	if index < l.length {
		_, oldElement := l.iterateList(false, index, false, true)
		oldElement.Val = element
		return nil
	}
	return errors.New("index is greater than the list length")
}

// Size returns the length of the linked list
func (l *ListNode) Size() int {
	return l.length
}

// ToArray Returns an array containing all of the elements in this list in proper sequence (from first to last element)
func (l *ListNode) ToArray() []interface{} {
	listToArray := make([]interface{}, l.length)
	var i int
	curr := l.head
	for curr != nil {
		listToArray[i] = curr.Val
		curr = curr.Next
		i++
	}
	return listToArray
}

// Print prints the linked list element. optional debug argument as true or false can be provided
func (l *ListNode) Print(debug ...bool) {
	l.iterateList(true, nil, l.printHelper(debug...), false)
}

// PrintReverse prints the linked list in reverse order. optional debug argument as true or false can be provided
func (l *ListNode) PrintReverse(debug ...bool) {
	l.reverseIterateList(true, nil, l.printHelper(debug...), false)
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

func (l *ListNode) iterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool, searchByIndex bool) (int, *Node) {
	curr := l.head
	var i int
	for curr != nil {
		if shouldPrint {
			fmt.Print((*curr).Val)
			if shouldDebug {
				fmt.Printf("\t Index: %v \t Next: %p \t current: %p \t Prev: %p", i, curr.Next, curr, curr.Prev)
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

func (l *ListNode) reverseIterateList(shouldPrint bool, searchKey interface{}, shouldDebug bool, searchByIndex bool) (int, *Node) {
	curr := l.tail
	counter := l.length - 1
	for curr != nil {
		if shouldPrint {
			fmt.Print((*curr).Val)
			if shouldDebug {
				fmt.Printf("\t Index: %v \t Next: %p \t current: %p \t Prev: %p", counter, curr.Next, curr, curr.Prev)
			}
			fmt.Println()
		}
		if searchKey != nil && searchKey == curr.Val && !searchByIndex {
			return counter, nil
		}
		if searchByIndex {
			if counter == searchKey {
				return -1, curr
			}
		}
		curr = curr.Prev
		counter--
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

func (l *ListNode) peekHelper(searchNode *Node) (*Node, error) {
	if searchNode != nil {
		return searchNode, nil
	}
	return nil, errors.New("list is not yet initialized")
}

func (l *ListNode) printHelper(debug ...bool) bool {
	var shouldDebug bool
	if len(debug) != 0 {
		if debug[0] {
			shouldDebug = true
		}
	}
	return shouldDebug
}
