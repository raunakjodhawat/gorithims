package rsort

import (
	"github.com/raunakjodhawat/gorithims/src/standardAlgorithims/utility"
	"time"
)

// This is for insertion sort
func InsertionSort(input []int, reverse ...bool) []int {
	defer utility.GetExecutionTime(time.Now(), "InsertionSort")
	// return if length is 0
	if len(input) <= 1 {
		return input
	}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input = utility.Swap(input, i, j)
			}
		}
	}
	return utility.Reverse(input, reverse...)
}

func BubbleSort(input []int, reverse ...bool) []int {
	defer utility.GetExecutionTime(time.Now(), "BubbleSort")
	return utility.Reverse(recursiveBubbleSort(input, reverse...), reverse...)
}

func recursiveBubbleSort(input []int, reverse ...bool) []int {
	if len(input) <= 1 {
		return input
	}
	swapCount := 0
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) && input[i] > input[i+1] {
			input = utility.Swap(input, i, i+1)
			swapCount++
		}
	}
	if swapCount == 0 {
		return input
	}
	return recursiveBubbleSort(input, reverse...)
}

func SelectionSort(input []int, reverse ...bool) []int {
	defer utility.GetExecutionTime(time.Now(), "SelectionSort")
	if len(input) <= 1 {
		return input
	}

	inputCopy := make([]int, len(input))
	var minElement int
	currentPointer := 0 // starting point for inner loop
	changeIndex := 0    // swapping index

	for i := 0; i < len(input)-1; i++ {
		// loop to find min element
		minElement = input[currentPointer]
		for j := currentPointer; j < len(input); j++ {
			if input[j] < minElement {
				minElement = input[j] // change min element
				changeIndex = j       // store the index of minimum element
			}
		}
		inputCopy[currentPointer], input[changeIndex] = minElement, input[currentPointer]
		currentPointer++
	}
	// copy last element
	inputCopy[len(inputCopy)-1] = input[len(inputCopy)-1]
	return utility.Reverse(inputCopy, reverse...)
}
