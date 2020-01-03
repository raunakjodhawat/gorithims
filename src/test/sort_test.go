package test

import (
	"github.com/raunakjodhawat/gorithims/src/sort"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T){
	input := []int{25, 17, 31, 13, 2}
	actualInsetionSortResult := sort.InsertionSort(input)
	actualBubbleSortResult := sort.BubbleSort(input)
	expectedResult := []int{2, 13, 17, 25, 31}

	// Positive cases
	if !reflect.DeepEqual(actualInsetionSortResult, expectedResult) {
		t.Errorf("Insertion Sort failed for input, %v", input)
	}
	if !reflect.DeepEqual(actualBubbleSortResult, expectedResult) {
		t.Errorf("Bubble Sort failed for input, %v", input)
	}

	input = []int{2}
	actualInsetionSortResult = sort.InsertionSort(input)
	actualBubbleSortResult = sort.BubbleSort(input)
	expectedResult = []int{2}
	if !reflect.DeepEqual(actualInsetionSortResult, expectedResult) {
		t.Errorf("Insertion Sort failed for input, %v", input)
	}
	if !reflect.DeepEqual(actualBubbleSortResult, expectedResult) {
		t.Errorf("Bubble Sort failed for input, %v", input)
	}

	input = []int{}
	actualInsetionSortResult = sort.InsertionSort(input)
	actualBubbleSortResult = sort.BubbleSort(input)
	expectedResult = []int{}
	if !reflect.DeepEqual(actualInsetionSortResult, expectedResult) {
		t.Errorf("Insertion Sort failed for input, %v", input)
	}
	if !reflect.DeepEqual(actualBubbleSortResult, expectedResult) {
		t.Errorf("Bubble Sort failed for input, %v", input)
	}

	input = []int{2, 7, 7, 1, 5, 3}
	actualInsetionSortResult = sort.InsertionSort(input)
	actualBubbleSortResult = sort.BubbleSort(input)
	expectedResult = []int{1, 2, 3, 5, 7, 7}
	if !reflect.DeepEqual(actualInsetionSortResult, expectedResult) {
		t.Errorf("Insertion Sort failed for input, %v", input)
	}
	if !reflect.DeepEqual(actualBubbleSortResult, expectedResult) {
		t.Errorf("Bubble Sort failed for input, %v", input)
	}

	input = []int{99, 100, 12}
	actualInsetionSortResult = sort.InsertionSort(input)
	actualBubbleSortResult = sort.BubbleSort(input)
	expectedResult = []int{1, 2, 3, 5, 7}
	if reflect.DeepEqual(actualInsetionSortResult, expectedResult) {
		t.Errorf("Insertion Sort failed for input, %v", input)
	}
	if reflect.DeepEqual(actualBubbleSortResult, expectedResult) {
		t.Errorf("Bubble Sort failed for input, %v", input)
	}

	// Reverse Sort
	input = []int{2, 7, 7, 1, 5, 3}
	actualInsetionSortResult = sort.InsertionSort(input, true)
	actualBubbleSortResult = sort.BubbleSort(input, true)
	expectedResult = []int{7,7,5,3,2,1}
	if !reflect.DeepEqual(actualInsetionSortResult, expectedResult) {
		t.Errorf("Insertion Sort failed for input, %v", input)
	}
	if !reflect.DeepEqual(actualBubbleSortResult, expectedResult) {
		t.Errorf("Bubble Sort failed for input, %v", input)
	}

	input = []int{25, 17, 31, 13, 2}
	actualInsetionSortResult = sort.InsertionSort(input, true)
	actualBubbleSortResult = sort.BubbleSort(input, true)
	expectedResult = []int{31, 25, 17, 13, 2}

	// Positive cases
	if !reflect.DeepEqual(actualInsetionSortResult, expectedResult) {
		t.Errorf("Insertion Sort failed for input, %v", input)
	}
	if !reflect.DeepEqual(actualBubbleSortResult, expectedResult) {
		t.Errorf("Bubble Sort failed for input, %v", input)
	}
}
