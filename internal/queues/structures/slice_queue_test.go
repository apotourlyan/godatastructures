package structures

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// TestSliceQueue_Empty verifies behavior on an empty queue.
// Ensures Peek and Dequeue return appropriate errors.
func TestSliceQueue_Empty(t *testing.T) {
	q := NewSliceQueueWithConfig[int](
		SliceQueueConfig{
			CompactOnEnqueue:    false,
			ReallocateOnDequeue: false,
		})

	p, pErr := q.Peek()
	test.GotWant(t, p, 0)
	test.GotWantError(t, pErr, ErrorEmptyQueue)
	test.GotWant(t, q.Size(), 0)
	test.GotWant(t, q.IsEmpty(), true)

	d, dErr := q.Dequeue()
	test.GotWant(t, d, 0)
	test.GotWantError(t, dErr, ErrorEmptyQueue)
}

// TestSliceQueue_InitialValues verifies constructor with initial values.
func TestSliceQueue_InitialValues(t *testing.T) {
	q := NewSliceQueueWithConfig(
		SliceQueueConfig{
			CompactOnEnqueue:    false,
			ReallocateOnDequeue: false,
		}, 1, 2, 3)

	p, pErr := q.Peek()
	test.GotWant(t, p, 1)
	test.GotWantError(t, pErr, "")
	test.GotWant(t, q.Size(), 3)
	test.GotWant(t, q.IsEmpty(), false)

	d, dErr := q.Dequeue()
	test.GotWant(t, d, 1)
	test.GotWantError(t, dErr, "")
}

// TestSliceQueue_AlternativeType verifies queue works with non-int types.
func TestSliceQueue_AlternativeType(t *testing.T) {
	q := NewSliceQueueWithConfig(
		SliceQueueConfig{
			CompactOnEnqueue:    false,
			ReallocateOnDequeue: false,
		}, "hello", "world")

	d, _ := q.Dequeue()
	test.GotWant(t, d, "hello")
	test.GotWant(t, q.Size(), 1)
}

// TestSliceQueue_FirstInFirstOutOrder verifies FIFO ordering is maintained.
// Elements should be dequeued in the same order they were enqueued.
func TestSliceQueue_FirstInFirstOutOrder(t *testing.T) {
	q := NewSliceQueueWithConfig[int](
		SliceQueueConfig{
			CompactOnEnqueue:    false,
			ReallocateOnDequeue: false,
		})

	size := 0
	for i := range 3 {
		size++
		q.Enqueue(i)

		p, pErr := q.Peek()
		test.GotWant(t, p, 0)
		test.GotWantError(t, pErr, "")
		test.GotWant(t, q.Size(), size)
		test.GotWant(t, q.IsEmpty(), false)
	}

	for i := range 3 {
		p, pErr := q.Peek()
		test.GotWant(t, p, i)
		test.GotWantError(t, pErr, "")
		test.GotWant(t, q.Size(), size)
		test.GotWant(t, q.IsEmpty(), size == 0)

		d, dErr := q.Dequeue()
		test.GotWant(t, d, i)
		test.GotWantError(t, dErr, "")
		size--
	}

	p, pErr := q.Peek()
	test.GotWant(t, p, 0)
	test.GotWantError(t, pErr, ErrorEmptyQueue)
	test.GotWant(t, q.Size(), 0)
	test.GotWant(t, q.IsEmpty(), true)
}

// TestSliceQueue_PeekDoesNotModify verifies Peek is non-destructive.
// Multiple peeks should return the same value without changing queue state.
func TestSliceQueue_PeekDoesNotModify(t *testing.T) {
	q := NewSliceQueueWithConfig(
		SliceQueueConfig{
			CompactOnEnqueue:    false,
			ReallocateOnDequeue: false,
		}, 1, 2, 3)

	for range 5 {
		p, pErr := q.Peek()
		test.GotWant(t, p, 1)
		test.GotWantError(t, pErr, "")
		test.GotWant(t, q.Size(), 3)
		test.GotWant(t, q.IsEmpty(), false)
	}
}

// TestSliceQueue_ReusableAfterEmpty verifies queue can be reused after becoming empty.
func TestSliceQueue_ReusableAfterEmpty(t *testing.T) {
	q := NewSliceQueueWithConfig[int](
		SliceQueueConfig{
			CompactOnEnqueue:    false,
			ReallocateOnDequeue: false,
		})

	// Fill and empty
	q.Enqueue(1)
	q.Dequeue()

	test.GotWant(t, q.IsEmpty(), true)

	q.Enqueue(2)
	p, _ := q.Peek()
	test.GotWant(t, p, 2)
}

// TestSliceQueue_LargeScale verifies correctness with large number of operations.
func TestSliceQueue_LargeScale(t *testing.T) {
	q := NewSliceQueueWithConfig[int](
		SliceQueueConfig{
			CompactOnEnqueue:    false,
			ReallocateOnDequeue: false,
		})

	for i := range 10000 {
		q.Enqueue(i)
	}

	test.GotWant(t, q.Size(), 10000)

	for i := range 10000 {
		d, _ := q.Dequeue()
		test.GotWant(t, d, i)
	}

	test.GotWant(t, q.IsEmpty(), true)
}

// TestSliceQueue_Compaction verifies compaction optimization triggers correctly.
// When waste exceeds threshold, compaction should reset curr to 0 and reuse capacity.
func TestSliceQueue_Compaction(t *testing.T) {
	q := NewSliceQueueWithConfig[int](SliceQueueConfig{
		CompactOnEnqueue:      true,
		ReallocateOnDequeue:   false,
		MinOptimizationLength: 10,
		CompactWastePercent:   50,
	})

	for i := range 100 {
		q.Enqueue(i)
	}

	// Create large waste
	for range 60 {
		q.Dequeue()
	}

	test.GotWant(t, q.curr > 0, true)
	// This enqueue should trigger compaction
	q.Enqueue(999)
	test.GotWant(t, q.curr, 0)
	test.GotWant(t, q.Size(), 41) // 40 remaining + 1 new
}

// TestSliceQueue_Reallocation verifies reallocation optimization triggers correctly.
// When waste exceeds threshold, reallocation should shrink capacity.
func TestSliceQueue_Reallocation(t *testing.T) {
	q := NewSliceQueueWithConfig[int](SliceQueueConfig{
		CompactOnEnqueue:       false,
		ReallocateOnDequeue:    true,
		MinOptimizationLength:  10,
		ReallocateWastePercent: 75,
	})

	// Create large capacity
	for i := range 1000 {
		q.Enqueue(i)
	}

	capBefore := cap(q.data)
	// Dequeue most (leaving < 25%, should trigger reallocation)
	for range 850 {
		q.Dequeue()
	}

	capAfter := cap(q.data)
	// Verify capacity has shrunk
	test.GotWant(t, capAfter < capBefore, true)
	test.GotWant(t, q.Size(), 150)
}
