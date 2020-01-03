package sort

import "github.com/raunakjodhawat/gorithims/src/utility"

func Swap(input []int, iIndex int, jIndex int) []int {
	temp := input[iIndex]
	input[iIndex] = input[jIndex]
	input[jIndex] = temp
	return input
}

func Reverse(input []int, reverse ...bool) []int {
	if utility.EvaluateOptionalBoolFlag(reverse) {
		for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
			input[i], input[j] = input[j], input[i]
		}
	}
	return input
}
