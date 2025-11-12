package structures

import (
	"errors"

	lists "github.com/apotourlyan/godatastructures/internal/lists/structures"
)

// Compile-time interface verifications
var _ Queue[int] = &LinkedListQueue[int]{}

// LinkedListQueue is a FIFO queue backed by a singly-linked list.
//
// This implementation uses a BasicLinkedList as its underlying storage,
// providing true O(1) enqueue and dequeue operations without memory
// reallocation or compaction overhead.
type LinkedListQueue[T any] struct {
	data lists.BasicList[T] // Underlying basic list storage
}

// Creates a new LinkedListQueue with optional initial values.
//
// Values are enqueued in the order provided. If no values are given,
// an empty queue is created.
//
// Time complexity: O(n) where n is the number of initial values.
//
// Example:
//
//	empty := NewLinkedListQueue[int]()
//	withValues := NewLinkedListQueue(1, 2, 3)
func NewLinkedListQueue[T any](values ...T) *LinkedListQueue[T] {
	data := lists.NewBasicLinkedList(values...)
	return &LinkedListQueue[T]{data}
}

// Adds a value to the back of the queue.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	q := NewLinkedListQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	q.Enqueue(3)  // Queue is now [1, 2, 3]
func (q *LinkedListQueue[T]) Enqueue(value T) {
	q.data.AddLast(value)
}

// Removes and returns the value from the front of the queue.
//
// Returns ErrorEmptyQueue if the queue is empty.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	q := NewLinkedListQueue(1, 2, 3)
//	value, _ := q.Dequeue()  // Returns 1, queue is now [2, 3]
//	value, _ = q.Dequeue()   // Returns 2, queue is now [3]
func (q *LinkedListQueue[T]) Dequeue() (T, error) {
	f, err := q.data.First()
	if err != nil {
		var zero T
		return zero, errors.New(ErrorEmptyQueue)
	}

	q.data.RemoveFirst()
	return f, nil
}

// Returns the value at the front of the queue without removing it.
//
// Returns ErrorEmptyQueue if the queue is empty.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	q := NewLinkedListQueue(1, 2, 3)
//	value, _ := q.Peek()  // Returns 1, queue unchanged
//	value, _ = q.Peek()   // Returns 1 again
func (q *LinkedListQueue[T]) Peek() (T, error) {
	f, err := q.data.First()
	if err != nil {
		var zero T
		return zero, errors.New(ErrorEmptyQueue)
	}

	return f, nil
}

// Returns true if the queue contains no elements.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	q := NewLinkedListQueue[int]()
//	q.IsEmpty()  // Returns true
//	q.Enqueue(1)
//	q.IsEmpty()  // Returns false
func (q *LinkedListQueue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

// Returns the number of elements in the queue.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	q := NewLinkedListQueue(1, 2, 3)
//	q.Size()  // Returns 3
func (q *LinkedListQueue[T]) Size() int {
	return q.data.Size()
}
