package main

import (
	"fmt"
	"github.com/raunakjodhawat/gorithims/src/customAlgorithims/multisorted"
	"github.com/raunakjodhawat/gorithims/src/dataStructure/collection/linkedList"
	"github.com/raunakjodhawat/gorithims/src/dataStructure/collection/queue"
	"github.com/raunakjodhawat/gorithims/src/dataStructure/collection/stack"
	"github.com/raunakjodhawat/gorithims/src/standardAlgorithims/rsort"
)

func main() {
	fmt.Printf("\nWelcome to gorithims! this file executes some the examples of how you can leverage the codebase present on https://github.com/raunakjodhawat/gorithims\n\n")
	fmt.Println("Starting first with the usage of\n##########    Custom Algorithms    ##########")
	multiSortExecution_docker()
	fmt.Println()
	fmt.Println("\nNow the next implementation\n##########    Standard Algorithms    ##########")
	algorithmExecution_docker()
	fmt.Println("\n##########    Standard Data structures usage    ##########")
	dataStructureExecution_docker()
	fmt.Println("########## go to executionExample.go file to look at more concrete implementation of all the functions")
}

func multiSortExecution_docker() {
	fmt.Println("\tYou can use multisort package to sort any JSON objects having a slice as any of its value. You can give in multiple input search parameters or keys and sort provide a sort direction")
	fmt.Printf("\tMultisort, Exports two functions, [MultiSorted() & Help()]\n")
	fmt.Println("\tImagine a type Person defined like this\n\ttype Person struct {\n\t\tName string\n\t\tAge  int\n\t}")
	fmt.Println("Create a slice containing variables of type Person, like: \n\t[ p1 := Person{\n\t\tName: 'Joe',\n\t\tAge:  26\n\t}, p2 := ... ]\nNow, we want to sort this particular slice based on different key pairs and in different direction")

	type Person struct {
		Name string
		Age  int
	}
	p1 := Person{
		Name: "Joe",
		Age:  26,
	}
	p2 := Person{
		Name: "Azin",
		Age:  14,
	}
	p3 := Person{
		Name: "Bold",
		Age:  11,
	}
	p4 := Person{
		Name: "AAND",
		Age:  14,
	}
	multisortExamplePersons := []Person{p1, p2, p3, p4}

	fmt.Println("\nImplementation:\nSorting, first by Name in ascending order and then by Age in Descending order")
	sortKeys := []string{"Name", "Age"}
	ascendingOrder := []bool{true, false}
	sortedSlice, err := multisorted.MultiSorted(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range sortedSlice {
		sortedSlice[i] = sortedSlice[i].(Person)
	}

	fmt.Println("Notice that AAND comes before AZIN when the age is constant")
	fmt.Println(sortedSlice) // [{Joe 26} {AAND 14} {Azin 14} {Bold 11}]
}

func algorithmExecution_docker() {
	func() {
		// Sort Algorithm execution examples
		fmt.Println("\nCurrently, rsort only have one algorithm to perform sorting using multiple different technique. Additionally rsort just works on int slice")

		unsortedSlice := []int{25, 17, 31, 13, 2}

		fmt.Println("to Sort []int{25, 17, 31, 13, 2} using Insertion Sort, use rsort.InsertionSort(unsortedSlice)")
		fmt.Println("Alternatively, you can also use Selection as well as Bubble sort to perform the same sort")
		fmt.Println("If you provide a true argument, while calling the appropriate sort function, sorting will take place in Descending order. default is Ascending")
		fmt.Println("Implementation (Sorting int slice in Descending order):\n\trsort.InsertionSort(unsortedSlice, true)\trsort.BubbleSort(unsortedSlice, true)\trsort.SelectionSort(unsortedSlice, true)")

		sortedSlice := rsort.InsertionSort(unsortedSlice, true) // Can use Bubble as well as Selection sort instead
		fmt.Println("Sorted using Insertion Sort in Descending order", sortedSlice)

		fmt.Println("Implementation (Sorting int slice in Ascending order):\n\trsort.InsertionSort(unsortedSlice)\trsort.BubbleSort(unsortedSlice)\trsort.SelectionSort(unsortedSlice)")
		sortedSlice = rsort.BubbleSort(unsortedSlice) // for ascending pass in false flag or dont pass the second argument
		fmt.Println("Sorted using Bubble Sort in Ascending order", sortedSlice)
	}()
}

func dataStructureExecution_docker() {
	func() {
		fmt.Println("##### 1. Linked List execution")
		fmt.Println("\nAny type can be used to create linked list\n\tlist := linkedlist.List{} creates the linkedList")
		list := linkedlist.List{} // creates a instance of linked list node
		for i := 0; i < 5; i++ {
			err := list.Add(i) // Adds elements to the list
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		fmt.Println("\tAdd element using list.Add() or list.AddAll()")
		err := list.Add(10, 0) // Add 10 at index 0
		if err != nil {
			fmt.Println(err)
			return
		}
		err = list.Add(11) // Add 11 at end
		if err != nil {
			fmt.Println(err)
			return
		}
		err = list.Add(12, 3) // Add 12 at middle of the list
		if err != nil {
			fmt.Println(err)
			return
		}
		err = list.Add(13, 8) // Add 13 at end of the list
		if err != nil {
			fmt.Println(err)
			return
		}
		err = list.Add(14, 10)
		if err != nil {
			fmt.Println(err) // Prints error indicating failure to add 14, because its greater than list length
		}

		err = list.AddAll([]string{"raunak", "jodhawat"}, 0) // Add string slice to end of list. notice how list can handle different types
		if err != nil {
			fmt.Println(err)
			return
		}
		err = list.AddAll([]int{15, 16}, 2) // Add string slice at index 2.
		if err != nil {
			fmt.Println(err)
			return
		}
		err = list.AddAll([]string{"end of list"}) // Add a string slice at end of the list
		if err != nil {
			fmt.Println(err)
			return
		}

		err = list.AddFirst("first element") // Add first element
		if err != nil {
			fmt.Println(err)
			return
		}

		err = list.AddLast(17) // Add last element
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("\tlist.Print() prints the list, one element in a line.\nAlternatively if you want information to previous and next pointer use:\n\tlist.Print(true), true to enable debug")
		// list.Print() // Print all elements of the list

		fmt.Println("\tlist.PrintPretty() can be used to print all the elements of list as a slice, all in one line")
		fmt.Printf("\t")
		list.PrintPretty()

		fmt.Println("Additional function to clone, clear the list are also made available")
		fmt.Println("\tlist.Clone()\tlist.Clear()")
		newList := list.Clone() // Clone the list, make a new copy
		list.Clear()            // Clears the list
		// List.print() will be empty, as the list is cleared
		// newList.Print() will print the cloned list, indicating cloning does work
		list = newList // copy newlist to list

		fmt.Println("Check if a element is present in the list\n\tlist.Contains(17) //returns true\n\tlist.Contains('17') //returns false")
		_ = list.Contains(17) // check if a element is present in the list

		fmt.Println("Getting element by index, has never been easier.\n\tlist.Element() //gets head element\n\tlist.Get(3) //get element at index 3(0-based)\n\tlist.GetFirst() //get head element\n\tlist.GetLast() //get tail element")
		head := list.Element()                                           // Gets the head element
		fmt.Println("Head element:\t\t", head.Next, head.Val, head.Prev) // print head value, next node and previous node

		valueAtIndex, _ := list.Get(1)                     // Get element at index 1
		fmt.Println("Element at index 1:\t", valueAtIndex) // "raunak"
		valueAtIndex, _ = list.GetFirst()                  // Get element at head
		fmt.Println("First element:\t\t", valueAtIndex)    // "first element"
		valueAtIndex, _ = list.GetLast()                   // Get element at tail
		fmt.Println("Last element:\t\t", valueAtIndex)     // "17"

		fmt.Println("You can get a index of a element in the list (0-based), if present. Otherwise it returns -1\n\tlist.IndexOf(17) // prints 15\n\tlist.LastIndexOf('raunak') // prints 15")

		fmt.Println("You can get a iterator starting from the desired index\n\tlist.ListIterator(5)\n\tusage:\n\t\tfor iterator != nil {\n\t\t\tfmt.Println(iterator.Val)\n\t\t\titerator=iterator.Next\n\t\t}\n\t//[10,0,1,12,2,3,4,11,13,end of list,17]")

		fmt.Println("Some other included functions that can be used are\n\tlist.Offer(100) // offer/add element to end of the list\n\tlist.OfferFirst('starting of the list') // offer/add element to start of the list\n\tlist.OfferLast('ending of the list') // offer/add element to end of the list")

		fmt.Println("\tlist.Peek()     // Get front element of the list, but dont remove from the list\n\tlist.PeekLast()  // get last element\n\tlist.PeekFirst() // get first element")

		fmt.Println("\tlist.Poll()     // Get front element of the list, and remove it from the list\n\tlist.PollLast()  // get last element\n\tlist.PollFirst() // get first element")

		fmt.Println("\tlist.Push('hello') // push hello to top of the list\n\tlist.Pop() // gets the value at head, and remove the node from the list")

		fmt.Println("\tlist.Remove(5)    // remove from the list, with index\n\tlist.RemoveFirst() // remove first element (head)\n\tlist.RemoveLast()  // remove last element (tail)")

		fmt.Println("\tlist.RemoveFirstOccurrence('raunak') // search the list for raunak and delete it\n\tlist.RemoveLastOccurrence(13) // remove the last occurrence of 13")

		fmt.Println("\tlist.Size() //returns the size of list\n\tlist.ToArray() // Converts list to slice")

		fmt.Println("Printing the list\n\tlist.PrintReverse(true) // print in reverse with extra debug option\n\tlist.Print() // print without extra debug option\n\t")
	}()

	func() {
		fmt.Println("##### 2. Queue execution")
		fmt.Println("\nAny type can be used to create queue\n\tq := queue.Queue{} creates the Queue")

		q := queue.Queue{} // Initializing a queue
		for i := 0; i < 10; i++ {
			isAdded := q.Add(i) // Adding element, you can also use q.offer(element)
			if !isAdded {
				fmt.Println("unable to add element to the queue")
			}
		}
		fmt.Println("\tq.Add(100) // Add 100 to the queue")
		_ = q.Element() // Get last element added to the queue or you can also use q.peek()
		fmt.Println("\tq.Element() // Get last element added to the queue or you can also use q.peek()")

		_, err := q.Poll() // remove the last added element
		if err != nil {
			fmt.Println("Error in removing last element")
		}
		fmt.Println("\tq.Poll() // remove the last added element")

		_, err = q.Remove() // remove a element from queue following (FIFO pattern)
		if err != nil {
			fmt.Println("Error in removing last element")
		}

		fmt.Println("\tq.Remove() // remove a element from queue following (FIFO pattern)")
		fmt.Println("\tq.IsEmpty() // check if queue is empty")
		fmt.Println("\tq.Size() // Get Q's size")
		fmt.Println("Printing the queue\n\tq.PrintReverse(true) // print in reverse with extra debug option\n\tq.Print() // print without extra debug option\n\t")
		/**
		Printing the Queue
			q.PrintReverse(true) // print in reverse with extra debug option
			q.PrintReverse() // print in reverse without extra debug option
			q.Print(true) // print with extra debug option
			q.Print() // print without extra debug option
			q.PrintPretty() // Print as a slice
		*/
	}()

	func() {
		fmt.Println("##### 3. Stack execution")
		fmt.Println("\nAny type can be used to create stack\n\ts := stack.Stack{} creates the Stack")

		fmt.Println("\ts.Push(100) // pushes 100 to stack\n\ts.Peek() // retrieves 100, but does not remove it\n\ts.Pop() //retrieves 100 and removes from the stack\n\ts.Empty() // prints true as the stack is now empty\n\ts.Search(100) // prints -1, Search, returns 1-based index if element is present in the queue, else returns -1")

		s := stack.Stack{} // Initialize the stack

		fmt.Println("Printing the stack\n\ts.PrintReverse(true) // print in reverse with extra debug option\n\tq.Print() // print without extra debug option\n\ts.PrintPretty() // Print stack as a slice eg: ")
		// Add 0 to 9 element on stack
		for i := 0; i < 10; i++ {
			err := s.Push(i) // Adding element, to the stack
			if err != nil {
				fmt.Println("unable to add element to the stack")
			}
		}
		fmt.Printf("\t")
		s.PrintPretty() // Print stack
		fmt.Println()
		/**
		Printing the Stack
			s.PrintReverse(true) // print in reverse with extra debug option
			s.PrintReverse() // print in reverse without extra debug option
			s.Print(true) // print with extra debug option
			s.Print() // print without extra debug option
			s.PrintPretty() // Print as a slice
		*/
	}()
}
