package algorithms

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Test Coverage
// =============
// Compact:
//  ✓ Negative start index
//  ✓ Start index equals length
//  ✓ Start index greater than length
//  ✓ Empty slice with nonzero start
//  ✓ Negative min size
//  ✓ Negative waste percent
//  ✓ Waste percent greater than 100
//  ✓ Empty slice
//  ✓ Used size below min size
//  ✓ Waste below threshold
//  ✓ Waste just below threshold
//  ✓ Min size zero but waste below threshold
//  ✓ No waste already at start
//  ✓ Standard compaction
//  ✓ Min size boundary
//  ✓ Waste percent boundary
//  ✓ Min size zero with waste above threshold
//  ✓ Waste percent zero with any waste

// Verifies that Compact panics with appropriate error messages for invalid parameters
func TestCompact_InvalidArgs(t *testing.T) {
	cases := []struct {
		name string
		data []int
		p    SliceCompactionParams
		want string
	}{
		{
			name: "negative_start_index",
			data: []int{1, 2, 3},
			p: SliceCompactionParams{
				UsedStart:    -1,
				MinSize:      1,
				WastePercent: 50,
			},
			want: `"start index" must be >= 0, got -1`,
		},
		{
			name: "start_index_equals_length",
			data: []int{1, 2, 3},
			p: SliceCompactionParams{
				UsedStart:    3,
				MinSize:      1,
				WastePercent: 50,
			},
			want: `"start index" must be < 3, got 3`,
		},
		{
			name: "start_index_greater_than_length",
			data: []int{1, 2, 3},
			p: SliceCompactionParams{
				UsedStart:    5,
				MinSize:      1,
				WastePercent: 50,
			},
			want: `"start index" must be < 3, got 5`,
		},
		{
			name: "empty_slice_with_nonzero_start",
			data: []int{},
			p: SliceCompactionParams{
				UsedStart:    1,
				MinSize:      1,
				WastePercent: 50,
			},
			want: `"start index" must be == 0, got 1`,
		},
		{
			name: "negative_min_size",
			data: []int{1, 2, 3},
			p: SliceCompactionParams{
				UsedStart:    0,
				MinSize:      -5,
				WastePercent: 50,
			},
			want: `"min compaction trigger size" must be >= 0, got -5`,
		},
		{
			name: "negative_waste_percent",
			data: []int{1, 2, 3},
			p: SliceCompactionParams{
				UsedStart:    0,
				MinSize:      1,
				WastePercent: -10,
			},
			want: `"waste percent" must be >= 0, got -10`,
		},
		{
			name: "waste_percent_greater_than_100",
			data: []int{1, 2, 3},
			p: SliceCompactionParams{
				UsedStart:    0,
				MinSize:      1,
				WastePercent: 150,
			},
			want: `"waste percent" must be <= 100, got 150`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			test.GotWantPanic(t, func() {
				Compact(c.data, c.p)
			}, c.want)
		})
	}
}

// Verifies that Compact returns unchanged data when compaction conditions are not met
func TestCompact_NotTriggered(t *testing.T) {
	cases := []struct {
		name string
		data []int
		p    SliceCompactionParams
	}{
		{
			name: "empty_slice",
			data: []int{},
			p: SliceCompactionParams{
				UsedStart:    0,
				MinSize:      1,
				WastePercent: 50,
			},
		},
		{
			name: "used_size_below_min_size",
			data: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 2}, // length=10, used=2, waste=80%
			p: SliceCompactionParams{
				UsedStart:    8,
				MinSize:      5,  // ← Testing: 2 < 5
				WastePercent: 50, // ✓ Waste: 80% >= 50%
			},
		},
		{
			name: "waste_below_threshold",
			data: []int{0, 0, 0, 1, 2, 3, 4, 5, 6, 7}, // length=10, used=7, waste=30%
			p: SliceCompactionParams{
				UsedStart:    3,
				MinSize:      5,  // ✓ Used: 7 >= 5
				WastePercent: 50, // ← Testing: 30% < 50%
			},
		},
		{
			name: "waste_just_below_threshold",
			data: []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5}, // length=10, used=5, waste=50%
			p: SliceCompactionParams{
				UsedStart:    5,
				MinSize:      1,  // ✓ Used: 5 >= 1
				WastePercent: 51, // ← Testing: 50% < 51% (boundary)
			},
		},
		{
			name: "min_size_zero_but_waste_below_threshold",
			data: []int{0, 1, 2, 3, 4}, // length=5, used=4, waste=20%
			p: SliceCompactionParams{
				UsedStart:    1,
				MinSize:      0,  // ✓ Used: 4 >= 0 (edge case)
				WastePercent: 50, // ← Testing: 20% < 50%
			},
		},
		{
			name: "no_waste_at_all",
			data: []int{1, 2, 3, 4, 5}, // length=5, used=5, waste=0%
			p: SliceCompactionParams{
				UsedStart:    0,
				MinSize:      1, // ✓ Used: 5 >= 1
				WastePercent: 0, // ← 0% >= 0%, but UsedStart=0 prevents compaction
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, start := Compact(c.data, c.p)
			test.GotWantSlice(t, data, c.data)
			test.GotWant(t, start, c.p.UsedStart)
		})
	}
}

// Verifies that Compact correctly shifts elements to the start and returns compacted slice
func TestCompact_Triggered(t *testing.T) {
	cases := []struct {
		name     string
		data     []int
		p        SliceCompactionParams
		wantData []int
	}{
		{
			name: "standard_compaction",
			data: []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5}, // length=10, used=5, waste=50%
			p: SliceCompactionParams{
				UsedStart:    5,  // ✓ > 0
				MinSize:      3,  // ✓ 5 >= 3
				WastePercent: 40, // ✓ 50% >= 40%
			},
			wantData: []int{1, 2, 3, 4, 5},
		},
		{
			name: "min_size_boundary",
			data: []int{0, 0, 0, 1, 2, 3}, // length=6, used=3, waste=50%
			p: SliceCompactionParams{
				UsedStart:    3,  // ✓ > 0
				MinSize:      3,  // ← Testing: 3 >= 3 (boundary)
				WastePercent: 40, // ✓ 50% >= 40%
			},
			wantData: []int{1, 2, 3},
		},
		{
			name: "waste_percent_boundary",
			data: []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5}, // length=10, used=5, waste=50%
			p: SliceCompactionParams{
				UsedStart:    5,  // ✓ > 0
				MinSize:      1,  // ✓ 5 >= 1
				WastePercent: 50, // ← Testing: 50% >= 50% (boundary)
			},
			wantData: []int{1, 2, 3, 4, 5},
		},
		{
			name: "min_size_zero_with_waste_above_threshold",
			data: []int{0, 0, 0, 0, 1}, // length=5, used=1, waste=80%
			p: SliceCompactionParams{
				UsedStart:    4,  // ✓ > 0
				MinSize:      0,  // ← Testing: 1 >= 0 (edge case)
				WastePercent: 50, // ✓ 80% >= 50%
			},
			wantData: []int{1},
		},
		{
			name: "waste_percent_zero_with_any_waste",
			data: []int{0, 1, 2, 3, 4}, // length=5, used=4, waste=20%
			p: SliceCompactionParams{
				UsedStart:    1, // ✓ > 0 (has waste)
				MinSize:      1, // ✓ 4 >= 1
				WastePercent: 0, // ← Testing: 20% >= 0% (any waste triggers)
			},
			wantData: []int{1, 2, 3, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotData, start := Compact(c.data, c.p)
			test.GotWantSlice(t, gotData, c.wantData)
			test.GotWant(t, start, 0)
		})
	}
}
