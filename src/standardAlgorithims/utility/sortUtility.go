package utility

func Swap(input []int, iIndex int, jIndex int) []int {
	input[iIndex], input[jIndex] = input[jIndex], input[iIndex]
	return input
}

func Reverse(input []int, reverse ...bool) []int {
	if EvaluateOptionalBoolFlag(reverse) {
		for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
			input = Swap(input, i, j)
		}
	}
	return input
}
