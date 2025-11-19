package algorithms

import "github.com/apotourlyan/godatastructures/internal/utilities/panics"

// Controls when and how to compact a slice-based data structure.
type SliceCompactionParams struct {
	UsedStart    int // Index of first used element
	MinSize      int // Minimum used size to trigger compaction (0 means always compact if waste threshold is met)
	WastePercent int // Compact if waste >= this percent (0-100)
}

// Validates compaction parameters against slice length.
//
// Panics if parameters are invalid:
//   - UsedStart outside [0, length)
//   - MinSize < 0
//   - WastePercent outside [0, 100]
//
// Special case: For empty slices (length=0), requires UsedStart=0.
func (p *SliceCompactionParams) validate(length int) {
	panics.RequireNonNegative(p.UsedStart, "start index")
	if length > 0 {
		panics.RequireLessThan(p.UsedStart, length, "start index")
	} else {
		panics.RequireEqualTo(p.UsedStart, length, "start index")
	}
	panics.RequireNonNegative(p.MinSize, "min compaction trigger size")
	panics.RequireNonNegative(p.WastePercent, "waste percent")
	panics.RequireLessThanOrEqualTo(p.WastePercent, 100, "waste percent")
}

// Compact shifts elements to the beginning of the slice to reclaim wasted capacity.
//
// Compaction occurs when ALL conditions are met:
//   - Used size >= MinSize (avoid expensive compaction on small ranges)
//   - Waste percent >= WastePercent (enough waste to justify cost)
//   - UsedStart > 0 (not already at beginning)
//
// If compaction occurs, elements at [UsedStart:length] are moved to [0:used],
// the resliced data[:used] and the new start index are returned.
// Otherwise, the original data and start index are returned.
//
// Parameters:
//   - data: The underlying slice to compact (modified in-place if compaction occurs)
//   - p: Compaction parameters controlling when and how to compact
//
// Returns:
//   - cData: Compacted data
//   - start: New index of first used element
//
// Time complexity:
//   - Best case: O(1) when no compaction needed
//   - Worst case: O(n) when compaction occurs (n = used size)
//
// Space complexity: O(1) - compacts in-place
//
// Panics if parameters are invalid.
//
// Example:
//
//	// Queue after many dequeue operations
//	// wasted: 5, used: 3, length: 8
//	data := [_, _, _, _, _, 1, 2, 3]
//	//      ^---wasted---^ ^-used-^
//	params := SliceCompactionParams{
//	  UsedStart:    5,
//	  MinSize:      1,
//	  WastePercent: 50,  // Compact if waste >= 50% length
//	}
//
//	// Waste: 5/8 = 63% >= 50% => compaction triggered
//	data, start := Compact(data, params)
//	// Result: data = [1, 2, 3]  // Re-sliced to used size
//	//         start = 0
//
// Use cases:
//   - Slice-based queues (elements removed from front)
//   - Slice-based deques (elements removed from front & back)
//   - Any structure with sliding window over slice
func Compact[T any](data []T, p SliceCompactionParams) (cData []T, start int) {
	length := len(data)
	p.validate(length)

	if length == 0 {
		return data, 0
	}

	used := length - p.UsedStart
	wastePercent := 100 - 100*used/length
	shouldCompact := used >= p.MinSize &&
		wastePercent >= p.WastePercent &&
		p.UsedStart > 0
	if shouldCompact {
		copy(data, data[p.UsedStart:])
		return data[:used], 0
	}

	return data, p.UsedStart
}
