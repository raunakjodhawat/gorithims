package main

import (
	"github.com/raunakjodhawat/gorithims/src/algos"
)

func main() {
	l := algos.ListNode{}
	l.Push(2)
	l.Push(31)
	l.Push(21)
	l.Push(41)
	l.Push(51)
	l.Push(61)
	l.DebugPrintList()
	l.PrintReverseList()

}
