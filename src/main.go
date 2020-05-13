package main

import (
	"fmt"
	"github.com/raunakjodhawat/gorithims/src/customAlgorithims/multisorted"
	rsort "github.com/raunakjodhawat/gorithims/src/standardAlgorithims/sort"
	linkedlist "github.com/raunakjodhawat/gorithims/src/standardDataStructures/linkedList"
)

func main() {
	fmt.Println("Starting Custom Algorithm usage")
	multiSortExecution()
	fmt.Println("Standard Algorithm usage")
	algorithmExecution()
	fmt.Println("Standard Data structures usage")
	dataStructureExecution()
}

func multiSortExecution() {
	fmt.Printf("Multisort, Exports two functions \t MultiSorted() and \tHelp()\n")
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

	// Sort first by Name in ascending order and age in descending order
	sortKeys := []string{"Name", "Age"}
	ascendingOrder := []bool{true, false}
	sortedSlice, err := multisorted.MultiSorted(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range sortedSlice {
		sortedSlice[i] = sortedSlice[i].(Person)
	}
	// Notice that AAND comes before AZIN when the age is constant
	fmt.Println(sortedSlice) // will print: [{Joe 26} {AAND 14} {Azin 14} {Bold 11}]
}

func algorithmExecution() {
	func() {
		// Sort Algorithm execution examples
		fmt.Println("rsort only support int slices")
		unsortedSlice := []int{25, 17, 31, 13, 2}
		sortedSlice := rsort.InsertionSort(unsortedSlice, true) // Can use Bubble as well as Selection sort instead
		fmt.Println("Sorted in Descending order", sortedSlice)

		sortedSlice = rsort.InsertionSort(unsortedSlice) // for ascending pass in false flag or dont pass the second argument
		fmt.Println("Sorted in Ascending order", sortedSlice)
	}()
}

func dataStructureExecution() {
	func() {
		fmt.Println("Linked List execution")
		fmt.Println("Any type can be used to create linked list")
		list := linkedlist.ListNode{} // creates a instance of linked list node
		for i := 0; i < 5; i++ {
			list.Add(i) // Adds elements to the list
		}
		list.Add(10, 0)                                // Add 10 at index 0
		list.Add(11, 3)                                // Add 11 at index 3
		list.AddAll([]string{"raunak", "jodhawat"}, 4) // Add string slice at index 4. notice how list can handle different types
		list.AddAll([]string{"RJ", "jodhawat"})        // Add string slice be default at last index. notice how list can handle different types
		list.AddFirst("First element")                 // Add element at the head
		//list.AddLast("Last element") // Add element at the tail
		list.Print() // Print all nodes

	}()
}
