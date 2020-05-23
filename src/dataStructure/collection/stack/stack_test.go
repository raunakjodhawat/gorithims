package stack

import "testing"

func TestStack_Empty(t *testing.T) {
	s := Stack{}
	if !s.Empty() {
		t.Errorf("Expected a empty stack, but stack was not empty")
	}
	s.Push(100)
	if s.Empty() {
		t.Errorf("Expected a non-empty stack, but stack was empty")
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack{}
	topElement := s.Peek()
	if topElement != nil {
		t.Errorf("Expected top element to be nil, got %v", topElement)
	}
	s.Push(100)
	topElement = s.Peek()
	if topElement.Val != 100 {
		t.Errorf("Expected top element to be 100, got %v", topElement.Val)
	}

	topElement = s.Peek()
	if topElement.Val != 100 {
		t.Errorf("Expected top element to be 100, got %v", topElement.Val)
	}

	s.Push(101)
	s.PrintPretty()
	if topElement.Val != 101 {
		t.Errorf("Expected top element to be 101, got %v", topElement.Val)
	}
}