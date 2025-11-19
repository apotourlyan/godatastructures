package algorithms

import "github.com/apotourlyan/godatastructures/internal/utilities/panics"

// Controls when and how to reallocate a slice-based data structure.
type SliceReallocationParams struct {
	UsedStart    int // Index of first used element
	UsedEnd      int // Exclusive index of last used element
	MinSize      int // Minimum used size to trigger reallocation (0 means always reallocate if waste threshold is met)
	WastePercent int // Reallocate if waste >= this percent (0-100)
	WasteBuffer  int // Target waste as percent of threshold (0-99, e.g. 80 means target 80% of threshold)
}

// Validates reallocation parameters against slice length.
//
// Panics if parameters are invalid:
//   - UsedStart outside [0, UsedEnd)
//   - UsedEnd outside [0, length)
//   - MinSize < 0
//   - WastePercent outside [0, 100]
//   - WasteBuffer outside [0, 100]
//
// Special case: For empty slices (length=0), requires UsedStart=0 & UsedEnd=0.
func (p *SliceReallocationParams) validate(length int) {
	panics.RequireNonNegative(p.UsedStart, "start index")
	panics.RequireNonNegative(p.UsedEnd, "end index")
	if length > 0 {
		panics.RequireLessThan(p.UsedStart, p.UsedEnd, "start index")
		panics.RequireLessThanOrEqualTo(p.UsedEnd, length, "end index")
	} else {
		panics.RequireEqualTo(p.UsedStart, 0, "start index")
		panics.RequireEqualTo(p.UsedEnd, 0, "end index")
	}
	panics.RequireNonNegative(p.MinSize, "min reallocation trigger size")
	panics.RequireNonNegative(p.WastePercent, "waste percent")
	panics.RequireLessThanOrEqualTo(p.WastePercent, 100, "waste percent")
	panics.RequireNonNegative(p.WasteBuffer, "waste buffer")
	panics.RequireLessThanOrEqualTo(p.WasteBuffer, 99, "waste buffer")
}

// Reallocate creates a new slice with reduced capacity to reclaim wasted space.
//
// Reallocation occurs when ALL conditions are met:
//   - Used size >= MinSize (avoid expensive reallocation on small slices)
//   - Waste percent >= WastePercent (enough waste to justify cost)
//
// If reallocation occurs, a new slice with capacity sized to keep waste at
// WasteBuffer% of WastePercent is created, and used elements are copied to
// the new slice starting at index 0. Otherwise, original slice and indices
// are returned unchanged.
//
// Parameters:
//   - data: The underlying slice to reallocate
//   - p: Reallocation parameters controlling when and how to reallocate
//
// Returns:
//   - rData: Reallocated slice (or original if no reallocation)
//   - start: New start index (0 if reallocated, UsedStart otherwise)
//   - end: New end index (len if reallocated, UsedEnd otherwise)
//
// Time complexity:
//   - Best case: O(1) when no reallocation needed
//   - Worst case: O(n) when reallocation occurs (n = used size)
//
// Space complexity:
//   - O(1) when no reallocation
//   - O(n) when reallocation occurs (new slice allocated)
//
// Panics if parameters are invalid.
//
// Example:
//
//	// Deque with high waste
//	data := [_, _, 1, 2, 3, 4, 5, 6, _, ..., _]  // cap=20, used=6, waste=70%
//	rData, start, end := Reallocate(data, SliceReallocationParams{
//	    UsedStart:    2,
//	    UsedEnd:      5,
//	    MinSize:      1,
//	    WastePercent: 50,  // Trigger at 50% waste
//	    WasteBuffer:  80,  // Target 40% waste (80% of 50%)
//	})
//	// Result: rData [1, 2, 3, 4, 5, 6, _, _, _, _], start=0, end=6
//	//         New waste: 40% (4 unused slots out of 10)
//
// Use cases:
//   - Slice-based queues (elements removed from front)
//   - Slice-based stacks (elements removed from back)
//   - Slice-based deques (elements removed from front & back)
//   - Any structure with sliding window over slice
func Reallocate[T any](data []T, p SliceReallocationParams) (rData []T, start int, end int) {
	length := len(data)
	p.validate(length)

	if length == 0 {
		return data, 0, 0
	}

	used := p.UsedEnd - p.UsedStart
	wastePercent := 100 - 100*used/cap(data)
	shouldReallocate := used >= p.MinSize && wastePercent >= p.WastePercent
	if shouldReallocate {
		// Calculate new capacity to keep waste at a fraction of the threshold
		targetPercent := p.WastePercent * p.WasteBuffer / 100
		targetCapacity := max(used*100/(100-targetPercent), 10) // min practical capacity 10
		usedData := data[p.UsedStart:p.UsedEnd]
		rData = make([]T, 0, targetCapacity)
		rData = append(rData, usedData...)
		return rData, 0, len(rData)
	}

	return data, p.UsedStart, p.UsedEnd
}
