package structures

import "errors"

// SliceQueue implements a FIFO queue using a dynamic slice with configurable
// memory optimizations. It supports two optimization strategies:
//
// 1. CompactOnEnqueue: Shifts elements to front when waste > threshold
//   - Best for: balanced ops, oscillating size, long-running queues
//   - Benefit: 2-3x faster, a lot less memory vs unoptimized
//   - Tradeoff: Copy overhead on compaction
//
// 2. ReallocateOnDequeue: Shrinks capacity when waste > threshold
//   - Best for: permanent shrinkage, memory-constrained environments
//   - Benefit: ~97-99% memory freed after shrinkage
//   - Tradeoff: Reallocation overhead
//
// Default configuration enables both optimizations for balanced performance.
// See benchmarks in slice_queue_bench_test.go for detailed comparisons.
type SliceQueue[T any] struct {
	curr   int              // Index of front element
	data   []T              // Underlying slice storage
	config SliceQueueConfig // Optimization configuration
}

// NewSliceQueue creates a queue with default optimizations enabled.
// Suitable for most workloads including balanced operations, oscillating
// sizes, and mixed growth/shrinkage patterns.
//
// For specific workloads, use NewSliceQueueWithConfig:
//   - Pure growth: disable both optimizations
//   - Balanced/oscillating: enable CompactOnEnqueue only
//   - Permanent shrinkage: enable ReallocateOnDequeue
//   - Unknown/mixed: use default (both enabled)
func NewSliceQueue[T any](values ...T) *SliceQueue[T] {
	config := SliceQueueConfig{
		CompactOnEnqueue:       true,
		ReallocateOnDequeue:    true,
		MinOptimizationLength:  100,
		CompactWastePercent:    50,
		ReallocateWastePercent: 75,
	}

	return NewSliceQueueWithConfig(config, values...)
}

// NewSliceQueueWithConfig creates a queue with custom optimization settings.
// See SliceQueueConfig for configuration options and tuning guidance.
//
// Example:
//
//	config := SliceQueueConfig{
//	    CompactOnEnqueue:      true,
//	    ReallocateOnDequeue:   false,
//	    MinOptimizationLength: 1000,
//	    CompactWastePercent:   60,
//	}
//	q := NewSliceQueueWithConfig(config, 1, 2, 3)
func NewSliceQueueWithConfig[T any](config SliceQueueConfig, values ...T) *SliceQueue[T] {
	q := &SliceQueue[T]{
		data: make([]T, 0, len(values)),
	}

	q.data = append(q.data, values...)
	q.config = config
	return q
}

// Enqueue adds an element to the back of the queue.
// If CompactOnEnqueue is enabled and waste exceeds the threshold,
// compaction occurs before enqueuing to reuse capacity.
//
// Time complexity: O(1) amortized, O(n) when compaction triggers
func (q *SliceQueue[T]) Enqueue(value T) {
	// Resize before enqueuing when waste is significant (> 'CompactWastePercent')
	optimize := q.config.CompactOnEnqueue &&
		q.curr >= q.config.MinOptimizationLength &&
		100.0*q.Size() < q.config.CompactWastePercent*len(q.data)

	if optimize {
		copy(q.data, q.data[q.curr:])
		q.data = q.data[:len(q.data)-q.curr]
		q.curr = 0
	}

	q.data = append(q.data, value)
}

// Dequeue removes and returns the element at the front of the queue.
// Returns an error if the queue is empty.
// If ReallocateOnDequeue is enabled and waste exceeds the threshold,
// reallocation occurs after dequeuing to free memory.
//
// Time complexity: O(1) amortized, O(n) when reallocation triggers
func (q *SliceQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New(ErrorEmptyQueue)
	}

	v := q.data[q.curr]
	q.curr++

	// Reallocate after dequeue when waste is significant (> 'ReallocateWastePercent')
	optimize := q.config.ReallocateOnDequeue &&
		q.curr >= q.config.MinOptimizationLength &&
		100.0*q.Size() < (100-q.config.ReallocateWastePercent)*cap(q.data)

	if optimize {
		data := q.data[q.curr:]
		q.data = make([]T, 0, max(len(data)*2, 10))
		q.data = append(q.data, data...)
		q.curr = 0
	}

	return v, nil
}

// Peek returns the element at the front of the queue without removing it.
// Returns an error if the queue is empty.
//
// Time complexity: O(1)
func (q *SliceQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New(ErrorEmptyQueue)
	}

	return q.data[q.curr], nil
}

// IsEmpty returns true if the queue contains no elements.
//
// Time complexity: O(1)
func (q *SliceQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Size returns the number of elements currently in the queue.
//
// Time complexity: O(1)
func (q *SliceQueue[T]) Size() int {
	return len(q.data) - q.curr
}
