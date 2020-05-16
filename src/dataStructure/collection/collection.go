// Package collection contains all functions pertaining to working with linked list, Stacks, Queue in GoLang. As per java SE 14 (https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/Collection.html)
package collection

import (
	"errors"
	"fmt"
	"reflect"
)

// Node type holds the data element and pointer to next and previous node in the linked list
type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

// Collection type, stores the Head and Tail. It represents a Collection of address with each Collection instantiated having a default nil.Head and Tail
type Collection struct {
	Head   *Node
	Tail   *Node
	Length int
}

// Add Inserts the specified element at the specified position in this Collection. By default add to the end of Collection
func (c *Collection) Add(element interface{}, startIndexSlice ...int) error {

	startIndex, err := getStartingIndex(c, startIndexSlice...)
	if err != nil {
		return err
	}

	slice := reflect.ValueOf(element)
	if slice.Kind() != reflect.Slice {
		n := &Node{Val: element}
		if c.Head == nil {
			c.Head = n
			c.Tail = n
		} else {
			if startIndex == 0 {
				HeadCopy := c.Head
				c.Head = n
				n.Next = HeadCopy
				HeadCopy.Prev = c.Head
			} else if startIndex != c.Length {
				curr := c.Head
				for i := 0; i < c.Length; i++ {
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
			} else if startIndex == c.Length {
				prevCopy := c.Tail
				prevCopy.Next = n
				n.Prev = prevCopy
				c.Tail = n
			} else {
				c.Tail.Prev = n
			}
		}
		c.Length++
	} else {
		elementsSlice := make([]interface{}, slice.Len())

		for i := 0; i < slice.Len(); i++ {
			elementsSlice[i] = slice.Index(i).Interface()
		}

		curr := c.Head
		if startIndex == 0 {
			for i := 0; i < slice.Len(); i++ {
				err := c.Add(elementsSlice[i], startIndex)
				if err != nil {
					return err
				}
				startIndex++
			}
		} else if startIndex != c.Length {
			for i := 0; i < c.Length; i++ {
				if i == startIndex-1 {
					nextCopy := curr.Next
					for j := 0; j < slice.Len(); j++ {
						innerNode := &Node{Val: elementsSlice[j]}
						curr.Next = innerNode
						currCopy := curr
						curr = curr.Next
						curr.Prev = currCopy
						c.Length++
					}
					curr.Next = nextCopy
					if nextCopy != nil {
						nextCopy.Prev = curr
					}
					if i-1+startIndex == c.Length {
						c.Tail = curr
					}
					break
				}
				curr = curr.Next
			}
		} else {
			currTail := c.Tail
			for i := 0; i < slice.Len(); i++ {
				n := &Node{Val: elementsSlice[i]}
				currTail.Next = n
				currTail = currTail.Next
				currTail.Prev = c.Tail
				c.Tail = c.Tail.Next
				c.Length++
			}
		}
	}
	return nil
}

// AddAll Inserts all of the elements in the specified collection into this Collection, starting at the specified position. or by default at the end
func (c *Collection) AddAll(elementsInterface interface{}, startIndexSlice ...int) error {
	slice := reflect.ValueOf(elementsInterface)
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf("A slice input is required for this function expected: Slice, got: %v", slice.Kind())
	}
	return c.Add(elementsInterface, startIndexSlice...)
}

// Clear Removes all of the elements from this Collection.
func (c *Collection) Clear() {
	c.Head = nil
	c.Tail = nil
	c.Length = 0
}

// Contains Returns true if this Collection contains the specified element.
func (c *Collection) Contains(element interface{}) bool {
	matchIndex, _ := Iterate(c, false, element, false, false)
	return matchIndex != -1
}

// ContainsAll returns true if all the elements provided as input are present in the collection
func (c *Collection) ContainsAll(element interface{}) bool {
	slice := reflect.ValueOf(element)
	if slice.Kind() != reflect.Slice {
		_, elementAtIndex := Iterate(c, false, element, false, false)
		return elementAtIndex != nil
	}

	elementsSlice := make([]interface{}, slice.Len())

	for i := 0; i < slice.Len(); i++ {
		elementsSlice[i] = slice.Index(i).Interface()
	}

	for i := 0; i < slice.Len(); i++ {
		_, elementAtIndex := Iterate(c, false, elementsSlice[i], false, false)
		if elementAtIndex == nil {
			return false
		}
	}
	return true

}

// Equals check if all the element value in one list is equal to another list
func (c *Collection) Equals(newCollection Collection) bool {
	return equalHelper(c, newCollection, false)
}

// IsEmpty checks if a collection is empty
func (c *Collection) IsEmpty() bool {
	return c.Head == nil
}

// Iterator Returns a iterator of the elements in this collection (in proper sequence)
func (c *Collection) Iterator() *Node {
	_, itrPointer := Iterate(c, false, 0, false, true)
	return itrPointer
}

// Remove Retrieves and removes the Head (first element) of this Collection.
func (c *Collection) Remove(startIndexSlice ...int) (*Node, error) {
	if c.Head == nil {
		return nil, errors.New("list is not initialized")
	}
	if len(startIndexSlice) > 0 {
		if startIndexSlice[0] < c.Length {
			if startIndexSlice[0] == 0 {
				return removeFromStartHelper(c)
			}
			if c.Length-1 == startIndexSlice[0] {
				return removeFromEndHelper(c)
			}
			_, element := Iterate(c, false, startIndexSlice[0], false, true)
			prevCopy := element.Prev
			prevCopy.Next = element.Next
			element.Next.Prev = prevCopy
			c.Length--
			return element, nil
		}
		return nil, fmt.Errorf("index %v can not be removed, because its greater than list Length", startIndexSlice[0])
	}
	return removeFromStartHelper(c)
}

// RemoveAll Removes all the values and pointers from the collection
func (c *Collection) RemoveAll() {
	c.Head = nil
	c.Tail = nil
	c.Length = 0
}

// Size returns the Length of the collection
func (c *Collection) Size() int {
	return c.Length
}

// ToArray Returns an array containing all of the elements in this collection in proper sequence (from first to last element)
func (c *Collection) ToArray() []interface{} {
	listToArray := make([]interface{}, c.Length)
	var i int
	curr := c.Head
	for curr != nil {
		listToArray[i] = curr.Val
		curr = curr.Next
		i++
	}
	return listToArray
}

// DeepEquals checks if a collection is made from the clone of a list or not. For two collection to be deepEqual all pointers to next and previous must also be equal
func (c *Collection) DeepEquals(newCollection Collection) bool {
	return equalHelper(c, newCollection, true)
}

// Print prints the collection element. optional debug argument as true or false can be provided
func (c *Collection) Print(debug ...bool) {
	Iterate(c, true, nil, printHelper(debug...), false)
}

// PrintReverse prints the collection in reverse order. optional debug argument as true or false can be provided
func (c *Collection) PrintReverse(debug ...bool) {
	ReverseIterate(c, true, nil, printHelper(debug...), false)
}

// PrintPretty prints the collection as a slice, so it is looks pretty on the screen
func (c *Collection) PrintPretty() {
	fmt.Println(c.ToArray())
}

// Iterate function, iterates over the collection to perform either searching or printing. iteration is sequential
func Iterate(c *Collection, shouldPrint bool, searchKey interface{}, shouldDebug bool, searchByIndex bool) (int, *Node) {
	curr := c.Head
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
			return i, curr
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

// ReverseIterate function, iterates over the collection to perform either searching or printing. iteration is in reverse order
func ReverseIterate(c *Collection, shouldPrint bool, searchKey interface{}, shouldDebug bool, searchByIndex bool) (int, *Node) {
	curr := c.Tail
	counter := c.Length - 1
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

// unexported and utility functions used by above functions
func getStartingIndex(c *Collection, startIndexSlice ...int) (int, error) {
	if len(startIndexSlice) > 0 {
		if len(startIndexSlice) > 2 {
			return -1, fmt.Errorf("only one argument expected as start index, received %v", startIndexSlice)
		} else if startIndexSlice[0] > c.Length {
			return startIndexSlice[0], fmt.Errorf("starting index cant be greater than Length of list, expected a number less than %v, got %v", c.Length+1, startIndexSlice[0])
		} else {
			return startIndexSlice[0], nil
		}
	}
	return c.Length, nil
}

func printHelper(debug ...bool) bool {
	var shouldDebug bool
	if len(debug) != 0 {
		if debug[0] {
			shouldDebug = true
		}
	}
	return shouldDebug
}

func equalHelper(oldCollection *Collection, newCollection Collection, deepEqual bool) bool {
	if oldCollection.Length != newCollection.Length {
		return false
	}

	currHeadOfOldCollection := oldCollection.Head
	currHeadOfNewCollection := newCollection.Head
	for i := 0; i < oldCollection.Length; i++ {
		if currHeadOfOldCollection.Val != currHeadOfNewCollection.Val || (deepEqual && (currHeadOfOldCollection.Next != currHeadOfNewCollection.Next ||
			currHeadOfOldCollection.Prev != currHeadOfNewCollection.Prev)) {
			return false
		}
	}
	return true
}

func removeFromStartHelper(c *Collection) (*Node, error) {
	if c.Head != nil {
		HeadCopy := c.Head
		if c.Head.Next == nil {
			c.Head = nil
			c.Length = 0
			c.Tail = nil
			return HeadCopy, nil
		}
		c.Head = c.Head.Next
		c.Head.Prev = nil
		c.Length--
		return HeadCopy, nil
	}

	return nil, errors.New("list is not yet initialized")
}

func removeFromEndHelper(c *Collection) (*Node, error) {
	if c.Head != nil {
		tailCopy := c.Tail
		if c.Tail.Prev == nil {
			c.Head = nil
			c.Length = 0
			c.Tail = nil
			return tailCopy, nil
		}
		c.Tail = c.Tail.Prev
		c.Tail.Next = nil
		c.Length--
		return tailCopy, nil
	}
	return nil, errors.New("list is not yet initialized")
}
