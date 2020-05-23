package stack

import (
	"fmt"
	"testing"
)

func BenchmarkStack_Push(b *testing.B) {
	s := Stack{}
	for i := 0; i < b.N; i++ {
		err := s.Push(i)
		if err != nil {
			fmt.Printf("There's a Error in pushing %v to the stack.\n", i)
		}
	}
}

func BenchmarkStack_Peek(b *testing.B) {
	s := Stack{}
	for i := 0; i < b.N; i++ {
		err := s.Push(i)
		if err != nil {
			fmt.Printf("There's a Error in pushing %v to the stack.\n", i)
		}
		s.Peek()
	}
}

func BenchmarkStack_Empty(b *testing.B) {
	s := Stack{}
	for i := 0; i < b.N; i++ {
		err := s.Push(i)
		if err != nil {
			fmt.Printf("There's a Error in pushing %v to the stack.\n", i)
		}
		s.Empty()
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	s := Stack{}
	for i := 0; i < b.N; i++ {
		err := s.Push(i)
		if err != nil {
			fmt.Printf("There's a Error in pushing %v to the stack.\n", i)
		}
		_, _ = s.Pop()
	}
}

func BenchmarkStack_Search(b *testing.B) {
	s := Stack{}
	for i := 0; i < b.N; i++ {
		err := s.Push(i)
		if err != nil {
			fmt.Printf("There's a Error in pushing %v to the stack.\n", i)
		}
		s.Search(i)
	}
}
