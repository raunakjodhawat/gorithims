package utility

import (
	"fmt"
	"time"
)

// These are common functions that can be used by any packages/files

// Internal Utility functions
// Credits: https://blog.stathat.com/2012/10/10/time_any_function_in_go.html
func GetExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n%s took %s\n\n", name, elapsed)
}

func ShouldPrint(toPrint ...[]bool) bool {
	if toPrint[0] == nil {
		return false
	} else {
		if toPrint[0][0] == false {
			return false
		}
		return true
	}
}

func Print(input [][]string, heading string) {
	fmt.Printf("\n%v\n",heading)
	for i:=0; i< len(input); i++ {
		for j:=0; j< len(input[i]); j++ {
			fmt.Printf("%v\t\t", input[i][j])
		}
		fmt.Println()
	}
}