![Go](https://travis-ci.org/raunakjodhawat/gorithims.svg?branch=master)
[![GoDoc](https://godoc.org/github.com/raunakjodhawat/multisort?status.svg)](https://pkg.go.dev/mod/github.com/raunakjodhawat/gorithims?tab=overview)

# Purpose
This repo aims to provide all the functionality related to algorithms and data structures in go.   
One of the thing that I miss in Go Language is the ability to use wide range of data structure available in Java.  
As people transition from languages like java and javascript they will find using the data structures listed in the following project, helpful.  
Currently, this is a WIP. Contributors are most welcome.  

# Installation
```
go get -d github.com/raunakjodhawat/gorithims/src
```

## Contents
Following project contains implementation for Data Structures in Go. (with java and Javascript alike functions)
This project have implementation for:
- Standard Algorithms
    - Sorting (Currently works just with Integer type. Goal is to make it generic)
        - Insertion
        - Bubble
        - Selection
- Standard Data Structures
    - **Linked List** (Work with any type of data type)
-     func (l *ListNode) Add(element interface{}, startIndexSlice ...int) error
      func (l *ListNode) AddAll(elementsInterface interface{}, startIndexSlice ...int) error
      func (l *ListNode) AddFirst(elementsInterface interface{}) error
      func (l *ListNode) AddLast(elementsInterface interface{}) error
      func (l *ListNode) Clear()
      func (l *ListNode) Clone() ListNode
      func (l *ListNode) Contains(element interface{}) bool
      func (l *ListNode) Element() *Node
      func (l *ListNode) Get(index int) (interface{}, error)
      func (l *ListNode) GetFirst() (interface{}, error)
      func (l *ListNode) GetLast() (interface{}, error)
      func (l *ListNode) IndexOf(searchKey interface{}) int
      func (l *ListNode) LastIndexOf(searchKey interface{}) int
      func (l *ListNode) ListIterator(index int) (*Node, error)
      func (l *ListNode) Offer(element interface{}) (bool, error)
      func (l *ListNode) OfferFirst(element interface{}) (bool, error)
      func (l *ListNode) OfferLast(element interface{}) (bool, error)
      func (l *ListNode) Peek() (*Node, error)
      func (l *ListNode) PeekFirst() (*Node, error)
      func (l *ListNode) PeekLast() (*Node, error)
      func (l *ListNode) Poll() (*Node, error)
      func (l *ListNode) PollFirst() (*Node, error)
      func (l *ListNode) PollLast() (*Node, error)
      func (l *ListNode) Pop() (interface{}, error)
      func (l *ListNode) Print(debug ...bool)
      func (l *ListNode) PrintReverse(debug ...bool)
      func (l *ListNode) Push(element interface{}) error
      func (l *ListNode) Remove(startIndexSlice ...int) (*Node, error)
      func (l *ListNode) RemoveFirst() (*Node, error)
      func (l *ListNode) RemoveFirstOccurrence(searchKey interface{}) (*Node, error)
      func (l *ListNode) RemoveLast() (*Node, error)
      func (l *ListNode) RemoveLastOccurrence(searchKey interface{}) (*Node, error)
      func (l *ListNode) Set(index int, element interface{}) error
      func (l *ListNode) Size() int
      func (l *ListNode) ToArray() []interface{}

- Custom Algorithms
    - multisort  
```cassandraql
func MultiSorted(unsortedSlice interface{}, inputSortKeys []string, ascendingSortOrder []bool) ([]multiSortInterface, error)
func Help() string
```
## Implementations
Implementation for all the functions is given in `main.go` file
 
More data structures and algorithms coming soon.  
Look at the project section, to figure out what's the next release cycle

TL; DR; fun fact: gorithims stands for Go + Algorithms 
