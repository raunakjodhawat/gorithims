package queue

import (
	"fmt"
	"testing"
)

func BenchmarkQueue_Add(b *testing.B) {
	q := Queue{}
	for i := 0; i < b.N; i++ {
		isAdded := q.Add(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
}

func BenchmarkQueue_Element(b *testing.B) {
	q := Queue{}
	for i := 0; i < b.N; i++ {
		_ = q.Element()
	}
}

//func BenchmarkQueue_IsEmpty(b *testing.B){
//	q := Queue{}
//	fmt.Println(q.IsEmpty())
//	for i := 0; i < 3; i++ {
//		isAdded := q.Add(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	fmt.Println(q.IsEmpty())
//	// Output:
//	//true
//	//false
//}
//
//func BenchmarkQueue_Offer(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 3; i++ {
//		isAdded := q.Offer(i)
//		fmt.Println(isAdded)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	// Output:
//	//true
//	//true
//	//true
//}
//
//func BenchmarkQueue_Peek(b *testing.B){
//	q := Queue{}
//	lastAddedElement := q.Peek()
//	fmt.Println("last added element is", lastAddedElement)
//	for i := 0; i < 3; i++ {
//		isAdded := q.Add(i)
//		lastAddedElement = q.Peek()
//		fmt.Println("last added element is", lastAddedElement.Val)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	// Output:
//	//last added element is <nil>
//	//last added element is 0
//	//last added element is 1
//	//last added element is 2
//}
//
//func BenchmarkQueue_Poll(b *testing.B){
//	q := Queue{}
//	lastAddedElement, err := q.Poll()
//	if err != nil {
//		fmt.Println("Error in getting the last element")
//	}
//	fmt.Println("last added element is", lastAddedElement)
//	for i := 0; i < 3; i++ {
//		isAdded := q.Add(i)
//		lastAddedElement, err = q.Poll()
//		fmt.Println("last added element is", lastAddedElement.Val)
//		if !isAdded || err != nil {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	// Output:
//	//Error in getting the last element
//	//last added element is <nil>
//	//last added element is 0
//	//last added element is 1
//	//last added element is 2
//}
//
//func BenchmarkQueue_Remove(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	removedElement, err := q.Remove()
//	if err != nil {
//		fmt.Println("Error in removing from the queue", err)
//	}
//	fmt.Println(removedElement.Val) // Output: 0
//}
//
//func BenchmarkQueue_Size(b *testing.B){
//	q := Queue{}
//	fmt.Println(q.Size())
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	fmt.Println(q.Size())
//	// Output: 0
//	//10
//}
//
//func BenchmarkQueue_PrintPretty(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	q.PrintPretty(b *testing.B)// Pretty print a queue
//}
//
//func BenchmarkQueue_Print(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	q.Print(b *testing.B)// print queue, without debugging
//}
//
//func BenchmarkQueue_Print2(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	q.Print(true) // print queue with debugging
//}
//
//func BenchmarkQueue_Print3(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	q.Print(false) // print queue without debugging
//}
//
//func BenchmarkQueue_PrintReverse(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	q.PrintReverse(b *testing.B)// reverse print queue without debugging
//}
//
//func BenchmarkQueue_PrintReverse2(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	q.PrintReverse(false) // reverse print queue without debugging
//}
//
//func BenchmarkQueue_PrintReverse3(b *testing.B){
//	q := Queue{}
//	for i := 0; i < 10; i++ {
//		isAdded := q.Offer(i)
//		if !isAdded {
//			fmt.Println("unable to add element to the queue")
//		}
//	}
//	q.PrintReverse(true) // reverse print queue without debugging
//}
