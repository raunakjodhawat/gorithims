package linkedlist

import (
	"reflect"
	"testing"
)

func TestListNode_Add(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	for i := 0; i < 5; i++ {
		val, err := list.Get(i)
		if val != i || err != nil {
			t.Errorf("expected %v, got %v", i, val)
		}
	}

	err := list.Add(10, 0) // Add 10 at index 0
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err := list.Get(0)
	if val != 10 || err != nil {
		t.Errorf("expected 10, got %v", val)
	}

	err = list.Add(11) // Add 11 at end
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err = list.Get(list.Size() - 1)
	if val != 11 || err != nil {
		t.Errorf("expected 10, got %v", val)
	}

	err = list.Add(12, 3) // Add 12 at index 3
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err = list.Get(3)
	if val != 12 || err != nil {
		t.Errorf("expected 12, got %v", val)
	}

	err = list.Add(13, list.Size()+1) // Add 12 at invalid index
	if err == nil {
		// Expect a error
		t.Errorf(err.Error())
	}
}

func TestListNode_AddAll(t *testing.T) {
	list := ListNode{}                       // creates a instance of linked list node
	err := list.AddAll([]int{0, 1, 2, 3, 4}) // Adds elements to the list
	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < list.Size(); i++ {
		val, err := list.Get(i)
		if val != i || err != nil {
			t.Errorf("expected %v, got %v", i, val)
		}
	}

	err = list.AddAll([]string{"hello", "world"}, 0) // Add at index 0
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err := list.Get(0)
	if val != "hello" || err != nil {
		t.Errorf("expected hello, got %v", val)
	}

	err = list.AddAll([]int{11}) // Add 11 at end
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err = list.Get(list.Size() - 1)
	if val != 11 || err != nil {
		t.Errorf("expected 11, got %v", val)
	}

	err = list.AddAll([]bool{true}, 3) // Add 12 at index 3
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err = list.Get(3)
	if reflect.ValueOf(val) == reflect.ValueOf(true) || err != nil {
		t.Errorf("expected true, got %v", val)
	}

	err = list.AddAll([]float32{13.34}, list.Size()+1) // Add 12 at invalid index
	if err == nil {
		// Expect a error
		t.Errorf(err.Error())
	}
}

func TestListNode_AddFirst(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	val, err := list.Get(0)
	if val != 0 || err != nil {
		t.Errorf("expected 0, got %v", val)
	}

	err = list.AddFirst("first element") // Add first element
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err = list.Get(0)
	if val != "first element" || err != nil {
		t.Errorf("expected first element, got %v", val)
	}
}

func TestListNode_AddLast(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	val, err := list.Get(list.Size() - 1)
	if val != 4 || err != nil {
		t.Errorf("expected 4, got %v", val)
	}

	err = list.AddLast("last element") // Add first element
	if err != nil {
		t.Errorf(err.Error())
	}
	val, err = list.Get(list.Size() - 1)
	if val != "last element" || err != nil {
		t.Errorf("expected last element, got %v", val)
	}
}

func TestListNode_Clone(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	newList := list.Clone()
	if list.Size() != newList.Size() {
		t.Errorf("Expected new list size to be %v, got %v", list.Size(), newList.Size())
	}

	for i := 0; i < list.Size(); i++ {
		val, err := list.Get(i)
		newVal, newErr := newList.Get(i)
		if val != newVal || err != nil || newErr != nil {
			t.Errorf("expected %v, got %v", val, newVal)
		}
	}

	list.Clear()
	if newList.Size() == 0 {
		t.Errorf("Expected new list size to be %v, got 0", newList.Size())
	}
}

func TestListNode_Clear(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Expected new list size to be %v, got 0", list.Size())
	}
}
