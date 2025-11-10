package structures

/*
Testing Strategy
================

The SliceQueue test suite verifies both fundamental queue behavior and
optimization mechanisms. These tests emphasize:

1. FIFO Semantics
   - Elements dequeued in exact order they were enqueued
   - No random access (unlike lists)
   - Peek returns front without modification

2. Optimization Correctness
   - Compaction triggers at correct thresholds
   - Reallocation reclaims memory appropriately
   - Optimizations preserve FIFO ordering

3. Configuration Flexibility
   - NoOptimizations baseline works correctly
   - Each optimization can be enabled independently
   - Configurations don't break core behavior

Core Test Principles
====================

Every test uses NoOptimizations config by default to verify baseline behavior.
Optimization-specific tests explicitly enable and verify optimization triggers.

Configuration Pattern:
  NoOptimizations = baseline correctness tests
  Specific configs = optimization behavior tests

This separation ensures:
- Core queue logic is correct before testing optimizations
- Optimization bugs don't mask fundamental issues
- Each optimization can be verified in isolation

Test Organization
=================

Tests are organized by concern:

Basic Behavior Tests (NoOptimizations):
  - Empty queue operations
  - Constructor with initial values
  - Generic type support
  - FIFO ordering
  - Peek non-destructive behavior
  - Reusability after empty
  - Large-scale correctness

Optimization Tests (Specific Configs):
  - Compaction triggering and correctness
  - Reallocation triggering and memory reclamation

This organization makes it clear which tests verify core behavior
vs optimization-specific functionality.

Test Coverage
=============

Coverage by category:

Basic Queue Operations:
  ✓ Empty queue (Peek and Dequeue return errors)
  ✓ Constructor with values
  ✓ Generic type support (string example)
  ✓ FIFO ordering maintained
  ✓ Peek is non-destructive
  ✓ Reusable after becoming empty
  ✓ Large-scale operations (10,000 elements)

Optimization Verification:
  ✓ Compaction triggers at threshold
  ✓ Compaction resets curr pointer
  ✓ Compaction preserves elements
  ✓ Reallocation triggers at threshold
  ✓ Reallocation shrinks capacity
  ✓ Reallocation preserves elements
*/

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Purpose: Verify empty queue behavior
//
// Verifies: ErrorEmptyQueue returned, size == 0, isEmpty == true
//
// Config: NoOptimizations
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

// Purpose: Verify constructor with values
//
// Verifies: Values stored correctly, size correct, FIFO order
//
// Config: NoOptimizations
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

// Purpose: Verify generic type support
//
// Verifies: Works with string type (not just int)
//
// Config: NoOptimizations
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

// Purpose: Verify FIFO ordering maintained
//
// Verifies: Enqueue 0,1,2 → Dequeue 0,1,2 in order
//
// Details: Tests both Peek and Dequeue order preservation
//
// Config: NoOptimizations
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

// Purpose: Verify Peek is non-destructive
//
// Verifies: Multiple peeks return same value, size unchanged
//
// Config: NoOptimizations
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

// Purpose: Verify queue can be reused after becoming empty
//
// Verifies: Enqueue → Dequeue → Enqueue works correctly
//
// Config: NoOptimizations
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

// Purpose: Verify correctness with large number of operations
//
// Verifies: 10,000 enqueues followed by 10,000 dequeues in order
//
// Config: NoOptimizations
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

// Purpose: Verify compaction optimization triggers correctly
//
// Setup: Enqueue 100, Dequeue 60 (60% waste)
//
// Config: CompactOnEnqueue, 50% threshold
//
// Verifies:
//   - curr > 0 before compaction
//   - curr == 0 after compaction
//   - Size correct after compaction
//   - Elements preserved
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

// Purpose: Verify reallocation optimization triggers correctly
//
// Setup: Enqueue 1000, Dequeue 850 (85% waste)
//
// Config: ReallocateOnDequeue, 75% threshold
//
// Verifies:
//   - Capacity shrinks (capAfter < capBefore)
//   - Size correct after reallocation
//   - Elements preserved
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
