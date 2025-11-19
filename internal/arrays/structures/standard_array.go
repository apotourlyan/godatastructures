package structures

import "errors"

// Compile-time interface verifications
var _ Array[int] = &StandardArray[int]{}

// StandardArray implements a fixed-size array using a slice.
//
// Once created, the array size cannot be changed. All elements are
// accessible by zero-based index in O(1) time. The underlying storage
// is a slice, providing efficient random access.
type StandardArray[T any] struct {
	data []T
}

// NewStandardArray creates a fixed-size array initialized with the provided values.
// The array size equals the number of values provided. An empty variadic
// argument creates an empty array.
//
// The values are copied into the array, so modifications to the original
// slice do not affect the array.
//
// Example:
//
//	arr := NewStandardArray(1, 2, 3)  // Creates array of size 3
//	empty := NewStandardArray[int]()   // Creates empty array
//
// Time complexity: O(n) where n is the number of values
func NewStandardArray[T any](values ...T) *StandardArray[T] {
	data := make([]T, len(values))
	copy(data, values)
	return &StandardArray[T]{data}
}

// GetAt returns the element at the specified index.
// Valid indices are 0 to Size()-1.
// Returns ErrorIndexOutOfRange if index is invalid.
//
// Time complexity: O(1)
func (a *StandardArray[T]) GetAt(index int) (T, error) {
	if index < 0 || index >= len(a.data) {
		var zero T
		return zero, errors.New(ErrorIndexOutOfRange)
	}

	return a.data[index], nil
}

// UpdateAt updates the value at the specified index and returns the old value.
// Valid indices are 0 to Size()-1.
// Returns ErrorIndexOutOfRange if index is invalid.
//
// Time complexity: O(1)
func (a *StandardArray[T]) UpdateAt(index int, value T) (T, error) {
	if index < 0 || index >= len(a.data) {
		var zero T
		return zero, errors.New(ErrorIndexOutOfRange)
	}

	old := a.data[index]
	a.data[index] = value
	return old, nil
}

// IsEmpty returns true if the array contains no elements.
//
// Time complexity: O(1)
func (a *StandardArray[T]) IsEmpty() bool {
	return len(a.data) == 0
}

// Size returns the number of elements in the array.
//
// Time complexity: O(1)
func (a *StandardArray[T]) Size() int {
	return len(a.data)
}
