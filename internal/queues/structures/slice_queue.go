package structures

import "errors"

type SliceQueue[T any] struct {
	curr   int
	data   []T
	config SliceQueueConfig
}

type SliceQueueConfig struct {
	CompactOnEnqueue       bool
	ReallocateOnDequeue    bool
	MinOptimizationLength  int
	CompactWastePercent    int
	ReallocateWastePercent int
}

func NewSliceQueue[T any](values ...T) *SliceQueue[T] {
	config := SliceQueueConfig{
		CompactOnEnqueue:       true,
		ReallocateOnDequeue:    false,
		MinOptimizationLength:  100,
		CompactWastePercent:    50,
		ReallocateWastePercent: 75,
	}

	return NewSliceQueueWithConfig(config, values...)
}

func NewSliceQueueWithConfig[T any](config SliceQueueConfig, values ...T) *SliceQueue[T] {
	q := &SliceQueue[T]{
		data: make([]T, 0, len(values)),
	}

	q.data = append(q.data, values...)
	q.config = config
	return q
}

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

func (q *SliceQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New(ErrorEmptyQueue)
	}

	return q.data[q.curr], nil
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *SliceQueue[T]) Size() int {
	return len(q.data) - q.curr
}
