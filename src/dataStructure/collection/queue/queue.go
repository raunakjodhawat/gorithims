/*
Package queue is a collection designed for holding elements prior to processing. Besides queues provide insertion, extraction, and inspection operations.
Each of these methods exists in two forms: one returns an error if the operation fails, the other returns a special value (either null or false, depending on the operation).
https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/Queue.html
*/
package queue

import (
	col "github.com/raunakjodhawat/gorithims/src/dataStructure/collection"
)

// Queue holds all the element of the queue. Using queueNode as a wrapper for collection module
type Queue struct {
	queueNode col.Collection
}

// Add Inserts the specified element into this queue, returning true upon success
func (q *Queue) Add(element interface{}) bool {
	err := q.queueNode.Add(element, 0)
	if err != nil {
		return false
	}
	return true
}

// Element Retrieves, last element at the tail (last element added)
func (q *Queue) Element() *col.Node {
	return q.queueNode.Head
}

// Offer Inserts the specified element into this queue, returning true upon success
func (q *Queue) Offer(element interface{}) bool {
	return q.Add(element)
}

// Peek Retrieves, but does not remove, the head of this queue, or returns null if this queue is empty.
func (q *Queue) Peek() *col.Node {
	return q.queueNode.Head
}

// Poll Retrieves and removes the tail of this queue, or returns null if this queue is empty. it removes the last input element
func (q *Queue) Poll() (*col.Node, error) {
	return q.queueNode.Remove(0)
}

// Remove Retrieves and removes the head of this queue
func (q *Queue) Remove() (*col.Node, error) {
	return q.queueNode.Remove(q.queueNode.Size() - 1)
}

// IsEmpty returns true of the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.queueNode.IsEmpty()
}

// Size returns the size of the list
func (q *Queue) Size() int {
	return q.queueNode.Size()
}

// PrintPretty prints the queue as a slice
func (q *Queue) PrintPretty() {
	q.queueNode.PrintPretty()
}

// Print the queue. with each element on a new line, optional debug parameter can be passed in to print extra information. usage: q.Print(true), q.Print(), q.Print(false)
func (q *Queue) Print(shouldDebug ...bool) {
	q.queueNode.Print(shouldDebug...)
}

// PrintReverse prints the queue in reverse order, with each element on a new line, optional debug parameter can be passed in to print extra information. usage: q.Print(true), q.Print(), q.Print(false)
func (q *Queue) PrintReverse(shouldDebug ...bool) {
	q.queueNode.PrintReverse(shouldDebug...)
}
