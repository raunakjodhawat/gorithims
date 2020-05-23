package stack

import "fmt"

func ExampleStack_Empty() {
	s := Stack{}
	fmt.Println(s.Empty())
	_ = s.Push(100)
	fmt.Println(s.Empty())
	//Output:
	//true
	//false
}

func ExampleStack_Peek() {
	s := Stack{}
	fmt.Println(s.Peek())
	_ = s.Push(100)
	fmt.Println(s.Peek().Val)
	//Output:
	//<nil>
	//100
}

func ExampleStack_Pop() {
	s := Stack{}
	_ = s.Push(100)
	poppedElement, _ := s.Pop()
	fmt.Println(poppedElement.Val) // Output: 100
}

func ExampleStack_Push() {
	s := Stack{}
	fmt.Println(s.Push(100)) //Output: <nil>
}

func ExampleStack_Search() {
	s := Stack{}
	fmt.Println(s.Search(100))
	_ = s.Push(100)
	fmt.Println(s.Search(100))
	//Output:
	//-1
	//1
}
