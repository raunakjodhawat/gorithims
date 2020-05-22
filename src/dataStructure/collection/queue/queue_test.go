package queue

import "testing"

func TestQueue_Add(t *testing.T) {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}
	if q.Size() != 10 {
		t.Errorf("Expected queue size to be 10, got %v", q.Size())
	}
}

func TestQueue_Element(t *testing.T) {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}

	if q.Element().Val != 9 {
		t.Errorf("Expected tail element to be 9, got %v", q.Element().Val)
	}

	if q.Element().Val != 9 {
		t.Errorf("Expected tail element to be 9, got %v", q.Element().Val)
	}
}

func TestQueue_Offer(t *testing.T) {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}
	if q.Size() != 10 {
		t.Errorf("Expected queue size to be 10, got %v", q.Size())
	}
}

func TestQueue_Peek(t *testing.T) {
	q := Queue{}
	if q.Peek() != nil {
		t.Errorf("Expected head element to be nil, but got %v", q.Peek())
	}
	for i := 0; i < 10; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}

	if q.Peek().Val != 9 {
		t.Errorf("Expected tail element to be 9, got %v", q.Peek().Val)
	}

	if q.Peek().Val != 9 {
		t.Errorf("Expected tail element to be 9, got %v", q.Peek().Val)
	}
}

func TestQueue_Poll(t *testing.T) {
	q := Queue{}
	polledElement, err := q.Poll()

	if polledElement != nil || err == nil {
		t.Errorf("No element was expected to be removed, and an error was expected, but got no errors and polledElement is %v ", polledElement)
	}
	for i := 0; i < 10; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}
	polledElement, err = q.Poll()

	if polledElement.Val != 9 || err != nil {
		t.Errorf("Expected polled element to be 9, got %v", polledElement.Val)
	}

	polledElement, err = q.Poll()
	if polledElement.Val != 8 || err != nil {
		t.Errorf("Expected polled element to be 8, got %v", polledElement.Val)
	}
}

func TestQueue_Remove(t *testing.T) {
	q := Queue{}
	removedElement, err := q.Remove()

	if removedElement != nil || err == nil {
		t.Errorf("No element was expected to be removed, and an error was expected, but got no errors and removed Element is %v ", removedElement)
	}
	for i := 0; i < 10; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}
	removedElement, err = q.Remove()

	if removedElement.Val != 0 || err != nil {
		t.Errorf("Expected polled element to be 0, got %v", removedElement.Val)
	}

	removedElement, err = q.Remove()
	if removedElement.Val != 1 || err != nil {
		t.Errorf("Expected polled element to be 1, got %v", q.Peek().Val)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := Queue{}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but current queue size is %v", q.Size())
	}
	for i := 0; i < 10; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}
	if q.IsEmpty() {
		t.Errorf("Expected queue to be not empty, but current queue size empty")
	}
}

func TestQueue_Size(t *testing.T) {
	q := Queue{}
	if q.Size() != 0 {
		t.Errorf("Expected queue to be empty, but current queue size is %v", q.Size())
	}
	for i := 0; i < 10; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			t.Errorf("unable to add element to the queue")
		}
	}
	if q.Size() != 10 {
		t.Errorf("Expected queue size to be 10, got %v", q.Size())
	}
}
