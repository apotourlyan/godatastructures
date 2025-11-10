package lists

/*
Testing Strategy
================

The LinkedList test suite uses a comprehensive approach to verify correctness:

1. Edge Cases
   - Empty lists
   - Single-element lists
   - Boundary conditions (first/last elements)

2. Core Operations
   - All mutation operations (Add, Remove, InsertAt, RemoveAt)
   - All query operations (GetAt, IndexOf, Contains, First, Last)
   - All state operations (IsEmpty, Size)

3. Invariant Verification
   Each test verifies critical invariants after operations:
   - Size matches expected value
   - Head points to first element
   - Tail points to last element
   - Tail.Next is always nil (no cycles)

4. Order Preservation
   Special "Order" tests verify that:
   - Elements maintain insertion order
   - List structure is not corrupted
   - All nodes are properly linked

5. Error Conditions
   Tests verify proper error handling for:
   - Invalid indices (negative, out of range)
   - Operations on empty lists

Test Organization
=================

Tests are organized by operation and scenario:
- TestLinkedList_<Operation>_<Scenario>

Examples:
- TestLinkedList_Add_OneValue_EmptyList
- TestLinkedList_Remove_LastValue_TwoElementList
- TestLinkedList_InsertAt_Middle_ManyElementList

This naming convention makes it immediately clear:
1. What operation is being tested
2. What the test scenario is
3. What state the list starts in

Benefits:
- Easy to identify missing test cases
- Clear failure messages
- Self-documenting test suite

Test Coverage
=============

Coverage by operation:

Constructor (NewLinkedList):
  ✓ Empty list
  ✓ Single value
  ✓ Multiple values
  ✓ Order preservation

Add:
  ✓ Add to empty list (1 and 2 values)
  ✓ Add to non-empty list (1 and 2 values)
  ✓ Order preservation

Remove:
  ✓ Remove from empty list
  ✓ Remove single element (list becomes empty)
  ✓ Remove first of two elements
  ✓ Remove last of two elements
  ✓ Remove middle element
  ✓ Remove non-existent element
  ✓ Order preservation after removal

InsertAt:
  ✓ Negative index (error)
  ✓ Invalid index (error)
  ✓ Insert into empty list (index 0)
  ✓ Insert at start (single and many elements)
  ✓ Insert at end/append (single and many elements)
  ✓ Insert in middle
  ✓ Order preservation after insertion

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

First/Last/IsEmpty/Size:
  ✓ On empty list
  ✓ On non-empty list
*/

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Purpose: Verify empty list creation
//
// Verifies: size == 0, head == nil, tail == nil
func TestLinkedList_NewLinkedList_Empty(t *testing.T) {
	l := NewLinkedList[int]()
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify single value constructor
//
// Verifies: size == 1, head == tail, values correct
func TestLinkedList_NewLinkedList_OneValue(t *testing.T) {
	l := NewLinkedList(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify multiple values constructor
//
// Verifies: size correct, head and tail point to correct elements
func TestLinkedList_NewLinkedList_ManyValues(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify constructor maintains insertion order
//
// Verifies: All elements accessible in correct order by traversing nodes
func TestLinkedList_NewLinkedList_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Purpose: Verify Add to empty list
//
// Verifies: size == 1, head == tail, element stored correctly
func TestLinkedList_Add_OneValue_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	test.GotWant(t, l.size, 1)
	test.GotWant(t, l.head, l.tail)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Add two values to empty list
//
// Verifies: size == 2, head and tail correct, order preserved
func TestLinkedList_Add_TwoValues_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(2)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Add to non-empty list
//
// Verifies: size increases, tail updated, head unchanged
func TestLinkedList_Add_OneValue_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.Add(3)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Add multiple values to non-empty list
//
// Verifies: size increases correctly, tail updated, head unchanged
func TestLinkedList_Add_TwoValues_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2)
	l.Add(3)
	l.Add(4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Add maintains insertion order
//
// Verifies: All added elements accessible in correct order
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

// Purpose: Verify Remove from empty list
//
// Verifies: Returns false, list remains empty
func TestLinkedList_Remove_OneValue_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	r := l.Remove(1)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify Remove single element (list becomes empty)
//
// Verifies: Returns true, size == 0, head == nil, tail == nil
func TestLinkedList_Remove_OneValue_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	r := l.Remove(1)
	test.GotWant(t, r, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify Remove first of two elements
//
// Verifies: Returns true, size == 1, head == tail, correct value remains
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

// Purpose: Verify Remove last of two elements
//
// Verifies: Returns true, size == 1, head == tail, tail updated correctly
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

// Purpose: Verify Remove middle element
//
// Verifies: Returns true, element removed, size decreased, head/tail unchanged
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

// Purpose: Verify Remove non-existent element
//
// Verifies: Returns false, list unchanged, size unchan
func TestLinkedList_Remove_NonExistent_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	r := l.Remove(10)
	test.GotWant(t, r, false)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Remove maintains order
//
// Verifies: After removing middle element, remaining elements in correct order
func TestLinkedList_Remove_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 99, 3, 4)
	l.Remove(99)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Purpose: Verify InsertAt with negative index
//
// Verifies: Returns ErrorIndexOutOfRange, list unchanged
func TestLinkedList_InsertAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	err := l.InsertAt(-1, 1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify InsertAt with invalid index
//
// Verifies: Returns ErrorIndexOutOfRange, list unchanged
func TestLinkedList_InsertAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(4, 4)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify InsertAt into empty list
//
// Verifies: Element inserted, size == 1, head == tail
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

// Purpose: Verify InsertAt start of single-element list
//
// Verifies: Element inserted at head, size == 2, head updated, tail unchanged
func TestLinkedList_InsertAt_Start_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.InsertAt(0, 0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 0)
	test.GotWant(t, l.tail.Value, 1)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify InsertAt end of single-element list (append)
//
// Verifies: Element inserted at tail, size == 2, head unchanged, tail updated
func TestLinkedList_InsertAt_End_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.InsertAt(1, 2)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify InsertAt start of multi-element list
//
// Verifies: Element inserted at head, size increased, head updated
func TestLinkedList_InsertAt_Start_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(0, 0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 0)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify InsertAt end of multi-element list (append)
//
// Verifies: Element inserted at tail, size increased,
func TestLinkedList_InsertAt_End_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.InsertAt(3, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify InsertAt middle position
//
// Verifies: Element inserted at correct position, size increased
func TestLinkedList_InsertAt_Middle_ManyElementList(t *testing.T) {
	l := NewLinkedList(1, 2, 4)
	err := l.InsertAt(2, 3)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify InsertAt maintains order
//
// Verifies: After insertion, all elements in correct order
func TestLinkedList_InsertAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 4, 5)
	l.InsertAt(2, 3)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Purpose: Verify RemoveAt with negative index
//
// Verifies: Returns ErrorIndexOutOfRange, list unchanged
func TestLinkedList_RemoveAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	err := l.RemoveAt(-1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify RemoveAt with invalid index
//
// Verifies: Returns ErrorIndexOutOfRange, list unchanged
func TestLinkedList_RemoveAt_InvalidIndex(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(3)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify RemoveAt single element (list becomes empty)
//
// Verifies: Element removed, size == 0, head == nil, tail == nil
func TestLinkedList_RemoveAt_OneElementList(t *testing.T) {
	l := NewLinkedList(1)
	err := l.RemoveAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify RemoveAt start of list
//
// Verifies: First element removed, size decreased, head updated
func TestLinkedList_RemoveAt_Start(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 2)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify RemoveAt end of list
//
// Verifies: Last element removed, size decreased, tail updated
func TestLinkedList_RemoveAt_End(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(2)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 2)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify RemoveAt middle position
//
// Verifies: Middle element removed, size decreased, head/tail unchanged
func TestLinkedList_RemoveAt_Middle(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	err := l.RemoveAt(1)
	test.GotWant(t, err, nil)
	test.GotWant(t, l.size, 2)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify RemoveAt maintains order
//
// Verifies: After removal, remaining elements in correct order
func TestLinkedList_RemoveAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 99, 3, 4)
	l.RemoveAt(2)

	node := l.head
	for i := range l.size {
		test.GotWant(t, node.Value, i+1)
		node = node.Next
	}
}

// Purpose: Verify GetAt with negative index
//
// Verifies: Returns ErrorIndexOutOfRange, zero value returned
func TestLinkedList_GetAt_NegativeIndex(t *testing.T) {
	l := NewLinkedList[int]()
	v, err := l.GetAt(-1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, v, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify GetAt with invalid index
//
// Verifies: Returns ErrorIndexOutOfRange, zero value returned
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

// Purpose: Verify GetAt first element
//
// Verifies: Returns correct value, list unchanged
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

// Purpose: Verify GetAt last element
//
// Verifies: Returns correct value, list unchanged
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

// Purpose: Verify GetAt middle element
//
// Verifies: Returns correct value, list unchanged
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

// Purpose: Verify GetAt for all elements
//
// Verifies: All elements accessible in correct order by index
func TestLinkedList_GetAt_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for i := range l.size {
		v, err := l.GetAt(i)
		test.GotWant(t, err, nil)
		test.GotWant(t, v, i+1)
	}
}

// Purpose: Verify IndexOf in empty list
//
// Verifies: Returns -1, list unchanged
func TestLinkedList_IndexOf_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	i := l.IndexOf(99)
	test.GotWant(t, i, -1)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify IndexOf for non-existing element
//
// Verifies: Returns -1, list unchanged
func TestLinkedList_IndexOf_NonExisting(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	i := l.IndexOf(99)
	test.GotWant(t, i, -1)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify IndexOf for existing element
//
// Verifies: Returns correct index of first occurrence
func TestLinkedList_IndexOf_Existing(t *testing.T) {
	l := NewLinkedList(1, 2, 3)
	i := l.IndexOf(1)
	test.GotWant(t, i, 0)
	test.GotWant(t, l.size, 3)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 3)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify IndexOf finds all elements
//
// Verifies: All elements found at correct indices
func TestLinkedList_IndexOf_Order(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for j := range l.size {
		i := l.IndexOf(j + 1)
		test.GotWant(t, i, j)
	}
}

// Purpose: Verify Contains in empty list
//
// Verifies: Returns false, list unchanged
func TestLinkedList_Contains_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	c := l.Contains(99)
	test.GotWant(t, c, false)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify Contains for non-existing element
//
// Verifies: Returns false, list unchanged
func TestLinkedList_Contains_NonExisting(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	c := l.Contains(99)
	test.GotWant(t, c, false)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Contains for existing element
//
// Verifies: Returns true, list unchanged
func TestLinkedList_Contains_Existing(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	c := l.Contains(4)
	test.GotWant(t, c, true)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Contains for all elements
//
// Verifies: All elements found correctly
func TestLinkedList_Contains_All(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)

	for i := range l.size {
		c := l.Contains(i + 1)
		test.GotWant(t, c, true)
	}
}

// Purpose: Verify First on empty list
//
// Verifies: Returns ErrorEmptyList, zero value returned
func TestLinkedList_First_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	f, err := l.First()
	test.GotWantError(t, err, ErrorEmptyList)
	test.GotWant(t, f, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify First on non-empty list
//
// Verifies: Returns first element, list unchanged
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

// Purpose: Verify Last on empty list
//
// Verifies: Returns ErrorEmptyList, zero value returned
func TestLinkedList_Last_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	la, err := l.Last()
	test.GotWantError(t, err, ErrorEmptyList)
	test.GotWant(t, la, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify Last on non-empty list
//
// Verifies: Returns last element, list unchanged
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

// Purpose: Verify IsEmpty on empty list
//
// Verifies: Returns true
func TestLinkedList_IsEmpty_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	e := l.IsEmpty()
	test.GotWant(t, e, true)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify IsEmpty on non-empty list
//
// Verifies: Returns false
func TestLinkedList_IsEmpty_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	e := l.IsEmpty()
	test.GotWant(t, e, false)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}

// Purpose: Verify Size on empty list
//
// Verifies: Returns 0
func TestLinkedList_Size_EmptyList(t *testing.T) {
	l := NewLinkedList[int]()
	s := l.Size()
	test.GotWant(t, s, 0)
	test.GotWant(t, l.size, 0)
	test.GotWant(t, l.head, nil)
	test.GotWant(t, l.tail, nil)
}

// Purpose: Verify Size on non-empty list
//
// Verifies: Returns correct count of elements
func TestLinkedList_Size_NonEmptyList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4)
	s := l.Size()
	test.GotWant(t, s, 4)
	test.GotWant(t, l.size, 4)
	test.GotWant(t, l.head.Value, 1)
	test.GotWant(t, l.tail.Value, 4)
	test.GotWant(t, l.tail.Next, nil)
}
