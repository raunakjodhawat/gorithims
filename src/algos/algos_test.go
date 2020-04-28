package algos

import "testing"

func TestListNode_First(t *testing.T) {
	list := ListNode{}
	stringToPush := "raunak jodhawat"
	list.Push(stringToPush)
	firstElement, err := list.First()
	if err != nil || firstElement != stringToPush {
		t.Errorf("expected %v got %v", stringToPush, firstElement)
	}
}

func TestListNode_Last(t *testing.T) {
	list := ListNode{}
	for i:=0; i<20; i++ {
		list.Push(i)
	}
	lastElementToPush := "raunak jodhawat"
	list.Push(lastElementToPush)
	lastElement, err := list.Last()
	if err != nil || lastElement != lastElementToPush {
		t.Errorf("expected %v got %v", lastElementToPush, lastElement)
	}
}

func TestListNode_Length(t *testing.T) {
	list := ListNode{}
	for i:=0; i<20; i++ {
		list.Push(i)
	}
	listLen := list.Length()
	if listLen != 20 {
		t.Errorf("expected %v got %v", 20, listLen)
	}
}

func TestListNode_GetFirstMatchIndex(t *testing.T) {
	list := ListNode{}
	for i:=0; i<20; i++ {
		list.Push(i)
	}
	for i:=0; i<20; i++ {
		firstMatchIndex := list.GetFirstMatchIndex(i)
		if firstMatchIndex != i {
			t.Errorf("expected %v got %v", i, firstMatchIndex)
		}
	}
}

func TestListNode_RemoveFirst(t *testing.T) {
	list := ListNode{}
	for i:=0; i<5; i++ {
		list.Push(i)
	}
	initialActualResult, err := list.First()
	list.RemoveFirst()
	actualResult, err := list.First()
	if initialActualResult != 0 && actualResult != 1 && err == nil {
		t.Errorf("Expected first value to be deleted, got first value as %v", actualResult)
	}
}