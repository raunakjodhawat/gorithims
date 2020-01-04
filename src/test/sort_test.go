package test

import (
	"github.com/raunakjodhawat/gorithims/src/sort"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	inputs := [][]int{
		[]int{25, 17, 31, 13, 2},
		[]int{2},
		[]int{},
		[]int{2, 7, 7, 1, 5, 3},
		[]int{99, 100, 12},
		[]int{2, 7, 7, 1, 5, 3},
		[]int{25, 17, 31, 13, 2, 6},
	}
	expectedResults := [][]int{
		[]int{2, 13, 17, 25, 31},
		[]int{2},
		[]int{},
		[]int{1, 2, 3, 5, 7, 7},
		[]int{1, 2, 3, 5, 7},
		[]int{7, 7, 5, 3, 2, 1},
		[]int{31, 25, 17, 13, 6, 2},
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
		sort.InsertionSort,
		sort.BubbleSort,
		sort.SelectionSort,
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
