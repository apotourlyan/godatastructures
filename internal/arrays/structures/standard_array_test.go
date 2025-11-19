package structures

/*
Test Coverage
=============
Constructor (NewStandardArray):
  ✓ Empty list
  ✓ Single value
  ✓ Multiple values
  ✓ Order preservation

GetAt:
  ✓ Negative index (error)
  ✓ Invalid index (error)
  ✓ Get at start
  ✓ Get at end
  ✓ Get in middle
  ✓ Get all elements in order

UpdateAt:
  ✓ Negative index (error)
  ✓ Invalid index (error)
  ✓ Update at start
  ✓ Update at end
  ✓ Update in middle
  ✓ Order preservation after update

IsEmpty/Size:
  ✓ On empty list
  ✓ On non-empty list
*/

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Verifies the creation of an empty array
func TestStandardArray_NewStandardArray_Empty(t *testing.T) {
	a := NewStandardArray[int]()
	test.GotWant(t, a.Size(), 0)
	test.GotWant(t, a.IsEmpty(), true)
}

// Verifies the creation of one-element array
func TestStandardArray_NewStandardArray_OneValue(t *testing.T) {
	a := NewStandardArray(1)
	test.GotWant(t, a.Size(), 1)
	test.GotWant(t, a.IsEmpty(), false)
}

// Verifies the creation of multi-element array
func TestStandardArray_NewStandardArray_ManyValues(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	test.GotWant(t, a.Size(), 3)
	test.GotWant(t, a.IsEmpty(), false)
}

// Verifies constructor maintains insertion order
func TestStandardArray_NewStandardArray_Order(t *testing.T) {
	a := NewStandardArray(1, 2, 3, 4)
	for i := range a.Size() {
		v, _ := a.GetAt(i)
		test.GotWant(t, v, i+1)
	}
}

// Verifies getting at negative index
func TestStandardArray_GetAt_NegativeIndex(t *testing.T) {
	a := NewStandardArray[int]()
	v, err := a.GetAt(-1)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, v, 0)
}

// Verifies getting at invalid index
func TestStandardArray_GetAt_InvalidIndex(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	v, err := a.GetAt(3)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, v, 0)
}

// Verifies getting at the start of a multi-element array
func TestStandardArray_GetAt_Start(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	v, err := a.GetAt(0)
	test.GotWant(t, err, nil)
	test.GotWant(t, v, 1)
}

// Verifies getting at the end of a multi-element array
func TestStandardArray_GetAt_End(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	v, err := a.GetAt(2)
	test.GotWant(t, err, nil)
	test.GotWant(t, v, 3)
}

// Verifies getting at the middle of a multi-element array
func TestStandardArray_GetAt_Middle(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	v, err := a.GetAt(1)
	test.GotWant(t, err, nil)
	test.GotWant(t, v, 2)
}

// Verifies all elements are accessible in the correct order by index
func TestStandardArray_GetAt_Order(t *testing.T) {
	a := NewStandardArray(1, 2, 3, 4)

	for i := range a.Size() {
		v, err := a.GetAt(i)
		test.GotWant(t, err, nil)
		test.GotWant(t, v, i+1)
	}
}

// Verifies updating at negative index
func TestStandardArray_UpdateAt_NegativeIndex(t *testing.T) {
	a := NewStandardArray[int]()
	old, err := a.UpdateAt(-1, 0)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, old, 0)
}

// Verifies updating at invalid index
func TestStandardArray_UpdateAt_InvalidIndex(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	old, err := a.UpdateAt(3, 4)
	test.GotWantError(t, err, ErrorIndexOutOfRange)
	test.GotWant(t, old, 0)
}

// Verifies updating at the start of a multi-element array
func TestStandardArray_UpdateAt_Start(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	old, err := a.UpdateAt(0, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, old, 1)
	new, _ := a.GetAt(0)
	test.GotWant(t, new, 4)
}

// Verifies updating at the end of a multi-element array
func TestStandardArray_UpdateAt_End(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	old, err := a.UpdateAt(2, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, old, 3)
	new, _ := a.GetAt(2)
	test.GotWant(t, new, 4)
}

// Verifies updating the middle of a multi-element array
func TestStandardArray_UpdateAt_Middle(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	old, err := a.UpdateAt(1, 4)
	test.GotWant(t, err, nil)
	test.GotWant(t, old, 2)
	new, _ := a.GetAt(1)
	test.GotWant(t, new, 4)
}

// Verifies updating in order
func TestStandardArray_UpdateAt_Order(t *testing.T) {
	a := NewStandardArray(0, 1, 2)
	for i := range a.Size() {
		old, _ := a.UpdateAt(i, i+1)
		test.GotWant(t, old, i)
	}

	for i := range a.Size() {
		new, _ := a.GetAt(i)
		test.GotWant(t, new, i+1)
	}
}

// Verifies the empty state of an empty array
func TestStandardArray_IsEmpty_EmptyArray(t *testing.T) {
	a := NewStandardArray[int]()
	test.GotWant(t, a.IsEmpty(), true)
}

// Verifies the empty state of an non-empty array
func TestStandardArray_IsEmpty_NonEmptyArray(t *testing.T) {
	a := NewStandardArray(1)
	test.GotWant(t, a.IsEmpty(), false)
}

// Verifies the size of an empty array
func TestStandardArray_Size_EmptyArray(t *testing.T) {
	a := NewStandardArray[int]()
	test.GotWant(t, a.Size(), 0)
}

// Verifies the size of an non-empty array
func TestStandardArray_Size_NonEmptyArray(t *testing.T) {
	a := NewStandardArray(1, 2, 3)
	test.GotWant(t, a.Size(), 3)
}
