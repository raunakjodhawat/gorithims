package stack

import (
	"testing"
)

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
	topElement = s.Peek()

	if topElement.Val != 101 {
		t.Errorf("Expected top element to be 101, got %v", topElement.Val)
	}
}

func TestStack_Pop(t *testing.T) {
	s := Stack{}
	poppedElement, error := s.Pop()
	if error == nil {
		t.Errorf("Expected error, because stack is empty")
	}
	s.Push(100)
	poppedElement, error = s.Pop()
	if error != nil || poppedElement.Val != 100 {
		t.Errorf("Expected element to be removed at 100, removedElement: %v", poppedElement)
	}
	poppedElement, error = s.Pop()
	if error == nil {
		t.Errorf("Expected error, because stack is empty")
	}
	poppedElement, error = s.Pop()
	if error == nil {
		t.Errorf("Expected error, because stack is empty")
	}
	s.Push(100)
	s.Push(100)
	poppedElement, error = s.Pop()
	if error != nil || poppedElement.Val != 100 {
		t.Errorf("Expected element to be removed at 100, removedElement: %v", poppedElement)
	}
	poppedElement, error = s.Pop()
	if error != nil || poppedElement.Val != 100 {
		t.Errorf("Expected element to be removed at 100, removedElement: %v", poppedElement)
	}
}

func TestStack_Push(t *testing.T) {
	s := Stack{}
	for i := 0; i < 10; i++ {
		err := s.Push(i)
		if err != nil {
			t.Errorf("Got error while pushing to the stack")
		}
		peekElement := s.Peek()
		if peekElement.Val != i {
			t.Errorf("Expected peek element to be %v, got %v", i, peekElement.Val)
		}
	}
}

func TestStack_Search(t *testing.T) {
	s := Stack{}
	searchElementIndex := s.Search(0)
	if searchElementIndex != -1 {
		t.Errorf("Expected element to not be in the stack, but found it at %v: index (1-based)", searchElementIndex)
	}
	s.Push(0)
	searchElementIndex = s.Search(0)
	if searchElementIndex != 1 {
		t.Errorf("Expected element to be at 1st index (1-based), but found it at %v: index (1-based)", searchElementIndex)
	}

	for i := 0; i < 10; i++ {
		err := s.Push(i)
		if err != nil {
			t.Errorf("Got error while pushing to the stack")
		}
	}
	for i := 0; i < 10; i++ {
		searchElementIndex = s.Search(i)
		if searchElementIndex != (10 - i) {
			t.Errorf("Expected element to be at %v index (1-based), but found it at %v: index (1-based)", 10-i, searchElementIndex)
		}
	}
}
