package main

// Add collections sort method
import (
	"fmt"
	linkedlist "github.com/raunakjodhawat/gorithims/src/standardDataStructures/linkedList"
)

func main() {
	//fmt.Println("Starting Custom Algorithm usage")
	//multiSortExecution()
	//fmt.Println("Standard Algorithm usage")
	//algorithmExecution()
	//fmt.Println("Standard Data structures usage")
	//dataStructureExecution()
	list := linkedlist.ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		list.Add(i * 10) // Adds elements to the list
	}

	list.OfferFirst(20)
	list.OfferFirst(100)
	list.Push(111)

	list.Print()
	fmt.Println()
	list.Set(2, 1029)
	fmt.Println(list.Size())
	list.Print()
	fmt.Println(list.ToArray())

	// Test poll, pollFirst & pollLast
	//fmt.Println(a.Val)
	//l2 := list.Clone()
	//l2.Clear()
	//fmt.Println("l2 is cleared")
	//list.PrintListNode()
	//fmt.Println("le")
	//l2.PrintListNode()
	//
	//l2.Add(100, 0)
	//l2.Add(121, 0)
	//l2.Add(103, 0)

	//list.PrintListNode()
	//list.Print(true)
	//fmt.Println()
	// a := []int{99, 89, 79, 69}
	//list.Add(99, 0) // Not working with 0 and 5
	//list.PrintListNode()
}

//func multiSortExecution() {
//	fmt.Printf("Multisort, Exports two functions \t MultiSorted() and \tHelp()\n")
//	type Person struct {
//		Name string
//		Age  int
//	}
//	p1 := Person{
//		Name: "Joe",
//		Age:  26,
//	}
//	p2 := Person{
//		Name: "Azin",
//		Age:  14,
//	}
//	p3 := Person{
//		Name: "Bold",
//		Age:  11,
//	}
//	p4 := Person{
//		Name: "AAND",
//		Age:  14,
//	}
//	multisortExamplePersons := []Person{p1, p2, p3, p4}
//
//	// Sort first by Name in ascending order and age in descending order
//	sortKeys := []string{"Name", "Age"}
//	ascendingOrder := []bool{true, false}
//	sortedSlice, err := multisorted.MultiSorted(multisortExamplePersons, sortKeys, ascendingOrder)
//	if err != nil {
//		return
//	}
//	for i := range sortedSlice {
//		sortedSlice[i] = sortedSlice[i].(Person)
//	}
//	// Notice that AAND comes before AZIN when the age is constant
//	fmt.Println(sortedSlice) // will print: [{Joe 26} {AAND 14} {Azin 14} {Bold 11}]
//}
//
//func algorithmExecution() {
//	func() {
//		// Sort Algorithm execution examples
//		fmt.Println("Warn: rsort only support int slices")
//		unsortedSlice := []int{25, 17, 31, 13, 2}
//		sortedSlice := rsort.InsertionSort(unsortedSlice, true) // Can use Bubble as well as Selection sort instead
//		fmt.Println("Sorted in Descending order", sortedSlice)
//
//		sortedSlice = rsort.InsertionSort(unsortedSlice) // for ascending pass in false flag or dont pass the second argument
//		fmt.Println("Sorted in Ascending order", sortedSlice)
//	}()
//}
//
//func dataStructureExecution() {
//	func (){
//		fmt.Println("Linked List execution")
//		fmt.Println("Any type can be used to create linked list")
//		list := linkedlist.ListNode{} // creates a instance of linked list node
//		for i := 0; i <5; i++ {
//			list.Push(i) // Adds elements to the list
//		}
//		lastElement, err := list.Last()
//		if err != nil {
//			fmt.Println(lastElement)
//		}
//
//		firstElement, err := list.First()
//		if err != nil {
//			fmt.Println(firstElement)
//		}
//
//		list.PrintListNode() // Prints all elements in the node
//		list.PrintReverseList() // Prints list in reverse
//		list.DebugPrintList() // Prints pointer address to next and previous nodes
//
//		_, _ = list.RemoveFirst() // removes first node
//		fmt.Println(list.Size()) // Prints length of the list
//
//		fmt.Println(list.GetFirstMatchIndex(3)) // returns the index of a key in the list
//
//		fmt.Println(list.ToSlice()) // Convert list to slice
//	}()
//}
