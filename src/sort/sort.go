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
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
				continue
			}
		}
	}
	if utility.EvaluateOptionalBoolFlag(reverse) {
		for i, j := 0, len(input) - 1; i < j; i, j = i+1, j-1 {
			input[i], input[j] = input[j], input[i]
		}
	}
	return input
}