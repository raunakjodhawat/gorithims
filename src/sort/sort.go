package sort

import (
	"github.com/raunakjodhawat/gorithims/src/utility"
	"time"
)

func InsertionSort(input []int, reverse ...bool) []int{
	defer utility.GetExecutionTime(time.Now(), "InsertionSort")
	if len(input) <= 1 {
		return input
	}
	for i:=0; i<len(input); i++ {
		for j:= i+1; j<len(input); j++{
			if input[i] > input[j]{
				input = Swap(input, i, j)
			}
		}
	}
	return Reverse(input, reverse...)
}

func BubbleSort(input []int, reverse ...bool) []int{
	defer utility.GetExecutionTime(time.Now(), "BubbleSort")
	return Reverse(RecursiveBubbleSort(input, reverse...), reverse...)
}

func RecursiveBubbleSort(input []int, reverse ...bool) []int {
	if len(input) <= 1 {
		return input
	}
	swapCount := 0
	for i :=0; i<len(input); i++ {
		if i+1 < len(input) && input[i] > input[i+1]{
			input = Swap(input, i, i+1)
			swapCount++
		}
	}
	if swapCount == 0 {
		return input
	} else {
		return RecursiveBubbleSort(input, reverse...)
	}
}