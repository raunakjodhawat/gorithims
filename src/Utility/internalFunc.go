package Utility

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

