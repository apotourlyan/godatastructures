// Package structures provides generic list data structures and their implementations.
package structures

const ErrorEmptyList = "list is empty"
const ErrorIndexOutOfRange = "index is out of the range of possible values"

// Provides fundamental list operations without requiring element comparison.
type BasicList[T any] interface {
	// Prepends a value to the start of the list.
	// Time complexity depends on implementation.
	AddFirst(value T)

	// Appends a value to the end of the list.
	// Time complexity depends on implementation.
	AddLast(value T)

	// Removes a value from the start of the list.
	// Returns false if the list is empty.
	// Time complexity depends on implementation.
	RemoveFirst() bool

	// Removes a value from the end of the list.
	// Returns false if the list is empty.
	// Time complexity depends on implementation.
	RemoveLast() bool

	// Returns the first element in the list.
	// Returns ErrorEmptyList if the list is empty.
	// Time complexity depends on implementation.
	First() (T, error)

	// Returns the last element in the list.
	// Returns ErrorEmptyList if the list is empty.
	// Time complexity depends on implementation.
	Last() (T, error)

	// Returns true if the list contains no elements.
	// Time complexity: O(1)
	IsEmpty() bool

	// Returns the number of elements in the list.
	// Time complexity: O(1)
	Size() int
}

// Provides position-based access and mutation list operations.
type IndexedList[T any] interface {
	// Inserts a value at the specified index.
	// Valid indices are 0 to Size() inclusive (append at end).
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity depends on implementation.
	InsertAt(index int, value T) error

	// Updates a value at the specified index.
	// Valid indices are 0 to Size()-1.
	// Returns the old value at the specified index.
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity depends on implementation.
	UpdateAt(index int, value T) (T, error)

	// Removes the element at the specified index.
	// Valid indices are 0 to Size()-1.
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity depends on implementation.
	RemoveAt(index int) error

	// Returns the element at the specified index.
	// Valid indices are 0 to Size()-1.
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity depends on implementation.
	GetAt(index int) (T, error)
}

// Provides value-based search and manipulation list operations.
type SearchableList[T comparable] interface {
	// Returns the index of the first occurrence of the specified value.
	// Returns -1 if the value is not found.
	// Time complexity: O(n) where n is the number of elements.
	IndexOf(value T) int

	// Returns true if the list contains the specified value.
	// Time complexity: O(n) where n is the number of elements.
	Contains(value T) bool

	// Removes the first occurrence of the specified value.
	// Returns true if the value was found and removed, false otherwise.
	// Time complexity: O(n) where n is the number of elements.
	Remove(value T) bool

	// Updates the first occurrence of the specified value.
	// Returns true if the value was found and updated, false otherwise.
	// Time complexity: O(n) where n is the number of elements.
	Update(oldValue T, newValue T) bool
}

// Represents a complete generic list collection with all operations.
//
// This interface combines BasicList, IndexedList, and SearchableList,
// providing the full suite of list operations. Elements must be comparable
// to support search operations.
type List[T comparable] interface {
	BasicList[T]
	IndexedList[T]
	SearchableList[T]
}
