// Package lists provides generic list data structures and their implementations.
//
// The package defines a List interface and provides a LinkedList implementation
// with head and tail pointers for efficient operations at both ends.
package lists

const ErrorEmptyList = "list is empty"
const ErrorIndexOutOfRange = "index is out of the range of possible values"

// List represents a generic ordered collection of elements.
//
// All implementations maintain insertion order and provide O(1) size operations.
// Elements must be comparable to support search operations (IndexOf, Contains, Remove).
//
// The interface is designed for general list operations including:
//   - Adding and removing elements
//   - Indexed access and mutation
//   - Searching for elements
//   - Accessing boundary elements (first/last)
type List[T comparable] interface {
	// Add appends a value to the end of the list.
	// Time complexity depends on implementation.
	Add(value T)

	// Remove removes the first occurrence of the specified value.
	// Returns true if the value was found and removed, false otherwise.
	// Time complexity: O(n) where n is the number of elements.
	Remove(value T) bool

	// InsertAt inserts a value at the specified index.
	// Valid indices are 0 to Size() inclusive (append at end).
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity depends on implementation.
	InsertAt(index int, value T) error

	// RemoveAt removes the element at the specified index.
	// Valid indices are 0 to Size()-1.
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity depends on implementation.
	RemoveAt(index int) error

	// GetAt returns the element at the specified index.
	// Valid indices are 0 to Size()-1.
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity depends on implementation.
	GetAt(index int) (T, error)

	// IndexOf returns the index of the first occurrence of the specified value.
	// Returns -1 if the value is not found.
	// Time complexity: O(n) where n is the number of elements.
	IndexOf(value T) int

	// Contains returns true if the list contains the specified value.
	// Time complexity: O(n) where n is the number of elements.
	Contains(value T) bool

	// First returns the first element in the list.
	// Returns ErrorEmptyList if the list is empty.
	// Time complexity depends on implementation.
	First() (T, error)

	// Last returns the last element in the list.
	// Returns ErrorEmptyList if the list is empty.
	// Time complexity depends on implementation.
	Last() (T, error)

	// IsEmpty returns true if the list contains no elements.
	// Time complexity: O(1)
	IsEmpty() bool

	// Size returns the number of elements in the list.
	// Time complexity: O(1)
	Size() int
}
