package main

import (
	"fmt"
	"github.com/raunakjodhawat/gorithims/src/customAlgorithims/multisorted"
	"github.com/raunakjodhawat/gorithims/src/dataStructure/collection/linkedList"
	"github.com/raunakjodhawat/gorithims/src/dataStructure/collection/queue"
	"github.com/raunakjodhawat/gorithims/src/standardAlgorithims/rsort"
)

func main() {
	fmt.Printf("\nWelcome to gorithims! this file executes some the examples of how you can leverage the codebase present on https://github.com/raunakjodhawat/gorithims\n\n")
	fmt.Println("Starting first with the usage of\n##########    Custom Algorithms    ##########")
	multiSortExecution()
	fmt.Println()
	fmt.Println("\nNow the next implementation\n##########    Standard Algorithms    ##########")
	algorithmExecution()
	fmt.Println("Standard Data structures usage")
	dataStructureExecution()
}

func multiSortExecution() {
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

func algorithmExecution() {
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

func dataStructureExecution() {
	func() {
		fmt.Println("Linked List execution")
		fmt.Println("Any type can be used to create linked list")

		list := linkedlist.List{} // creates a instance of linked list node
		for i := 0; i < 5; i++ {
			err := list.Add(i) // Adds elements to the list
			if err != nil {
				fmt.Println(err)
				break
			}
		}
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

		list.Print() // Print all elements of the list

		newList := list.Clone() // Clone the list, make a new copy
		list.Clear()            // Clears the list
		fmt.Println("Does not print anything on the next line, as the list is cleared")
		list.Print() // prints nothing, as the list is cleared above
		fmt.Println("Print the cloned list, indicating cloning does work")
		newList.Print() // Print the new list
		list = newList  // copy newlist to list
		fmt.Println("Print the list, indicating its no more empty")
		list.Print()

		isPresent := list.Contains(17)            // check if a element is present in the list
		fmt.Println(isPresent)                    // true, as 17 is present in the list
		fmt.Println(list.Contains("17"))          // false
		fmt.Println(list.Contains("end of list")) // true

		head := list.Element()                      // Gets the head element
		fmt.Println(head.Next, head.Val, head.Prev) // print head value, next node and previous node

		valueAtIndex, _ := list.Get(1)    // Get element at index 1
		fmt.Println(valueAtIndex)         // "raunak"
		valueAtIndex, _ = list.GetFirst() // Get element at head
		fmt.Println(valueAtIndex)         // "first element"
		valueAtIndex, _ = list.GetLast()  // Get element at tail
		fmt.Println(valueAtIndex)         // "17"

		index := list.IndexOf(17) // get index of 17 in the list
		fmt.Println(index)        // 15
		index = list.LastIndexOf("raunak")
		fmt.Println(index) // 1
		err = list.Add("raunak")
		if err != nil {
			fmt.Println(err)
			return
		}
		index = list.LastIndexOf("raunak")
		fmt.Println(index) // 16

		iterator, _ := list.ListIterator(0) // Get a iterator for the list, starting from head
		for iterator != nil {               // traverse the list
			fmt.Printf("%v\t", iterator.Val) // first element   raunak  jodhawat        15      16      10      0       1       12      2       3       4       11      13      end of list     17      raunak
			iterator = iterator.Next
		}
		fmt.Println()
		iterator, _ = list.ListIterator(10) // Get a iterator for the list, starting from 10th index
		for iterator != nil {               // traverse the list
			fmt.Printf("%v\t", iterator.Val) // 3       4       11      13      end of list     17      raunak
			iterator = iterator.Next
		}
		fmt.Println()

		_, err = list.Offer(100) // offer/add element to end of the list
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = list.OfferFirst("starting of the list") // offer/add element to start of the list
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = list.OfferLast("ending of the list") // offer/add element to end of the list
		if err != nil {
			fmt.Println(err)
			return
		}

		peekNode, _ := list.Peek()     // Get front element of the list, but dont remove from the list
		fmt.Println(peekNode)          // &{starting of the list <Next address> <nil>}
		peekNode, _ = list.PeekLast()  // get last element
		fmt.Println(peekNode)          // &{ending of the list <nil> <Prev address>}
		peekNode, _ = list.PeekFirst() // get first element
		fmt.Println(peekNode)          // &{starting of the list <Next address> <nil>}

		fmt.Println()
		pollNode, _ := list.Poll()     // Get front element of the list, and remove it from the list
		fmt.Println(pollNode)          // &{starting of the list <Next address> <nil>}
		pollNode, _ = list.PollLast()  // get last element
		fmt.Println(pollNode)          // &{ending of the list <nil> <Prev address>}
		pollNode, _ = list.PollFirst() // get first element
		fmt.Println(pollNode)          // &{first element <Next address> <nil>}

		fmt.Println()
		err = list.Push("hello") // push hello to top of the list
		if err != nil {
			fmt.Println(err)
			return
		}
		topElement, _ := list.Pop() // gets the value at head, and remove the node from the list
		fmt.Println(topElement)

		fmt.Println()
		removedElement, _ := list.Remove(0)    // remove from the list, with index
		fmt.Println(removedElement)            // &{raunak <Next address> <nil>}
		removedElement, _ = list.Remove(5)     // remove from the list, element at index 5
		fmt.Println(removedElement)            // &{1 <Next address> <previous address>}
		removedElement, _ = list.RemoveFirst() // remove first element (head)
		fmt.Println(removedElement)            // &{jodhawat <Next address> <nil>}
		removedElement, _ = list.RemoveLast()  // remove last element (tail)
		fmt.Println(removedElement)            // &{100 <nil> <previous address>}

		_, err = list.RemoveFirstOccurrence("raunak") // search the list for "raunak" and delete it
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = list.RemoveFirstOccurrence("hello") // search the list for "hello" and delete it
		if err != nil {
			fmt.Println(err) // Since the list does not have "hello", an error is returned indicating the same
			return
		}
		_, err = list.RemoveLastOccurrence(13) // remove the last occurrence of 13
		if err != nil {
			fmt.Println(err)
			return
		}
		elementAtIndex, _ := list.Get(10)  // Get element at 10th index
		fmt.Println(elementAtIndex)        // 17
		err = list.Set(10, "Hello, world") // set element at 10th index to be hello, world
		if err != nil {
			fmt.Println(err)
			return
		}
		elementAtIndex, _ = list.Get(10)
		fmt.Println(elementAtIndex) // Hello, world

		listSize := list.Size() // gets the size of the list
		fmt.Println(listSize)   // 11

		listToSlice := list.ToArray() // Converts list to slice
		fmt.Println(listToSlice)      // [15 16 10 0 12 2 3 4 11 end of list Hello, world]

		/**
		Printing the list
			list.PrintReverse(true) // print in reverse with extra debug option
			list.PrintReverse() // print in reverse without extra debug option
			list.Print(true) // print with extra debug option
			list.Print() // print without extra debug option
			list.PrintPretty() // Print as a slice
		*/
		list.Print(true)
	}()

	func() {
		fmt.Println("Queue execution")
		fmt.Println("Any data type can be used to create Queue")

		q := queue.Queue{} // Initializing a queue
		for i := 0; i < 10; i++ {
			isAdded := q.Add(i) // Adding element, you can also use q.offer(element)
			if !isAdded {
				fmt.Println("unable to add element to the queue")
			}
		}

		lastElementAdded := q.Element() // Get last element added to the queue or you can also use q.peek()
		fmt.Println(lastElementAdded)

		lastElementAdded, err := q.Poll() // remove the last added element
		if err != nil {
			fmt.Println("Error in removing last element")
		}
		fmt.Println(lastElementAdded)

		dequeueElement, err := q.Remove() // remove a element from queue following (FIFO pattern)
		if err != nil {
			fmt.Println("Error in removing last element")
		}
		fmt.Println(dequeueElement)

		fmt.Println(q.IsEmpty()) // check if queue is empty

		fmt.Println(q.Size()) // Get Q's size
		/**
		Printing the Queue
			q.PrintReverse(true) // print in reverse with extra debug option
			q.PrintReverse() // print in reverse without extra debug option
			q.Print(true) // print with extra debug option
			q.Print() // print without extra debug option
			q.PrintPretty() // Print as a slice
		*/
	}()
}
