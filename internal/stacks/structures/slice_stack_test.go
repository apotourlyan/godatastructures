package structures

/*
Test Coverage
=============
Constructor (NewSliceStack):
  ✓ Empty stack
  ✓ Single value
  ✓ Multiple values

Push:
  ✓ Single value to empty stack
  ✓ Single value to non-empty stack
  ✓ Multiple values to empty stack
  ✓ Multiple values to non-empty stack

Pop:
  ✓ Single value from empty stack
  ✓ Single value from non-empty stack
  ✓ Multiple values from non-empty stack

Push/Pop:
  ✓ LIFO order
  ✓ Reusable after emptying the stack

Peek:
  ✓ Empty stack
  ✓ Non-empty stack (single peek)
  ✓ Non-empty stack (multiple peeks)

IsEmpty/Size:
  ✓ Empty stack
  ✓ Non-empty stack
*/

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Verifies the creation of an empty stack
func TestSliceStack_NewSliceStack_Empty(t *testing.T) {
	s := NewSliceStack[int]()
	test.GotWant(t, s.Size(), 0)
	test.GotWant(t, s.IsEmpty(), true)
}

// Verifies the creation of one-element stack
func TestSliceStack_NewSliceStack_OneValue(t *testing.T) {
	s := NewSliceStack(1)
	test.GotWant(t, s.Size(), 1)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the creation of multi-element stack
func TestSliceStack_NewSliceStack_ManyValues(t *testing.T) {
	s := NewSliceStack(1, 2, 3)
	test.GotWant(t, s.Size(), 3)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the pushing of an element in an empty stack
func TestSliceStack_Push_OneElement_EmptyStack(t *testing.T) {
	s := NewSliceStack[int]()
	s.Push(1)
	p, _ := s.Peek()
	test.GotWant(t, p, 1)
	test.GotWant(t, s.Size(), 1)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the pushing of multiple elements in an empty stack
func TestSliceStack_Push_ManyElements_EmptyStack(t *testing.T) {
	s := NewSliceStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	p, _ := s.Peek()
	test.GotWant(t, p, 3)
	test.GotWant(t, s.Size(), 3)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the pushing of an element in a non-empty stack
func TestSliceStack_Push_OneElement_NonEmptyStack(t *testing.T) {
	s := NewSliceStack(1, 2, 3)
	s.Push(4)
	p, _ := s.Peek()
	test.GotWant(t, p, 4)
	test.GotWant(t, s.Size(), 4)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the pushing of multiple elements in a non-empty stack
func TestSliceStack_Push_ManyElements_NonEmptyStack(t *testing.T) {
	s := NewSliceStack(1, 2, 3)
	s.Push(4)
	s.Push(5)
	p, _ := s.Peek()
	test.GotWant(t, p, 5)
	test.GotWant(t, s.Size(), 5)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the poping an element from an empty stack
func TestSliceStack_Pop_OneElement_EmptyStack(t *testing.T) {
	s := NewSliceStack[int]()
	d, err := s.Pop()
	test.GotWantError(t, err, ErrorEmptyStack)
	test.GotWant(t, d, 0)
	test.GotWant(t, s.Size(), 0)
	test.GotWant(t, s.IsEmpty(), true)
}

// Verifies the poping an element from a non-empty stack
func TestSliceStack_Pop_OneElement_NonEmptyStack(t *testing.T) {
	s := NewSliceStack(1, 2, 3)
	d, err := s.Pop()
	test.GotWant(t, err, nil)
	test.GotWant(t, d, 3)
	p, _ := s.Peek()
	test.GotWant(t, p, 2)
	test.GotWant(t, s.Size(), 2)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the poping multiple elements in a non-empty stack
func TestSliceStack_Pop_ManyElements_NonEmptyStack(t *testing.T) {
	s := NewSliceStack(1, 2, 3)
	s.Pop()
	s.Pop()
	d, err := s.Pop()
	test.GotWant(t, err, nil)
	test.GotWant(t, d, 1)
	test.GotWant(t, s.Size(), 0)
	test.GotWant(t, s.IsEmpty(), true)
}

// Verifies Last-In-First-Out element order
func TestSliceStack_PushPop_Order(t *testing.T) {
	s := NewSliceStack[int]()

	for i := range 5 {
		s.Push(i + 1)
		p, _ := s.Peek()
		test.GotWant(t, p, i+1)
	}

	for i := range 5 {
		p, _ := s.Peek()
		test.GotWant(t, p, 5-i)
		d, _ := s.Pop()
		test.GotWant(t, d, 5-i)
	}
}

// Verifies the stack is reusable
func TestSliceStack_PushPop_Reusability(t *testing.T) {
	s := NewSliceStack[int]()
	s.Push(1)
	s.Pop()
	test.GotWant(t, s.IsEmpty(), true)
	s.Push(2)
	p, _ := s.Peek()
	test.GotWant(t, p, 2)
}

// Verifies peeking into an empty stack
func TestSliceStack_Peek_EmptyStack(t *testing.T) {
	s := NewSliceStack[int]()
	p, err := s.Peek()
	test.GotWantError(t, err, ErrorEmptyStack)
	test.GotWant(t, p, 0)
	test.GotWant(t, s.Size(), 0)
	test.GotWant(t, s.IsEmpty(), true)
}

// Verifies peeking into an non-empty stack
func TestSliceStack_Peek_NonEmptyStack(t *testing.T) {
	s := NewSliceStack(1, 2, 3)
	p, err := s.Peek()
	test.GotWant(t, err, nil)
	test.GotWant(t, p, 3)
	test.GotWant(t, s.Size(), 3)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies peeking multiple times into an non-empty stack
func TestSliceStack_Peek_Many(t *testing.T) {
	s := NewSliceStack(1, 2, 3)

	for range 3 {
		p, err := s.Peek()
		test.GotWant(t, err, nil)
		test.GotWant(t, p, 3)
	}
}

// Verifies the empty state of an empty stack
func TestSliceStack_IsEmpty_EmptyStack(t *testing.T) {
	s := NewSliceStack[int]()
	test.GotWant(t, s.IsEmpty(), true)
}

// Verifies the empty state of an non-empty stack
func TestSliceStack_IsEmpty_NonEmptyStack(t *testing.T) {
	s := NewSliceStack(1)
	test.GotWant(t, s.IsEmpty(), false)
}

// Verifies the size of an empty stack
func TestSliceStack_Size_EmptyStack(t *testing.T) {
	s := NewSliceStack[int]()
	test.GotWant(t, s.Size(), 0)
}

// Verifies the size of an non-empty stack
func TestSliceStack_Size_NonEmptyStack(t *testing.T) {
	s := NewSliceStack(1, 2, 3)
	test.GotWant(t, s.Size(), 3)
}
