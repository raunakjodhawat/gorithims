package Utility

import (
	"fmt"
	"time"
)

/*
* External Utility functions
* Returns and prints a 2d string array of list of Description and References
*/
func GetCredits(toPrint ...bool) [][]string{
	defer getExecutionTime(time.Now(), "GetCredits")
	fmt.Print("\nHere's a list of all credits, that I have used for this repository:")
	executionTime := []string{"execution Time", "https://blog.stathat.com/2012/10/10/time_any_function_in_go.html"}
	complexity := []string{"complexity", "https://www.bigocheatsheet.com/"}
	credits := [][]string{executionTime, complexity}
	if toPrint != nil {
		if toPrint[0] == true {
			for i := 0; i < len(credits); i++ {
				fmt.Printf("\n%v\t%v", credits[i][0], credits[i][1])
			}
		}
	}
	return credits
}

// Internal Utility functions
// Credits: https://blog.stathat.com/2012/10/10/time_any_function_in_go.html
func getExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n%s took %s\n\n", name, elapsed)
}
