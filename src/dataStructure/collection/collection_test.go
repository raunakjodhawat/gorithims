package collection

import (
	"testing"
)

func TestCollection_Add(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	for i := 0; i < 5; i++ {
		if !c.Contains(i) {
			t.Errorf("expected %v, in the collection. But did not find it", i)
		}
	}

	err := c.Add(10, 0) // Add 10 at index 0
	if err != nil || !c.Contains(10) {
		t.Errorf("expected 10, in the collection. But did not find it")
	}

	err = c.Add(100, c.Size()-1) // Add 10 at index 0
	if err != nil || !c.Contains(100) {
		t.Errorf("expected 100, in the collection. But did not find it")
	}

	err = c.Add(12) // Add 10 at index 0
	if err != nil || !c.Contains(12) {
		t.Errorf("expected 12, in the collection. But did not find it")
	}
}

func TestCollection_AddAll(t *testing.T) {
	c := Collection{}
	err := c.AddAll([]int{0, 1, 2, 3, 4}) // Adds elements to the list
	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < 5; i++ {
		if !c.Contains(i) {
			t.Errorf("expected %v, in the collection. But did not find it", i)
		}
	}

	err = c.AddAll([]string{"hello", "world"}, 0)
	if err != nil || !c.Contains("hello") || !c.Contains("world") {
		t.Errorf("expected Hello and world, in the collection. But did not find it")
	}

	err = c.AddAll([]string{"raunak"}, c.Size()-1)
	if err != nil || !c.Contains("raunak") {
		t.Errorf("expected raunak, in the collection. But did not find it")
	}

	err = c.AddAll([]int{1020})
	if err != nil || !c.Contains(1020) {
		t.Errorf("expected 1020, in the collection. But did not find it")
	}
}

func TestCollection_Clear(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	c.Clear()
	if c.Size() != 0 {
		t.Errorf("Expected collection size to be 0, got %v", c.Size())
	}
}

func TestCollection_Contains(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	for i := 0; i < c.Size(); i++ {
		isPresent := c.Contains(i)
		if !isPresent {
			t.Errorf("expected %v to be in the collection, but could not find it", i)
		}
	}

	isPresent := c.Contains(100) // Adds elements to the list
	if isPresent {
		t.Errorf("expected 100 to not be in the list, but it was present")
	}
}

func TestCollection_ContainsAll(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	isPresent := c.ContainsAll(1)
	if !isPresent {
		t.Errorf("expected 1 to be in the collection, but could not find it")
	}
	isPresent = c.ContainsAll(5)
	if isPresent {
		t.Errorf("expected 5 to not be in the collection, but it was there!")
	}

	input := [][]int{
		{0, 1},
		{3},
		{4},
		{5},
		{0, 4},
		{1, 5},
		{0, 1, 2, 3, 4, 5, 6},
	}

	expectedResult := []bool{
		true,
		true,
		true,
		false,
		true,
		false,
		false,
	}

	for i := 0; i < len(input); i++ {
		isPresent = c.ContainsAll(input[i])
		if isPresent != expectedResult[i] {
			t.Errorf("expected %v to be %v, got %v", input[i], expectedResult[i], isPresent)
		}
	}

}

func TestCollection_Equals(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	copyC := Collection{}
	for i := 0; i < 5; i++ {
		err := copyC.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if !copyC.Equals(c) {
		t.Errorf("Collections should be equal")
	}

	err := copyC.Add(100)

	if copyC.Equals(c) && err != nil {
		t.Errorf("Collections should not be equal")
	}
}

func TestCollection_IsEmpty(t *testing.T) {
	c := Collection{}
	if !c.IsEmpty() {
		t.Errorf("Collection should be empty, but it is not")
	}

	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if c.IsEmpty() {
		t.Errorf("Collection should not be empty, but it is")
	}
}

func TestCollection_Iterator(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	iterator := c.Iterator() // Get a iterator for the Collection, starting from head
	var j int
	for iterator != nil { // traverse the list
		if iterator.Val != j {
			t.Errorf("expected %v, got %v", j, iterator.Val)
		}
		j++
		iterator = iterator.Next
	}
}

func TestCollection_Remove(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	removedElement, err := c.Remove()
	if err != nil || removedElement.Val != 0 || removedElement.Prev != nil {
		t.Errorf("Expected removed element to be 0, got %v", removedElement.Val)
	}

	if c.Size() != 4 {
		t.Errorf("Expected collection to have a size of 4, got %v", c.Size())
	}

	removedElement, err = c.Remove(0)
	if err != nil || removedElement.Val != 1 || removedElement.Prev != nil {
		t.Errorf("Expected removed element to be 1, got %v", removedElement.Val)
	}

	if c.Size() != 3 {
		t.Errorf("Expected collection to have a size of 3, got %v", c.Size())
	}

	_, err = c.Remove(c.Size())
	if err == nil {
		t.Errorf("Expected error, because list.size is not a valid index, which can be removed")
	}

	if c.Size() != 3 {
		t.Errorf("Expected collection to have a size of 3, got %v", c.Size())
	}

	for i := 0; i < 3; i++ {
		_, err = c.Remove()
		if err != nil {
			t.Errorf("Expected error, because list.size is not a valid index, which can be removed")
		}
	}
	if c.Size() != 0 {
		t.Errorf("Expected collection to have a size of 0, got %v", c.Size())
	}
}

func TestCollection_RemoveAll(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if c.Size() != 5 {
		t.Errorf("Expected collection to have a size of 5, got %v", c.Size())
	}

	c.RemoveAll()
	c.PrintPretty()
	if c.Size() != 0 {
		t.Errorf("Expected collection to have a size of 0, got %v", c.Size())
	}
}

func TestCollection_Size(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the list
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if c.Size() != 5 {
		t.Errorf("Expected collection to have a size of 5, got %v", c.Size())
	}

	c.RemoveAll()
	c.PrintPretty()
	if c.Size() != 0 {
		t.Errorf("Expected collection to have a size of 0, got %v", c.Size())
	}
}

func TestCollection_ToArray(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the Collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	slice := c.ToArray()
	for i := 0; i < len(slice); i++ {
		if slice[i] != i {
			t.Errorf("expected %v, got %v", i, slice[i])
		}
	}
}

func TestCollection_DeepEquals(t *testing.T) {
	c := Collection{}
	for i := 0; i < 5; i++ {
		err := c.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	copyC := Collection{}
	for i := 0; i < 5; i++ {
		err := copyC.Add(i) // Adds elements to the collection
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if copyC.DeepEquals(c) {
		t.Errorf("Collections are not deep equal")
	}
}
