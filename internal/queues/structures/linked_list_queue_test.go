package structures

/*
Test Coverage
=============
Constructor (NewBasicLinkedList):
  ✓ Empty queue
  ✓ Single value
  ✓ Multiple values

Enqueue:
  ✓ Single value to empty queue
  ✓ Single value to non-empty queue
  ✓ Multiple values to empty queue
  ✓ Multiple values to non-empty queue

Dequeue:
  ✓ Single value from empty queue
  ✓ Single value from non-empty queue
  ✓ Multiple values from non-empty queue

Enqueue/Dequeue:
  ✓ FIFO order
  ✓ Reusable after emptying the queue

Peek:
  ✓ Empty queue
  ✓ Non-empty queue (single peek)
  ✓ Non-empty queue (multiple peeks)

IsEmpty/Size:
  ✓ Empty queue
  ✓ Non-empty queue
*/

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Verifies the creation of an empty queue
func TestLinkedListQueue_NewLinkedListQueue_Empty(t *testing.T) {
	q := NewLinkedListQueue[int]()
	test.GotWant(t, q.Size(), 0)
	test.GotWant(t, q.IsEmpty(), true)
}

// Verifies the creation of one-element queue
func TestLinkedListQueue_NewLinkedListQueue_OneValue(t *testing.T) {
	q := NewLinkedListQueue(1)
	test.GotWant(t, q.Size(), 1)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the creation of multi-element queue
func TestLinkedListQueue_NewLinkedListQueue_ManyValues(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)
	test.GotWant(t, q.Size(), 3)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the enqueuing of an element in an empty queue
func TestLinkedListQueue_Enqueue_OneElement_EmptyQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	q.Enqueue(1)
	p, _ := q.Peek()
	test.GotWant(t, p, 1)
	test.GotWant(t, q.Size(), 1)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the enqueuing of multiple elements in an empty queue
func TestLinkedListQueue_Enqueue_ManyElements_EmptyQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	p, _ := q.Peek()
	test.GotWant(t, p, 1)
	test.GotWant(t, q.Size(), 3)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the enqueuing of an element in a non-empty queue
func TestLinkedListQueue_Enqueue_OneElement_NonEmptyQueue(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)
	q.Enqueue(4)
	p, _ := q.Peek()
	test.GotWant(t, p, 1)
	test.GotWant(t, q.Size(), 4)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the enqueuing of multiple elements in a non-empty queue
func TestLinkedListQueue_Enqueue_ManyElements_NonEmptyQueue(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)
	q.Enqueue(4)
	q.Enqueue(5)
	p, _ := q.Peek()
	test.GotWant(t, p, 1)
	test.GotWant(t, q.Size(), 5)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the dequiuing of an element from an empty queue
func TestLinkedListQueue_Dequeue_OneElement_EmptyQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	d, err := q.Dequeue()
	test.GotWantError(t, err, ErrorEmptyQueue)
	test.GotWant(t, d, 0)
	test.GotWant(t, q.Size(), 0)
	test.GotWant(t, q.IsEmpty(), true)
}

// Verifies the dequiuing of an element from a non-empty queue
func TestLinkedListQueue_Dequeue_OneElement_NonEmptyQueue(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)
	d, err := q.Dequeue()
	test.GotWant(t, err, nil)
	test.GotWant(t, d, 1)
	p, _ := q.Peek()
	test.GotWant(t, p, 2)
	test.GotWant(t, q.Size(), 2)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the dequiuing of multiple elements in a non-empty queue
func TestLinkedListQueue_Dequeue_ManyElements_NonEmptyQueue(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)
	q.Dequeue()
	q.Dequeue()
	d, err := q.Dequeue()
	test.GotWant(t, err, nil)
	test.GotWant(t, d, 3)
	test.GotWant(t, q.Size(), 0)
	test.GotWant(t, q.IsEmpty(), true)
}

// Verifies First-In-First-Out element order
func TestLinkedListQueue_EnqueueDequeue_Order(t *testing.T) {
	q := NewLinkedListQueue[int]()

	for i := range 5 {
		q.Enqueue(i + 1)
		p, _ := q.Peek()
		test.GotWant(t, p, 1)
	}

	for i := range 5 {
		p, _ := q.Peek()
		test.GotWant(t, p, i+1)
		d, _ := q.Dequeue()
		test.GotWant(t, d, i+1)
	}
}

// Verifies the queue is reusable
func TestLinkedListQueue_EnqueueDequeue_Reusability(t *testing.T) {
	q := NewLinkedListQueue[int]()
	q.Enqueue(1)
	q.Dequeue()
	test.GotWant(t, q.IsEmpty(), true)
	q.Enqueue(2)
	p, _ := q.Peek()
	test.GotWant(t, p, 2)
}

// Verifies peeking into an empty queue
func TestLinkedListQueue_Peek_EmptyQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	p, err := q.Peek()
	test.GotWantError(t, err, ErrorEmptyQueue)
	test.GotWant(t, p, 0)
	test.GotWant(t, q.Size(), 0)
	test.GotWant(t, q.IsEmpty(), true)
}

// Verifies peeking into an non-empty queue
func TestLinkedListQueue_Peek_NonEmptyQueue(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)
	p, err := q.Peek()
	test.GotWant(t, err, nil)
	test.GotWant(t, p, 1)
	test.GotWant(t, q.Size(), 3)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies peeking multiple times into an non-empty queue
func TestLinkedListQueue_Peek_Many(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)

	for range 3 {
		p, err := q.Peek()
		test.GotWant(t, err, nil)
		test.GotWant(t, p, 1)
	}
}

// Verifies the empty state of an empty queue
func TestLinkedListQueue_IsEmpty_EmptyQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	test.GotWant(t, q.IsEmpty(), true)
}

// Verifies the empty state of an non-empty queue
func TestLinkedListQueue_IsEmpty_NonEmptyQueue(t *testing.T) {
	q := NewLinkedListQueue(1)
	test.GotWant(t, q.IsEmpty(), false)
}

// Verifies the size of an empty queue
func TestLinkedListQueue_Size_EmptyQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	test.GotWant(t, q.Size(), 0)
}

// Verifies the size of an non-empty queue
func TestLinkedListQueue_Size_NonEmptyQueue(t *testing.T) {
	q := NewLinkedListQueue(1, 2, 3)
	test.GotWant(t, q.Size(), 3)
}
