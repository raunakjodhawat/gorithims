//nolint
package collection

import "testing"

func BenchmarkCollection_Add(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.Add("A")
	}
}

func BenchmarkCollection_AddAll(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.AddAll([]int{0, 1})
	}
}

func BenchmarkCollection_Clear(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.Clear()
	}
}

func BenchmarkCollection_Contains(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.Contains(0)
	}
}

func BenchmarkCollection_ContainsAll(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.ContainsAll([]int{0, 1})
	}
}

func BenchmarkCollection_DeepEquals(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.DeepEquals(Collection{})
	}
}

func BenchmarkCollection_Equals(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.Equals(Collection{})
	}
}

func BenchmarkCollection_IsEmpty(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.IsEmpty()
	}
}

func BenchmarkCollection_Iterator(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.Iterator()
	}
}

func BenchmarkCollection_Remove(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.Remove()
	}
}

func BenchmarkCollection_RemoveAll(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.RemoveAll()
	}
}

func BenchmarkCollection_Size(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.Size()
	}
}

func BenchmarkCollection_ToArray(b *testing.B) {
	c := Collection{}
	for i := 0; i < b.N; i++ {
		c.ToArray()
	}
}
