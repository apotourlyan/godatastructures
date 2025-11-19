package structures

import (
	"errors"

	"github.com/apotourlyan/godatastructures/internal/slices/algorithms"
)

// Compile-time interface verifications
var _ Stack[int] = &SliceStack[int]{}

// SliceStack implements a LIFO stack using a dynamic slice with optional
// memory optimization.
//
// Optimization Strategy:
//
// ReallocateOnPop: Shrinks capacity when waste > threshold after Pop operations
//   - Best for: stacks that grow large then permanently shrink
//   - Benefit: Reclaims ~97-99% of wasted memory after shrinkage
//   - Tradeoff: Reallocation overhead (one-time O(n) cost)
//
// Default configuration enables reallocation with conservative thresholds,
// suitable for most workloads. Disable for pure growth patterns or when
// memory overhead is acceptable.
type SliceStack[T any] struct {
	curr   int              // Exclusive index of back element
	data   []T              // Underlying slice storage
	config SliceStackConfig // Optimization configuration
}

// NewSliceStack creates a stack with default optimizations enabled.
// Suitable for most workloads including growth-shrink cycles and
// temporary large allocations.
//
// For specific workloads, use NewSliceStackWithConfig:
//   - Pure growth: disable ReallocateOnPop
//   - Memory-constrained: enable with aggressive thresholds (90-99% waste)
//   - CPU-constrained: disable or use conservative thresholds (60-70% waste)
//   - Unknown/mixed: use default (reallocation enabled, 75% threshold)
func NewSliceStack[T any](values ...T) *SliceStack[T] {
	c := SliceStackConfig{
		ReallocateOnPop:        true,
		MinOptimizationLength:  100,
		ReallocateWastePercent: 75,
		ReallocateWasteBuffer:  80,
	}

	return NewSliceStackWithConfig(c, values...)
}

// NewSliceStackWithConfig creates a stack with custom optimization settings.
// See SliceStackConfig for configuration options and tuning guidance.
//
// Example:
//
//	config := SliceStackConfig{
//	    ReallocateOnPop:        true,
//	    MinOptimizationLength:  500,
//	    ReallocateWastePercent: 80,
//	    ReallocateWasteBuffer:  70,
//	}
//	s := NewSliceStackWithConfig(config, 1, 2, 3)
func NewSliceStackWithConfig[T any](config SliceStackConfig, values ...T) *SliceStack[T] {
	s := &SliceStack[T]{
		data: make([]T, 0, len(values)),
	}

	s.data = append(s.data, values...)
	s.curr = len(values)
	s.config = config
	return s
}

// Push adds an element to the top of the stack.
//
// Time complexity: O(1) amortized
func (s *SliceStack[T]) Push(value T) {
	if s.curr == len(s.data) {
		s.data = append(s.data, value)
	} else {
		s.data[s.curr] = value
	}

	s.curr++
}

// Pop removes and returns the element at the top of the stack.
// Returns an error if the stack is empty.
// If ReallocateOnPop is enabled and waste exceeds the threshold,
// reallocation occurs after popping to free memory.
//
// Time complexity: O(1) amortized, O(n) when reallocation triggers
func (s *SliceStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New(ErrorEmptyStack)
	}

	v := s.data[s.curr-1]
	s.curr--

	// Reset when empty
	if s.curr == 0 {
		s.data = s.data[:0]
	} else if s.config.ReallocateOnPop {
		s.data, _, s.curr = algorithms.Reallocate(
			s.data, algorithms.SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      s.curr,
				MinSize:      s.config.MinOptimizationLength,
				WastePercent: s.config.ReallocateWastePercent,
				WasteBuffer:  s.config.ReallocateWasteBuffer,
			})
	}

	return v, nil
}

// Peek returns the element at the top of the stack without removing it.
// Returns an error if the stack is empty.
//
// Time complexity: O(1)
func (s *SliceStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New(ErrorEmptyStack)
	}

	return s.data[s.curr-1], nil
}

// IsEmpty returns true if the stack contains no elements.
//
// Time complexity: O(1)
func (s *SliceStack[T]) IsEmpty() bool {
	return s.curr == 0
}

// Size returns the number of elements currently in the stack.
//
// Time complexity: O(1)
func (s *SliceStack[T]) Size() int {
	return s.curr
}
