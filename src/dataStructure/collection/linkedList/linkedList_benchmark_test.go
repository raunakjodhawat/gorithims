//nolint
package linkedlist

import "testing"

func BenchmarkList_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Add(i)
	}
}

func BenchmarkList_AddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.AddAll([]int{i})
	}
}

func BenchmarkList_AddFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.AddFirst(i)
	}
}

func BenchmarkList_AddLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.AddLast(i)
	}
}

func BenchmarkList_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Contains(i)
	}
}

func BenchmarkList_Element(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Element()
	}
}

func BenchmarkList_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Get(i)
	}
}

func BenchmarkList_GetFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.GetFirst()
	}
}

func BenchmarkList_GetLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.GetLast()
	}
}

func BenchmarkList_IndexOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.IndexOf(i)
	}
}

func BenchmarkList_LastIndexOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.LastIndexOf(i)
	}
}

func BenchmarkList_ListIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.ListIterator(i)
	}
}

func BenchmarkList_Offer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Offer(i)
	}
}

func BenchmarkList_OfferFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.OfferFirst(i)
	}
}

// Down
func BenchmarkList_OfferLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.OfferLast(i)
	}
}

func BenchmarkList_Peek(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Peek()
	}
}

func BenchmarkList_PeekFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.PeekFirst()
	}
}

func BenchmarkList_PeekLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.PeekLast()
	}
}

func BenchmarkList_Poll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Poll()
	}
}

func BenchmarkList_PollFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.OfferFirst(i)
	}
}

func BenchmarkList_PollLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.PollLast()
	}
}

func BenchmarkList_Pop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.PollLast()
	}
}

func BenchmarkList_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Push(i)
	}
}

func BenchmarkList_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Remove(i)
	}
}

func BenchmarkList_Remove2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Remove(i, 0)
	}
}

func BenchmarkList_Remove3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Remove(i, l.Size()-1)
	}
}

func BenchmarkList_RemoveFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.RemoveFirst()
	}
}

func BenchmarkList_RemoveFirstOccurrence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.RemoveFirstOccurrence(i)
	}
}

func BenchmarkList_RemoveLast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.RemoveLast()
	}
}

func BenchmarkList_RemoveLastOccurrence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.RemoveLastOccurrence(i)
	}
}

func BenchmarkList_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Set(i, 100)
	}
}

func BenchmarkList_Size(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.Size()
	}
}

func BenchmarkList_ToArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{}
		l.ToArray()
	}
}
