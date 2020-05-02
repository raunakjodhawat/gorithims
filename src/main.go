package main

import (
	"fmt"
	"github.com/raunakjodhawat/gorithims/src/customAlgorithims/multisorted"
	rsort "github.com/raunakjodhawat/gorithims/src/standardAlgorithims/sort"
)

func main(){
	fmt.Println("Starting Custom Algorithm usage")
	multiSortExecution()
	fmt.Println("Standard Algorithm usage")
	algorithmExecution()
}


func multiSortExecution(){
	fmt.Printf("Multisort, Exports two functions \t MultiSorted() and \tHelp()\n")
	type Person struct {
		Name string
		Age int
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
	func (){
		// Sort Algorithm execution examples
		fmt.Println("Warn: rsort only support int slices")
		unsortedSlice := []int{25, 17, 31, 13, 2}
		sortedSlice := rsort.InsertionSort(unsortedSlice, true) // Can use Bubble as well as Selection sort instead
		fmt.Println("Sorted in Descending order",sortedSlice)

		sortedSlice = rsort.InsertionSort(unsortedSlice) // for ascending pass in false flag or dont pass the second argument
		fmt.Println("Sorted in Ascending order",sortedSlice)
	}()
}