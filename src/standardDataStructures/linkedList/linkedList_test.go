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

func TestListNode_Contains(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	for i := 0; i < list.Size(); i++ {
		isPresent := list.Contains(i) // Adds elements to the list
		if !isPresent {
			t.Errorf("expected %v to be in the list, but could not find it", i)
		}
	}

	isPresent := list.Contains(100) // Adds elements to the list
	if isPresent {
		t.Errorf("expected 100 to not be in the list, but it was present")
	}
}

func TestListNode_Element(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	headElement := list.Element() // Adds elements to the list
	if headElement.Val != 0 || headElement.Prev != nil {
		t.Errorf("expected 0 as thead element, and previous element to be nil. got %v, %v", headElement.Val, headElement.Prev)
	}
}

func TestListNode_Get(t *testing.T) {
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
}

func TestListNode_GetFirst(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	val, err := list.GetFirst()
	if val != 0 || err != nil {
		t.Errorf("expected 0, got %v", val)
	}
}

func TestListNode_GetLast(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	val, err := list.GetLast()
	if val != 4 || err != nil {
		t.Errorf("expected 4, got %v", val)
	}
}

func TestListNode_IndexOf(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	index := list.IndexOf(0)
	if index == -1 || index != 0 {
		t.Errorf("expected element to be at 0 index, but the element was not present in the list. got: %v", index)
	}

	index = list.IndexOf(100)
	if index != -1 {
		t.Errorf("expected element to be not be in the list and thus return -1")
	}
}

func TestListNode_LastIndexOf(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(0) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	index := list.LastIndexOf(0)
	if index == -1 || index != list.Size()-1 {
		t.Errorf("expected element to be at last index, but the element was not present in the list. got: %v", index)
	}

	index = list.LastIndexOf(100)
	if index != -1 {
		t.Errorf("expected element to be not be in the list and thus return -1")
	}
}

func TestListNode_ListIterator(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	iterator, err := list.ListIterator(0) // Get a iterator for the list, starting from head
	var j int
	for iterator != nil { // traverse the list
		if iterator.Val != j || err != nil {
			t.Errorf("expected %v, got %v", j, iterator.Val)
		}
		j++
		iterator = iterator.Next
	}

	j--
	iterator, err = list.ListIterator(list.Size() - 1) // Get a iterator for the list, starting from head
	for iterator != nil {
		if iterator.Val != j || err != nil {
			t.Errorf("expected %v, got %v", j, iterator.Val)
		}
		iterator = iterator.Next
	}
}

func TestListNode_Offer(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	isAdded, err := list.Offer(100) // offer/add element to end of the list
	if err != nil || !isAdded {
		t.Errorf("Failed to add 100 to end of the list")
	}
	lastElement, err := list.Get(list.Size() - 1)
	if lastElement != 100 || err != nil {
		t.Errorf("Expected last element to be 100, got %v", lastElement)
	}
}

func TestListNode_OfferFirst(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	isAdded, err := list.OfferFirst(100) // offer/add element to end of the list
	if err != nil || !isAdded {
		t.Errorf("Failed to add 100 to end of the list")
	}
	lastElement, err := list.Get(0)
	if lastElement != 100 || err != nil {
		t.Errorf("Expected last element to be 100, got %v", lastElement)
	}
}

func TestListNode_OfferLast(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	isAdded, err := list.OfferLast(100) // offer/add element to end of the list
	if err != nil || !isAdded {
		t.Errorf("Failed to add 100 to end of the list")
	}
	lastElement, err := list.Get(list.Size() - 1)
	if lastElement != 100 || err != nil {
		t.Errorf("Expected last element to be 100, got %v", lastElement)
	}
}

func TestListNode_Peek(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	headElement, err := list.Peek()
	if err != nil || headElement.Val != 0 || headElement.Prev != nil {
		t.Errorf("Head element value is not 0")
	}
}

func TestListNode_PeekFirst(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	headElement, err := list.PeekFirst()
	if err != nil || headElement.Val != 0 || headElement.Prev != nil {
		t.Errorf("Head element value is not 0")
	}
}

func TestListNode_PeekLast(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	tailElement, err := list.PeekLast()
	if err != nil || tailElement.Val != 4 || tailElement.Next != nil {
		t.Errorf("Head element value is not 0")
	}
}

func TestListNode_Poll(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	headElement, err := list.Poll()
	if err != nil || headElement.Val != 0 || headElement.Prev != nil {
		t.Errorf("Head element value is not 0, got %v", headElement.Val)
	}

	headElement, err = list.Poll()
	if err != nil || headElement.Val != 1 || headElement.Prev != nil {
		t.Errorf("Head element value is not 1, got %v", headElement.Val)
	}
}

func TestListNode_PollFirst(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	headElement, err := list.PollFirst()
	if err != nil || headElement.Val != 0 || headElement.Prev != nil {
		t.Errorf("Head element value is not 0, got %v", headElement.Val)
	}

	headElement, err = list.PollFirst()
	if err != nil || headElement.Val != 1 || headElement.Prev != nil {
		t.Errorf("Next head element is not 1, got %v", headElement.Val)
	}
}

func TestListNode_PollLast(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	tailElement, err := list.PollLast()
	if err != nil || tailElement.Val != 4 || tailElement.Next != nil {
		t.Errorf("Next tail element is not 4, got %v", tailElement.Val)
	}

	tailElement, err = list.PollLast()
	if err != nil || tailElement.Val != 3 || tailElement.Next != nil {
		t.Errorf("Next tail element is not 3, got %v", tailElement.Val)
	}
}

func TestListNode_Pop(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	headVal, err := list.Pop()
	if err != nil || headVal != 0 {
		t.Errorf("Head element is expected to be 0, got %v", headVal)
	}

	headVal, err = list.Pop()
	if err != nil || headVal != 1 {
		t.Errorf("Head element is expected to be 0, got %v", headVal)
	}
}

func TestListNode_Push(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	err := list.Push(10)
	if err != nil {
		t.Errorf(err.Error())
	}

	headElement, err := list.Get(0)
	if err != nil || headElement != 10 {
		t.Errorf("Expected head element to be 10, got %v", headElement)
	}
}

func TestListNode_Remove(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	removedElement, err := list.Remove()
	if err != nil || removedElement.Val != 0 || removedElement.Prev != nil {
		t.Errorf("Expected removed element to be 0, got %v", removedElement.Val)
	}

	newHead, err := list.Get(0)
	if err != nil || newHead != 1 {
		t.Errorf("Expected head element to be 1, got %v", newHead)
	}

	removedElement, err = list.Remove(0)
	if err != nil || removedElement.Val != 1 || removedElement.Prev != nil {
		t.Errorf("Expected removed element to be 0, got %v", removedElement.Val)
	}

	newHead, err = list.Get(0)
	if err != nil || newHead != 2 {
		t.Errorf("Expected head element to be 1, got %v", newHead)
	}

	removedElement, err = list.Remove(list.Size())
	if err == nil {
		t.Errorf("Expected error, because list.size is not a valid index, which can be removed")
	}
}

func TestListNode_RemoveFirst(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	removedElement, err := list.RemoveFirst()
	if err != nil || removedElement.Val != 0 || removedElement.Prev != nil {
		t.Errorf("Expected removed element to be 0, got %v", removedElement.Val)
	}

	newHead, err := list.Get(0)
	if err != nil || newHead != 1 {
		t.Errorf("Expected head element to be 1, got %v", newHead)
	}
}

func TestListNode_RemoveLast(t *testing.T) {
	list := ListNode{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	removedElement, err := list.RemoveLast()
	if err != nil || removedElement.Val != 4 || removedElement.Next != nil {
		t.Errorf("Expected removed element to be 4, got %v", removedElement.Val)
	}

	newHead, err := list.Get(0)
	if err != nil || newHead != 0 {
		t.Errorf("Expected head element to be 0, got %v", newHead)
	}

	newHead, err = list.Get(list.Size() - 1)
	if err != nil || newHead != 3 {
		t.Errorf("Expected tail element to be 3, got %v", newHead)
	}
}
