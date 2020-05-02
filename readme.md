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
As of April 28, 2020. This project have implementation for:
- Algorithms
    - Sorting (Currently works just with Integer type. Goal is to make it generic)
        - Insertion
        - Bubble
        - Selection
- Data Structures
    - Linked List (Work with any type of data type)
        - First()
        - Last()
        - Push()
        - Print()
        - Length()
        - Search()
        - Remove()
        
More data structures and algorithms coming soon.

## Data Structures
1. Linked List  
Supports any input data alter type 
```cassandraql
package main
## Import package
import "github.com/raunakjodhawat/gorithims/src/algos"

## Make a Object of type ListNode from algos package
list := algos.ListNode{}

## Push elements
list.Push(0)
list.Push("Hello world")
list.Push(true)
list.Push("Anything, any data type can be pushed into this list")
list.Push("Literally any type")

## Usage
list.First() // = 0, nil (returns actual element, error)
list.Last() // = Literally any type, nil (returns actual element, error)
list.Length() // = 4 (returns int)
list.GetFirstMatchIndex(true) // = 2 (returns int) 
list.RemoveFirst() // removes first element of the list
list.PrintListNode() // Prints the list

```

## Algorithms
1. Sorting
```
Note: Currently only supports []int type.
Supports the optional functional parameter, reverse
See the definition below in section Optional Parameters to understand more
```
Implements following sorting functions  
Insertion Sort
```
    import "github.com/raunakjodhawat/gorithims/src/sort"
    sort.InsertionSort(input []int, reverse ...bool)
    
    /*  
        Returns and print a sorted int array using insertion sort
    */
```

2. Bubble Sort  
Bubble Sort is internally implemented in recursive
```
    import "github.com/raunakjodhawat/gorithims/src/sort"
    sort.BubbleSort(input []int, reverse ...bool)
    
    /*  
        Returns and print a sorted int array using Bubble sort
    */
```
3. Selection Sort => selection()
```
    import "github.com/raunakjodhawat/gorithims/src/sort"
    sort.SelectionSort(input []int, reverse ...bool)
    
    /*  
        Returns and print a sorted int array using Selection sort
    */
```
4. Merge Sort => merge()
5. Heap Sort => heap()
6. Quick Sort => quick()
7. Bucket Sort => bucket()
8. Radix Sort => radix()

# Functions
## Sorting Functions
- Insertion()
- Bubble()
- BinaryInsertionSort()
- Selection()
- Merge()
- Heap()
- Quick()
- Bucket()
- Radix()

###### Optional Parameters
```
    toPrint:
    for functions with input parameter as (toPrint ...bool)
    toPrint is a boolean value, that is optional
        if true is provided, it will print to console
        if nothing (or false) is provided, it won't print
        regardless it sends the actual raw data in multi-dimensional string array

    reverse:
    for functions with input parameter as (reverse ...bool)
    reverse is a boolean value, that is optional
        if true is provided, it will sort in descending order
        if nothing (or false) is provided, it will sort in ascending order
        regardless it sends the actual sorted data
```
## Complexity functions
1. GetAllComplexity()
```
    import "github.com/raunakjodhawat/gorithims/src/utility"
    Utility.GetAllComplexity(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of
        Time and Space Complexity of sorting algorithims 
        defined for this repo
    */
```
2. GetAllTimeComplexity()
```
    import "github.com/raunakjodhawat/gorithims/src/utility"
    Utility.GetAllTimeComplexity(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of all
        Time Complexity of sorting algorithims 
        defined for this repo
    */
```
3. GetAllSpaceComplexity()
```
    import "github.com/raunakjodhawat/gorithims/src/utility"
    Utility.GetAllTimeComplexity(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of all
        Space Complexity of sorting algorithims 
        defined for this repo
    */
```
4. GetTimeComplexityByKey()
5. GetSpaceComplexityByKey()
6. GetAllComplexityByKey()
7. GetAllFunc()
#### Other functions
1. GetCredits()
```
    import "github.com/raunakjodhawat/gorithims/src/utility"
    Utility.GetCredits(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of References
        used in creation of this project
    */
```

## External Dependency
1. fmt
2. time
3. reflect
4. testing

TL; DR; fun fact: gorithims stands for Go + Algorithms