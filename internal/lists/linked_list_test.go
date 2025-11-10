package linkedlists

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

func TestLinkedList_NewLinkedList_Empty(t *testing.T) {
	l := NewLinkedList[int]()
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_NewLinkedList_OneValue(t *testing.T) {
	l := NewLinkedList(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_NewLinkedList_ManyValues(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_NewLinkedList_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

func TestLinkedList_Add_OneValue_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Add_TwoValues_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(2)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Add_OneValue_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.Add(3)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Add_TwoValues_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.Add(3)
	l.Add(4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Add_Order(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

func TestLinkedList_Remove_OneValue_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	r := l.Remove(1)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_Remove_OneValue_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	r := l.Remove(1)
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_Remove_FirstValue_TwoElementList(t *testing.T) {
	l := NewLinkedList(1, 2)
	r := l.Remove(1)
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 2)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Remove_LastValue_TwoElementList(t *testing.T) {
	l := NewLinkedList(1, 2)
	r := l.Remove(2)
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Remove_MidValue_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4, 5)
	r := l.Remove(3)
	c := l.Contains(3)
	test.GotWant(t, r, true)
	test.GotWant(t, c, false)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 5)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Remove_NonExistent_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	r := l.Remove(10)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Remove_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 99, 3, 4)
	l.Remove(99)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

func TestLinkedList_InsertAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	err := l.InsertAt(-1, 1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_InsertAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(4, 4)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_InsertAt_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	err := l.InsertAt(0, 1)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_InsertAt_Start_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.InsertAt(0, 0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 0)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_InsertAt_End_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.InsertAt(1, 2)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_InsertAt_Start_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(0, 0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 0)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_InsertAt_End_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(3, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_InsertAt_Middle_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 4)
	err := l.InsertAt(2, 3)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_InsertAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 4, 5)
	l.InsertAt(2, 3)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

func TestLinkedList_RemoveAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	err := l.RemoveAt(-1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_RemoveAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(3)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_RemoveAt_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.RemoveAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_RemoveAt_Start(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 2)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_RemoveAt_End(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(2)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_RemoveAt_Middle(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(1)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_RemoveAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 99, 3, 4)
	l.RemoveAt(2)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

func TestLinkedList_GetAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	v, err := l.GetAt(-1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, v, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_GetAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	v, err := l.GetAt(3)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, v, 0)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_GetAt_Start(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	v, err := l.GetAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, v, 1)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_GetAt_End(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	v, err := l.GetAt(2)
	test.GotWant(t, err, nil)
	test.GotWant(t, v, 3)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_GetAt_Middle(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	v, err := l.GetAt(1)
	test.GotWant(t, err, nil)
	test.GotWant(t, v, 2)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_GetAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for i := range l.size {
		v, err := l.GetAt(i)
		test.GotWant(t, err, nil)
		test.GotWant(t, v, i+1)
	}
}

func TestLinkedList_IndexOf_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	i := l.IndexOf(99)
	test.GotWant(t, i, -1)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_IndexOf_NonExisting(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	i := l.IndexOf(99)
	test.GotWant(t, i, -1)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_IndexOf_Existing(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	i := l.IndexOf(1)
	test.GotWant(t, i, 0)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_IndexOf_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for j := range l.size {
		i := l.IndexOf(j + 1)
		test.GotWant(t, i, j)
	}
}

func TestLinkedList_Contains_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	c := l.Contains(99)
	test.GotWant(t, c, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_Contains_NonExisting(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	c := l.Contains(99)
	test.GotWant(t, c, false)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Contains_Existing(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	c := l.Contains(4)
	test.GotWant(t, c, true)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Contains_All(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for i := range l.size {
		c := l.Contains(i + 1)
		test.GotWant(t, c, true)
	}
}

func TestLinkedList_First_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	f, err := l.First()
	test.GotWantError(t, err, ErrorEmptyList)
	test.GotWant(t, f, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_First_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	f, err := l.First()
	test.GotWant(t, err, nil)
	test.GotWant(t, f, 1)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Last_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	la, err := l.Last()
	test.GotWantError(t, err, ErrorEmptyList)
	test.GotWant(t, la, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_Last_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	la, err := l.Last()
	test.GotWant(t, err, nil)
	test.GotWant(t, la, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_IsEmpty_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	e := l.IsEmpty()
	test.GotWant(t, e, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_IsEmpty_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	e := l.IsEmpty()
	test.GotWant(t, e, false)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

func TestLinkedList_Size_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	s := l.Size()
	test.GotWant(t, s, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

func TestLinkedList_Size_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	s := l.Size()
	test.GotWant(t, s, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}
