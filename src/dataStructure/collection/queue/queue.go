//nolint
/**
Package queue is a collection designed for holding elements prior to processing. Besides queues provide insertion, extraction, and inspection operations.
Each of these methods exists in two forms: one returns an error if the operation fails, the other returns a special value (either null or false, depending on the operation).
*/
package queue

import (
	"fmt"
	linkedlist "github.com/raunakjodhawat/gorithims/src/dataStructure/collection/linkedList"
)

type Queue struct {
	list   *linkedlist.List
	length int
}

func (q *Queue) Add(element interface{}) error {
	if q.list == nil {
		q.list = &linkedlist.List{}
	}
	return q.list.Add(element, 0)
}

func (q *Queue) Remove() (*linkedlist.Node, error) {
	_, err := q.list.RemoveLast()
	return nil, err
}

func (q *Queue) Print() {
	slice := q.list.ToArray()
	fmt.Println(slice)
}
