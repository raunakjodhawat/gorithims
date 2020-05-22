package queue

import "fmt"

func ExampleQueue_Add() {
	q := Queue{}
	for i := 0; i < 3; i++ {
		isAdded := q.Add(i)
		fmt.Println(isAdded)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	// Output:
	//true
	//true
	//true
}

func ExampleQueue_Element() {
	q := Queue{}
	lastAddedElement := q.Element()
	fmt.Println("last added element is", lastAddedElement)
	for i := 0; i < 3; i++ {
		isAdded := q.Add(i)
		lastAddedElement = q.Element()
		fmt.Println("last added element is", lastAddedElement.Val)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	// Output:
	//last added element is <nil>
	//last added element is 0
	//last added element is 1
	//last added element is 2
}

func ExampleQueue_IsEmpty() {
	q := Queue{}
	fmt.Println(q.IsEmpty())
	for i := 0; i < 3; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	fmt.Println(q.IsEmpty())
	// Output:
	//true
	//false
}

func ExampleQueue_Offer() {
	q := Queue{}
	for i := 0; i < 3; i++ {
		isAdded := q.Offer(i)
		fmt.Println(isAdded)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	// Output:
	//true
	//true
	//true
}

func ExampleQueue_Peek() {
	q := Queue{}
	lastAddedElement := q.Peek()
	fmt.Println("last added element is", lastAddedElement)
	for i := 0; i < 3; i++ {
		isAdded := q.Add(i)
		lastAddedElement = q.Peek()
		fmt.Println("last added element is", lastAddedElement.Val)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	// Output:
	//last added element is <nil>
	//last added element is 0
	//last added element is 1
	//last added element is 2
}

func ExampleQueue_Poll() {
	q := Queue{}
	lastAddedElement, err := q.Poll()
	if err != nil {
		fmt.Println("Error in getting the last element")
	}
	fmt.Println("last added element is", lastAddedElement)
	for i := 0; i < 3; i++ {
		isAdded := q.Add(i)
		lastAddedElement, err = q.Poll()
		fmt.Println("last added element is", lastAddedElement.Val)
		if !isAdded || err != nil {
			fmt.Println("unable to add element to the queue")
		}
	}
	// Output:
	//Error in getting the last element
	//last added element is <nil>
	//last added element is 0
	//last added element is 1
	//last added element is 2
}

func ExampleQueue_Remove() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	removedElement, err := q.Remove()
	if err != nil {
		fmt.Println("Error in removing from the queue", err)
	}
	fmt.Println(removedElement.Val) // Output: 0
}

func ExampleQueue_Size() {
	q := Queue{}
	fmt.Println(q.Size())
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	fmt.Println(q.Size())
	// Output: 0
	//10
}

func ExampleQueue_PrintPretty() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	q.PrintPretty() // Pretty print a queue
}

func ExampleQueue_Print() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	q.Print() // print queue, without debugging
}

//nolint
func ExampleQueue_Print2() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	q.Print(true) // print queue with debugging
}

//nolint
func ExampleQueue_Print3() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	q.Print(false) // print queue without debugging
}

func ExampleQueue_PrintReverse() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	q.PrintReverse() // reverse print queue without debugging
}

//nolint
func ExampleQueue_PrintReverse2() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	q.PrintReverse(false) // reverse print queue without debugging
}

//nolint
func ExampleQueue_PrintReverse3() {
	q := Queue{}
	for i := 0; i < 10; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
	q.PrintReverse(true) // reverse print queue without debugging
}
