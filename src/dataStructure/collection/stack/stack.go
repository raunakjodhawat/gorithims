// package stack contains the implementation of all the functions as described according to Java 14 documentation (https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/Stack.html)
package stack

import (
	col "github.com/raunakjodhawat/gorithims/src/dataStructure/collection"
)

// type Stack used as a wrapped to use collection class functions
type Stack struct {
	stackNode col.Collection
}

// Empty returns a boolean indicating if this stack is empty (true) or not (false)
func (s *Stack) Empty() bool {
	return s.stackNode.IsEmpty()
}

// Peek returns the object at the top of this stack without removing it from the stack.
func (s *Stack) Peek() *col.Node {
	return s.stackNode.Head
}

// Pop returns the object at the top of this stack, and a error element while removing it from the stack.
func (s *Stack) Pop() (*col.Node, error) {
	return s.stackNode.Remove()
}

// Push adds the element to the stack and returns the error.
func (s *Stack) Push(element interface{}) error {
	return s.stackNode.Add(element)
}

// Search Returns the 1-based position where an object is on this stack. if element is not present in the list, it returns -1
func (s *Stack) Search(element interface{}) int {
	index, _ := col.Iterate(&s.stackNode, false, element, false, false)
	if index == -1 {
		return index
	}
	return index + 1
}

// Print the stack
func (s *Stack) Print(shouldDebug ...bool) {
	s.stackNode.Print(shouldDebug...)
}

// PrintReverse print the stack in reverse order
func (s *Stack) PrintReverse(shouldDebug ...bool) {
	s.stackNode.PrintReverse(shouldDebug...)
}

// PrintPretty prints the stack as a slice
func (s *Stack) PrintPretty() {
	s.stackNode.PrintPretty()
}
