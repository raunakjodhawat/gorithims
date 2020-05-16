// Package linkedlist contains all functions pertaining to working with linked list in Go. As per java SE 14 (https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/LinkedList.html)
package linkedlist

import (
	"errors"
	"fmt"
	col "github.com/raunakjodhawat/gorithims/src/dataStructure/collection"
	"reflect"
)

type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

// List type, stores the head and tail. It represents a list of address with each list instantiated having a default nil.head and tail

type List struct {
	listElements col.Collection
}

// Add Inserts the specified element at the specified position in this list. By default add to the end of list
func (l *List) Add(element interface{}, startIndexSlice ...int) error {
	return l.listElements.Add(element, startIndexSlice...)
}

// AddAll Inserts all of the elements in the specified collection into this list, starting at the specified position. or by default at the end
func (l *List) AddAll(elementsInterface interface{}, startIndexSlice ...int) error {
	slice := reflect.ValueOf(elementsInterface)
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf("A slice input is required for this function expected: Slice, got: %v", slice.Kind())
	}
	return l.listElements.Add(elementsInterface, startIndexSlice...)
}

// AddFirst Inserts the specified element at the beginning of this list.
func (l *List) AddFirst(elementsInterface interface{}) error {
	return l.listElements.Add(elementsInterface, 0)
}

// AddLast Appends the specified element to the end of this list.
func (l *List) AddLast(elementsInterface interface{}) error {
	return l.listElements.Add(elementsInterface, l.listElements.Size())
}

// Clear Removes all of the elements from this list.
func (l *List) Clear() {
	l.listElements.Head = nil
	l.listElements.Tail = nil
	l.listElements.Length = 0
}

// Clone Returns a shallow copy of this LinkedList.
func (l *List) Clone() List {
	return *l
}

// Contains Returns true if this list contains the specified element.
func (l *List) Contains(element interface{}) bool {
	matchIndex, _ := col.Iterate(&l.listElements, false, element, false, false)
	return matchIndex != -1
}

// Element Retrieves, but does not remove, the head (first element) of this list.
func (l *List) Element() *col.Node {
	return l.listElements.Head
}

// Get Returns the element at the specified position in this list.
func (l *List) Get(index int) (interface{}, error) {
	if index >= l.listElements.Size() {
		return nil, fmt.Errorf("received %v, which is greater than the list size: %v", index, l.listElements.Size())
	}
	_, NodeElement := col.Iterate(&l.listElements, false, index, false, true)
	return NodeElement.Val, nil
}

// GetFirst Returns the first element in this list.
func (l *List) GetFirst() (interface{}, error) {
	if l.listElements.Head != nil {
		return l.listElements.Head.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

// GetLast Returns the last element in this list.
func (l *List) GetLast() (interface{}, error) {
	if l.listElements.Head != nil {
		return l.listElements.Tail.Val, nil
	}
	return -1, errors.New("head is not yet initialized")
}

// IndexOf Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (l *List) IndexOf(searchKey interface{}) int {
	index, _ := col.Iterate(&l.listElements, false, searchKey, false, false)
	return index
}

// LastIndexOf Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (l *List) LastIndexOf(searchKey interface{}) int {
	index, _ := col.ReverseIterate(&l.listElements, false, searchKey, false, false)
	return index
}

// ListIterator Returns a list-iterator of the elements in this list (in proper sequence), starting at the specified position in the list.
func (l *List) ListIterator(index int) (*col.Node, error) {
	if index >= l.listElements.Length {
		return nil, fmt.Errorf("received %v, which is greater than the list size: %v", index, l.Size())
	}
	_, itrPointer := col.Iterate(&l.listElements, false, index, false, true)
	return itrPointer, nil
}

// Offer Adds the specified element as the tail (last element) of this list.
func (l *List) Offer(element interface{}) (bool, error) {
	return l.offerHelper(element)
}

// OfferFirst Inserts the specified element at the front of this list.
func (l *List) OfferFirst(element interface{}) (bool, error) {
	return l.offerHelper(element, 0)
}

// OfferLast Inserts the specified element at the end of this list.
func (l *List) OfferLast(element interface{}) (bool, error) {
	return l.offerHelper(element)
}

// Peek Retrieves, but does not remove, the head (first element) of this list.
func (l *List) Peek() (*col.Node, error) {
	return l.peekHelper(l.listElements.Head)
}

// PeekFirst Retrieves, but does not remove, the first element of this list, or returns null if this list is empty.
func (l *List) PeekFirst() (*col.Node, error) {
	return l.peekHelper(l.listElements.Head)
}

// PeekLast Retrieves, but does not remove, the last element of this list, or returns null if this list is empty.
func (l *List) PeekLast() (*col.Node, error) {
	return l.peekHelper(l.listElements.Tail)
}

// Poll Retrieves and removes the head (first element) of this list.
func (l *List) Poll() (*col.Node, error) {
	return l.listElements.Remove()
}

// PollFirst Retrieves and removes the first element of this list, or returns null if this list is empty.
func (l *List) PollFirst() (*col.Node, error) {
	return l.listElements.Remove(0)
}

// PollLast Retrieves and removes the last element of this list, or returns null if this list is empty.
func (l *List) PollLast() (*col.Node, error) {
	return l.listElements.Remove(l.Size() - 1)
}

// Pop Pops an element from the stack represented by this list.
func (l *List) Pop() (interface{}, error) {
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
func (l *List) Push(element interface{}) error {
	return l.listElements.Add(element, 0)
}

// Remove Retrieves and removes the head (first element) of this list.
func (l *List) Remove(startIndexSlice ...int) (*col.Node, error) {
	return l.listElements.Remove(startIndexSlice...)
}

// RemoveFirst Removes and returns the first element from this list.
func (l *List) RemoveFirst() (*col.Node, error) {
	return l.listElements.Remove()
}

// RemoveLast Removes and returns the last element from this list.
func (l *List) RemoveLast() (*col.Node, error) {
	return l.listElements.Remove(l.Size() - 1)
}

// RemoveFirstOccurrence Removes the first occurrence of the specified element in this list (when traversing the list from head to tail).
func (l *List) RemoveFirstOccurrence(searchKey interface{}) (*col.Node, error) {
	index, _ := col.Iterate(&l.listElements, false, searchKey, false, false)
	if index != -1 {
		return l.listElements.Remove(index)
	}
	return nil, fmt.Errorf("%v, is not present in the list", searchKey)
}

// RemoveLastOccurrence Removes the last occurrence of the specified element in this list (when traversing the list from head to tail).
func (l *List) RemoveLastOccurrence(searchKey interface{}) (*col.Node, error) {
	index, _ := col.ReverseIterate(&l.listElements, false, searchKey, false, false)
	if index != -1 {
		return l.listElements.Remove(index)
	}
	return nil, fmt.Errorf("%v, is not present in the list", searchKey)
}

// Set Replaces the element at the specified position in this list with the specified element.
func (l *List) Set(index int, element interface{}) error {
	if index < l.Size() {
		_, oldElement := col.Iterate(&l.listElements, false, index, false, true)
		oldElement.Val = element
		return nil
	}
	return errors.New("index is greater than the list length")
}

// Size returns the length of the linked list
func (l *List) Size() int {
	return l.listElements.Length
}

// ToArray Returns an array containing all of the elements in this list in proper sequence (from first to last element)
func (l *List) ToArray() []interface{} {
	listToArray := make([]interface{}, l.Size())
	var i int
	curr := l.listElements.Head
	for curr != nil {
		listToArray[i] = curr.Val
		curr = curr.Next
		i++
	}
	return listToArray
}

// Print prints the linked list element. optional debug argument as true or false can be provided
func (l *List) Print(debug ...bool) {
	col.Iterate(&l.listElements, true, nil, l.printHelper(debug...), false)
}

// PrintReverse prints the linked list in reverse order. optional debug argument as true or false can be provided
func (l *List) PrintReverse(debug ...bool) {
	col.ReverseIterate(&l.listElements, true, nil, l.printHelper(debug...), false)
}

// PrintPretty prints the linked list as a slice, so it is easier to debug
func (l *List) PrintPretty() {
	fmt.Println(l.ToArray())
}

// unexported and utility functions used by above functions
func (l *List) getStartingIndex(startIndexSlice ...int) (int, error) {
	if len(startIndexSlice) > 0 {
		if len(startIndexSlice) > 2 {
			return -1, fmt.Errorf("only one argument expected as start index, received %v", startIndexSlice)
		} else if startIndexSlice[0] > l.Size() {
			return startIndexSlice[0], fmt.Errorf("starting index cant be greater than length of list, expected a number less than %v, got %v", l.Size()+1, startIndexSlice[0])
		} else {
			return startIndexSlice[0], nil
		}
	}
	return l.Size(), nil
}

func (l *List) offerHelper(element interface{}, startIndex ...int) (bool, error) {
	err := l.listElements.Add(element, startIndex...)
	if err != nil {
		return false, err
	}
	return true, err
}

func (l *List) peekHelper(searchNode *col.Node) (*col.Node, error) {
	if searchNode != nil {
		return searchNode, nil
	}
	return nil, errors.New("list is not yet initialized")
}

func (l *List) printHelper(debug ...bool) bool {
	var shouldDebug bool
	if len(debug) != 0 {
		if debug[0] {
			shouldDebug = true
		}
	}
	return shouldDebug
}
