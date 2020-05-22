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

func BenchmarkQueue_IsEmpty(b *testing.B) {
	q := Queue{}
	for i := 0; i < b.N; i++ {
		_ = q.IsEmpty()
	}
}

func BenchmarkQueue_Offer(b *testing.B) {
	q := Queue{}
	for i := 0; i < b.N; i++ {
		isAdded := q.Offer(i)
		if !isAdded {
			fmt.Println("unable to add element to the queue")
		}
	}
}

func BenchmarkQueue_Peek(b *testing.B) {
	q := Queue{}
	for i := 0; i < b.N; i++ {
		_ = q.Peek()
	}
}

func BenchmarkQueue_Poll(b *testing.B) {
	q := Queue{}
	for i := 0; i < b.N; i++ {
		_, _ = q.Poll()
	}
}

func BenchmarkQueue_Remove(b *testing.B) {
	q := Queue{}
	for i := 0; i < 10; i++ {
		_, _ = q.Remove()
	}
}

func BenchmarkQueue_Size(b *testing.B) {
	q := Queue{}
	for i := 0; i < b.N; i++ {
		_ = q.Size()
	}
}
