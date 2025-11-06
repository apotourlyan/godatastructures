package structures

const ErrorEmptyQueue = "queue is empty"

// Queue defines the interface for a FIFO (First-In-First-Out) data structure.
// Elements are added to the back and removed from the front, maintaining insertion order.
//
// All Queue implementations guarantee:
//   - Enqueue operations add elements to the back
//   - Dequeue operations remove elements from the front
//   - Peek operations observe the front without removal
//   - Size and IsEmpty operations reflect current state
//
// Thread safety is implementation-dependent. Check specific implementation
// documentation for concurrency guarantees.
type Queue[T any] interface {
	// Enqueue adds an element to the back of the queue.
	Enqueue(value T)

	// Dequeue removes and returns the element at the front of the queue.
	// Returns an error if the queue is empty.
	Dequeue() (T, error)

	// Peek returns the element at the front of the queue without removing it.
	// Returns an error if the queue is empty.
	Peek() (T, error)

	// IsEmpty returns true if the queue contains no elements.
	IsEmpty() bool

	// Size returns the number of elements currently in the queue.
	Size() int
}
