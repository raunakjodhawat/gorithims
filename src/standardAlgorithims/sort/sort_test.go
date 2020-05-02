package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	inputs := [][]int{
		{25, 17, 31, 13, 2},
		{2},
		{},
		{2, 7, 7, 1, 5, 3},
		{99, 100, 12},
		{2, 7, 7, 1, 5, 3},
		{25, 17, 31, 13, 2, 6},
	}
	expectedResults := [][]int{
		{2, 13, 17, 25, 31},
		{2},
		{},
		{1, 2, 3, 5, 7, 7},
		{1, 2, 3, 5, 7},
		{7, 7, 5, 3, 2, 1},
		{31, 25, 17, 13, 6, 2},
	}
	reverseProperties := []bool{
		false,
		false,
		false,
		false,
		false,
		true,
		true,
	}

	positiveNegativeTestCases := []bool{
		false,
		false,
		false,
		false,
		true,
		false,
		false,
	}

	functionsToTest := []interface{}{
		algirthims.InsertionSort,
		algirthims.BubbleSort,
		algirthims.SelectionSort,
	}

	for i := 0; i < len(inputs); i++ {
		for _, f := range functionsToTest {
			actualSortedResult := f.(func([]int, ...bool) []int)(inputs[i], reverseProperties[i])
			expectedResult := expectedResults[i]
			if !reflect.DeepEqual(actualSortedResult, expectedResult) && !positiveNegativeTestCases[i] {
				t.Errorf("%T Sort failed for input, %v", f, inputs[i])
			}
		}
	}
}
