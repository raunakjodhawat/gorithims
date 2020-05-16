package linkedlist

import "fmt"

func ExampleList_Add() {
	list := List{} // creates a instance of linked list node
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err := list.Add(200, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.Add(200, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.Add(200, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleList_AddAll() {
	list := List{} // creates a instance of linked list node
	err := list.AddAll([]int{100, 200, 300})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.AddAll([]int{101, 201, 301}, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.AddAll([]string{"raunak"}, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleList_AddFirst() {
	list := List{}
	err := list.AddLast(100)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleList_AddLast() {
	list := List{}
	err := list.AddLast(100)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleList_Clear() {
	list := List{}
	list.Clear()
}

func ExampleList_Clone() {
	list := List{}
	newList := list.Clone()
	newList.Print()
}

func ExampleList_Contains() {
	list := List{}
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	isPresent := list.Contains(1)
	fmt.Println(isPresent) // Output: true
}

func ExampleList_Element() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	listHead := list.Element()
	fmt.Println(listHead.Val) // Output: 10
}

func ExampleList_Get() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	zeroElement, err := list.Get(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(zeroElement) // Output: 10
}

func ExampleList_GetFirst() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	zeroElement, err := list.GetFirst()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(zeroElement) // Output: 10
}

func ExampleList_GetLast() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	zeroElement, err := list.GetLast()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(zeroElement) // Output: 10
}

func ExampleList_IndexOf() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	index := list.IndexOf(10)
	fmt.Println(index) // Output: 0
}

func ExampleList_LastIndexOf() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	index := list.LastIndexOf(10)
	fmt.Println(index) // Output: 1
}

func ExampleList_ListIterator() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	iterator, err := list.ListIterator(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for iterator != nil {
		fmt.Printf("%v\t", iterator.Val) // Output: 10
		iterator = iterator.Next
	}
}

func ExampleList_Offer() {
	list := List{}
	elementAdded, err := list.Offer(100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(elementAdded) // Output: true
}

func ExampleList_OfferFirst() {
	list := List{}
	elementAdded, err := list.OfferFirst(100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(elementAdded) // Output: true
}

func ExampleList_OfferLast() {
	list := List{}
	elementAdded, err := list.OfferLast(100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(elementAdded) // Output: true
}

func ExampleList_Peek() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	topElement, err := list.Peek()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(topElement.Val) // Output: 10
}

func ExampleList_PeekFirst() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	topElement, err := list.PeekFirst()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(topElement.Val) // Output: 10
}

func ExampleList_PeekLast() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	topElement, err := list.PeekLast()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(topElement.Val) // Output: 10
}

func ExampleList_Poll() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	polledElement, err := list.Poll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(polledElement.Val) // Output: 10
}

func ExampleList_PollFirst() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	polledElement, err := list.PollFirst()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(polledElement.Val) // Output: 10
}

func ExampleList_PollLast() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	polledElement, err := list.PollLast()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(polledElement.Val) // Output: 10
}

func ExampleList_Pop() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	removedElementVal, err := list.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(removedElementVal) // Output: 10
}

func ExampleList_Print() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	list.Print()      // Print with debug option by default set to be false. equivalent to list.Print(false)
	list.Print(true)  // Print with debug option set as true
	list.Print(false) // Print with debug option set as false. equivalent to list.Print()
}

func ExampleList_PrintReverse() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	list.PrintReverse()      // Print in reverse with debug option by default set to be false. equivalent to list.Print(false)
	list.PrintReverse(true)  // Print in reverse with debug option set as true
	list.PrintReverse(false) // Print in reverse with debug option set as false. equivalent to list.Print()
}

func ExampleList_Push() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.Push(100)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleList_Remove() {
	list := List{}
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	removedElement, err := list.Remove()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(removedElement.Val, ", ")

	removedElement, err = list.Remove(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(removedElement.Val) // Output: 0, 3
}

func ExampleList_RemoveFirst() {
	list := List{}
	for i := 0; i < 5; i++ {
		err := list.Add(i) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	removedElement, err := list.RemoveFirst()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(removedElement.Val) // Output: 0
}

func ExampleList_RemoveLast() {
	list := List{}
	for i := 0; i < 5; i++ {
		err := list.Add(i * 10) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	removedElement, err := list.RemoveLast()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(removedElement.Val) // Output: 40
}

func ExampleList_RemoveFirstOccurrence() {
	list := List{}
	for i := 0; i < 5; i++ {
		err := list.Add(i%2 == 0) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	removedElement, err := list.RemoveFirstOccurrence(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(removedElement.Val) // Output: true
}

func ExampleList_RemoveLastOccurrence() {
	list := List{}
	for i := 0; i < 5; i++ {
		err := list.Add(i%2 == 0) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	removedElement, err := list.RemoveLastOccurrence(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(removedElement.Val) // Output: true
}

func ExampleList_Set() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.Set(0, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleList_Size() {
	list := List{}
	err := list.Add(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list.Size()) // Output: 1
}

func ExampleList_ToArray() {
	list := List{}
	for i := 0; i < 5; i++ {
		err := list.Add(i%2 == 0) // Adds elements to the list
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	listSlice := list.ToArray()
	fmt.Println(listSlice) // Output: [true false true false true]
}
