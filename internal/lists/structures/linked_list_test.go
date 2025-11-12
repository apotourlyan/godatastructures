package structures

/*
Test Coverage
=============
Constructor (NewBasicLinkedList):
  ✓ Empty list
  ✓ Single value
  ✓ Multiple values
  ✓ Order preservation

Constructor (NewLinkedList):
  ✓ Empty list
  ✓ Single value
  ✓ Multiple values
  ✓ Order preservation

AddFirst:
  ✓ Add to empty list (1 and 2 values)
  ✓ Add to non-empty list (1 and 2 values)
  ✓ Order preservation

AddLast:
  ✓ Add to empty list (1 and 2 values)
  ✓ Add to non-empty list (1 and 2 values)
  ✓ Order preservation

RemoveFirst:
  ✓ Remove from empty list
  ✓ Remove from one-element list
  ✓ Remove from two-element list
  ✓ Order preservation after removal

RemoveLast:
  ✓ Remove from empty list
  ✓ Remove from one-element list
  ✓ Remove from two-element list
  ✓ Order preservation after removal

First/Last/IsEmpty/Size:
  ✓ On empty list
  ✓ On non-empty list

InsertAt:
  ✓ Negative index (error)
  ✓ Invalid index (error)
  ✓ Insert into empty list (index 0)
  ✓ Insert at start (single and many elements)
  ✓ Insert at end/append (single and many elements)
  ✓ Insert in middle
  ✓ Order preservation after insertion

UpdateAt:
  ✓ Negative index (error)
  ✓ Invalid index (error)
  ✓ Update at start
  ✓ Update at end
  ✓ Update in middle
  ✓ Order preservation after update

RemoveAt:
  ✓ Negative index (error)
  ✓ Invalid index (error)
  ✓ Remove single element (list becomes empty)
  ✓ Remove at start
  ✓ Remove at end
  ✓ Remove in middle
  ✓ Order preservation after removal

GetAt:
  ✓ Negative index (error)
  ✓ Invalid index (error)
  ✓ Get at start
  ✓ Get at end
  ✓ Get in middle
  ✓ Get all elements in order

IndexOf:
  ✓ Search in empty list
  ✓ Search for non-existent element
  ✓ Search for existing element
  ✓ Find all elements in order

Contains:
  ✓ Search in empty list
  ✓ Search for non-existent element
  ✓ Search for existing element
  ✓ Verify all elements present

Remove:
  ✓ Remove from empty list
  ✓ Remove single element (list becomes empty)
  ✓ Remove first of two elements
  ✓ Remove last of two elements
  ✓ Remove middle element
  ✓ Remove non-existent element
  ✓ Order preservation after removal

Update:
  ✓ Update in empty list
  ✓ Update non-existent element
  ✓ Update existing element
  ✓ Update elements in order
*/

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Verifies empty basic list creation
func TestLinkedList_NewBasicLinkedList_Empty(t *testing.T) {
	l := NewBasicLinkedList[int]()
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies single value basic list creation
func TestLinkedList_NewBasicLinkedList_OneValue(t *testing.T) {
	l := NewBasicLinkedList(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies multiple values basic list creation
func TestLinkedList_NewBasicLinkedList_ManyValues(t *testing.T) {
	l := NewBasicLinkedList(1, 2, 3, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies empty list creation
func TestLinkedList_NewLinkedList_Empty(t *testing.T) {
	l := NewLinkedList[int]()
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies single value list creation
func TestLinkedList_NewLinkedList_OneValue(t *testing.T) {
	l := NewLinkedList(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies multiple values list creation
func TestLinkedList_NewLinkedList_ManyValues(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies constructor maintains insertion order
func TestLinkedList_NewLinkedList_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies prepending a single value to an empty list
func TestLinkedList_AddFirst_OneValue_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddFirst(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies prepending two values to an empty list
func TestLinkedList_AddFirst_TwoValues_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddFirst(1)
	l.AddFirst(2)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 2)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies prepending a single value to a non-empty list
func TestLinkedList_AddFirst_OneValue_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.AddFirst(0)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 0)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies prepending two values to a non-empty list
func TestLinkedList_AddFirst_TwoValues_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.AddFirst(0)
	l.AddFirst(-1)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, -1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies prepending maintains insertion order
func TestLinkedList_AddFirst_Order(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddFirst(3)
	l.AddFirst(2)
	l.AddFirst(1)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies appending a single value to an empty list
func TestLinkedList_AddLast_OneValue_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies appending two values to an empty list
func TestLinkedList_AddLast_TwoValues_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(1)
	l.AddLast(2)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies appending a single value to a non-empty list
func TestLinkedList_AddLast_OneValue_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.AddLast(3)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies appending two values to a non-empty list
func TestLinkedList_AddLast_TwoValues_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.AddLast(3)
	l.AddLast(4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies appending maintains insertion order
func TestLinkedList_AddLast_Order(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies removing from an empty list
func TestLinkedList_RemoveFirst_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	r := l.RemoveFirst()
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing from a one-element list
func TestLinkedList_RemoveFirst_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	r := l.RemoveFirst()
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing from a two-element list
func TestLinkedList_RemoveFirst_TwoElementList(t *testing.T) {
	l := NewLinkedList(1, 2)
	r := l.RemoveFirst()
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 2)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies order after removal
func TestLinkedList_RemoveFirst_Order(t *testing.T) {
	l := NewLinkedList(0, 1, 2, 3)
	l.RemoveFirst()

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies removing from an empty list
func TestLinkedList_RemoveLast_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	r := l.RemoveLast()
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing from a one-element list
func TestLinkedList_RemoveLast_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	r := l.RemoveLast()
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing from a two-element list
func TestLinkedList_RemoveLast_TwoElementList(t *testing.T) {
	l := NewLinkedList(1, 2)
	r := l.RemoveLast()
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies order after removal
func TestLinkedList_RemoveLast_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	l.RemoveLast()

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies getting first in an empty list
func TestLinkedList_First_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	f, err := l.First()
	test.GotWantError(t, err, ErrorEmptyList)
	test.GotWant(t, f, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies getting first in a non-empty list
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

// Verifies getting last in an empty list
func TestLinkedList_Last_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	la, err := l.Last()
	test.GotWantError(t, err, ErrorEmptyList)
	test.GotWant(t, la, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies getting last in a non-empty list
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

// Verifies empty list
func TestLinkedList_IsEmpty_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	e := l.IsEmpty()
	test.GotWant(t, e, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies non-empty list
func TestLinkedList_IsEmpty_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	e := l.IsEmpty()
	test.GotWant(t, e, false)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies size in an empty list
func TestLinkedList_Size_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	s := l.Size()
	test.GotWant(t, s, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies size in a non-empty list
func TestLinkedList_Size_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	s := l.Size()
	test.GotWant(t, s, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies inserting at negative index
func TestLinkedList_InsertAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	err := l.InsertAt(-1, 1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies inserting at invalid index
func TestLinkedList_InsertAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(4, 4)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies inserting in an empty list
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

// Verifies inserting at the start of a one-element list
func TestLinkedList_InsertAt_Start_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.InsertAt(0, 0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 0)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies inserting at the end of a one-element list
func TestLinkedList_InsertAt_End_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.InsertAt(1, 2)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies inserting at the start of a multi-element list
func TestLinkedList_InsertAt_Start_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(0, 0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 0)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies inserting at the end of a multi-element list
func TestLinkedList_InsertAt_End_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(3, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies inserting at the middle of a multi-element list
func TestLinkedList_InsertAt_Middle_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 4)
	err := l.InsertAt(2, 3)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies order after insertion
func TestLinkedList_InsertAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 4, 5)
	l.InsertAt(2, 3)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies updating at negative index
func TestLinkedList_UpdateAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	old, err := l.UpdateAt(-1, 0)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, old, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies updating at invalid index
func TestLinkedList_UpdateAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	old, err := l.UpdateAt(3, 4)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, old, 0)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies updating at the start of a multi-element list
func TestLinkedList_UpdateAt_Start(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	old, err := l.UpdateAt(0, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, old, 1)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 4)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies updating at the end of a multi-element list
func TestLinkedList_UpdateAt_End(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	old, err := l.UpdateAt(2, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, old, 3)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies updating the middle of a multi-element list
func TestLinkedList_UpdateAt_Middle(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	old, err := l.UpdateAt(1, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, old, 2)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies updating in order
func TestLinkedList_UpdateAt_Order(t *testing.T) {
	l := NewLinkedList(0, 1, 2)
	for i := range l.size {
		old, _ := l.UpdateAt(i, i+1)
		test.GotWant(t, old, i)
	}

	for i := range l.size {
		new, _ := l.GetAt(i)
		test.GotWant(t, new, i+1)
	}
}

// Verifies removing at negative index
func TestLinkedList_RemoveAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	err := l.RemoveAt(-1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing at invalid index
func TestLinkedList_RemoveAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(3)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies removing from a one-element list
func TestLinkedList_RemoveAt_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.RemoveAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing at the start of a multi-element list
func TestLinkedList_RemoveAt_Start(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 2)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies removing at the end of a multi-element list
func TestLinkedList_RemoveAt_End(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(2)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies removing at the middle of a multi-element list
func TestLinkedList_RemoveAt_Middle(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(1)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies order after removal
func TestLinkedList_RemoveAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 99, 3, 4)
	l.RemoveAt(2)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies getting at negative index
func TestLinkedList_GetAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	v, err := l.GetAt(-1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, v, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies getting at invalid index
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

// Verifies getting at the start of a multi-element list
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

// Verifies getting at the end of a multi-element list
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

// Verifies getting at the middle of a multi-element list
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

// Verifies all elements are accessible in the correct order by index
func TestLinkedList_GetAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for i := range l.size {
		v, err := l.GetAt(i)
		test.GotWant(t, err, nil)
		test.GotWant(t, v, i+1)
	}
}

// Verifies getting an index of any element in an empty list
func TestLinkedList_IndexOf_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	i := l.IndexOf(99)
	test.GotWant(t, i, -1)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies getting an index of a non-existing element
func TestLinkedList_IndexOf_NonExisting(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	i := l.IndexOf(99)
	test.GotWant(t, i, -1)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies getting an index of an existing element
func TestLinkedList_IndexOf_Existing(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	i := l.IndexOf(1)
	test.GotWant(t, i, 0)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies all elements are at the correct indices
func TestLinkedList_IndexOf_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for j := range l.size {
		i := l.IndexOf(j + 1)
		test.GotWant(t, i, j)
	}
}

// Verifies existence in an empty list
func TestLinkedList_Contains_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	c := l.Contains(99)
	test.GotWant(t, c, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies existence of a non-existing element
func TestLinkedList_Contains_NonExisting(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	c := l.Contains(99)
	test.GotWant(t, c, false)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies existence of a existing element
func TestLinkedList_Contains_Existing(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	c := l.Contains(4)
	test.GotWant(t, c, true)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies all elements are existing
func TestLinkedList_Contains_All(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for i := range l.size {
		c := l.Contains(i + 1)
		test.GotWant(t, c, true)
	}
}

// Verifies removing from an empty list
func TestLinkedList_Remove_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	r := l.Remove(1)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing from a one-element list
func TestLinkedList_Remove_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	r := l.Remove(1)
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies removing the first element from a two-element list
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

// Verifies removing the last element from a two-element list
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

// Verifies removing a mid element from a multi-element list
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

// Verifies removing a non-existent element
func TestLinkedList_Remove_NonExistent_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	r := l.Remove(10)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies order after removal
func TestLinkedList_Remove_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 99, 3, 4)
	l.Remove(99)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Verifies updating in an empty list
func TestLinkedList_Update_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	r := l.Update(1, 2)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Verifies updating a non-existing element
func TestLinkedList_Update_NonExisting(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	r := l.Update(0, 1)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies updating an existing element
func TestLinkedList_Update_Existing(t *testing.T) {
	l := NewLinkedList(0, 2, 3)
	r := l.Update(0, 1)
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Verifies updating in order
func TestLinkedList_Update_Order(t *testing.T) {
	l := NewLinkedList(0, 0, 0, 0)
	for i := range l.size {
		r := l.Update(0, i+1)
		test.GotWant(t, r, true)
	}

	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}
