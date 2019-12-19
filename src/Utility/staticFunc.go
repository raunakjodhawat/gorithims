package Utility

import (
	"fmt"
	"time"
)

// credits: https://www.bigocheatsheet.com/
var timeComplexity = [][]string{
	{"Sorting Algorithim", "Best", "Average", "Worst"},
	{"Insertion Sort", "Ω(n)", "Θ(n^2)", "O(n^2)"},
	{"Selection Sort", "Ω(n^2)", "Θ(n^2)", "O(n^2)"},
	{"Merge Sort", "Ω(n log(n))", "Θ(n log(n))", "O(n log(n))"},
	{"Heap Sort", "Ω(n log(n))", "Θ(n log(n))", "O(n log(n))"},
	{"Quick Sort", "Ω(n log(n))", "Θ(n log(n))", "O(n^2)"},
	{"Bubble Sort", "Ω(n)", "Θ(n^2)", "O(n^2)"},
	{"Bucket Sort", "Ω(n+k)", "Θ(n+k)", "O(n^2)"},
	{"Radix Sort", "Ω(nk)", "Θ(nk)", "O(nk)"},
	{"k is the size of bucket"},
}

// credits: https://www.bigocheatsheet.com/
var spaceComplexity = [][]string{
	{"Sorting Algorithim", "Worst"},
	{"Insertion Sort", "O(1)"},
	{"Selection Sort", "O(1)"},
	{"Merge Sort", "O(n)"},
	{"Heap Sort", "O(1)"},
	{"Quick Sort", "O(log(n))"},
	{"Bubble Sort", "O(1)"},
	{"Bucket Sort", "O(n)"},
	{"Radix Sort", "O(n+k)"},
	{"k is the maximum keylength"},
}

func GetAllComplexity(toPrint ...bool) [][][]string {
	defer GetExecutionTime(time.Now(), "GetAllComplexity")
	allComplexity := [][][]string{timeComplexity, spaceComplexity}
	if ShouldPrint(toPrint) {
		fmt.Print("\nHere's a list of all complexities, covered in this repo (credits: https://www.bigocheatsheet.com/)")
		Print(timeComplexity, "Time Complexity List")
		Print(spaceComplexity, "Space Complexity List")
	}
	return allComplexity
}

func GetAllTimeComplexity(toPrint ...bool) [][]string {
	defer GetExecutionTime(time.Now(), "GetAllTimeComplexity")
	if ShouldPrint(toPrint) {
		fmt.Print("\nHere's a list of all Time complexities, covered in this repo (credits: https://www.bigocheatsheet.com/)")
		Print(timeComplexity, "Time Complexity List")
	}
	return timeComplexity
}

func GetAllSpaceComplexity(toPrint ...bool) [][]string {
	defer GetExecutionTime(time.Now(), "GetAllSpaceComplexity")
	if ShouldPrint(toPrint) {
		fmt.Print("\nHere's a list of all Space complexities, covered in this repo (credits: https://www.bigocheatsheet.com/)")
		Print(spaceComplexity, "Time Complexity List")
	}
	return spaceComplexity
}

func GetTimeComplexityByKey(key []string) []string {
	timeComplexityBykey := []string{}
	return timeComplexityBykey
}

func GetSpaceComplexityByKey(key []string) []string {
	spaceComplexityByKey := []string{}
	return spaceComplexityByKey
}

func GetAllComplexityByKey(key []string) [][]string {
	timeComplexityBykey := GetTimeComplexityByKey(key)
	spaceComplexityByKey := GetSpaceComplexityByKey(key)
	return [][]string{timeComplexityBykey, spaceComplexityByKey}
}