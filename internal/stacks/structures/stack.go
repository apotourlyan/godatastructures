package structures

const ErrorEmptyStack = "stack is empty"

// Stack defines the interface for a LIFO (Last-In-First-Out) data structure.
// Elements are added to the top and removed from the top, maintaining reverse insertion order.
//
// All implementations guarantee:
//   - Push operations add elements to the top
//   - Pop operations remove elements from the top
//   - Peek operations observe the top without removal
//   - Size and IsEmpty operations reflect current state
//
// Thread safety is implementation-dependent. Check specific implementation
// documentation for concurrency guarantees.
type Stack[T any] interface {
	// Push adds an element to the top of the stack.
	Push(value T)

	// Pop removes and returns the element at the top of the stack.
	// Returns an error if the stack is empty.
	Pop() (T, error)

	// Peek returns the element at the top of the stack without removing it.
	// Returns an error if the stack is empty.
	Peek() (T, error)

	// IsEmpty returns true if the stack contains no elements.
	IsEmpty() bool

	// Size returns the number of elements currently in the stack.
	Size() int
}
