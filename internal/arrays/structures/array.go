package structures

const ErrorIndexOutOfRange = "index is out of the range of possible values"

// Array defines the interface for a fixed-size indexed collection.
// Elements are accessed and updated by zero-based index in O(1) time.
//
// Unlike dynamic collections, arrays have fixed size after creation.
// All implementations guarantee:
//   - GetAt operations retrieve elements by index
//   - UpdateAt operations modify elements by index and return old values
//   - Size and IsEmpty operations reflect current state
//   - Index bounds are validated (0 to Size()-1)
//
// Thread safety is implementation-dependent. Check specific implementation
// documentation for concurrency guarantees.
type Array[T any] interface {
	// GetAt returns the element at the specified index.
	// Valid indices are 0 to Size()-1.
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity: O(1)
	GetAt(index int) (T, error)

	// UpdateAt updates a value at the specified index.
	// Valid indices are 0 to Size()-1.
	// Returns the old value at the specified index.
	// Returns ErrorIndexOutOfRange if index is invalid.
	// Time complexity: O(1)
	UpdateAt(index int, value T) (T, error)

	// IsEmpty returns true if the array contains no elements.
	// Time complexity: O(1)
	IsEmpty() bool

	// Size returns the number of elements in the array.
	// Time complexity: O(1)
	Size() int
}
