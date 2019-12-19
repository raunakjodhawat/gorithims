package main

import (
	"fmt"
	complexity "github.com/gorithims/static"
	"time"
)

func main() {
	a := "getAllComplexity"
	switch a {
	case "getAllComplexity":
		defer getExecutionTime(time.Now(), a)
		print(complexity.GetAllComplexity())
		break
	default:
		fmt.Println("hi")
		break
	}
}

// External Utility functions
func GetCredits() [][]string {
	executionTime := []string{"execution Time", "https://blog.stathat.com/2012/10/10/time_any_function_in_go.html"}
	complexity := []string{"complexity", "https://www.bigocheatsheet.com/"}
	return [][]string{executionTime, complexity}
}

// Internal Utility functions
// Credits: https://blog.stathat.com/2012/10/10/time_any_function_in_go.html
func getExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func print(input [][][]string) {
	for i:=0; i<len(input[0]); i++ {
		fmt.Println(input[i])
	}
}
