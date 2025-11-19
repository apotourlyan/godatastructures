package algorithms

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/test"
)

// Test Coverage
// =============
// Reallocate:
//  ✓ Negative start index
//  ✓ Negative end index
//  ✓ Start index greater than or equal to end index
//  ✓ End index greater than length
//  ✓ Empty slice with nonzero start
//  ✓ Empty slice with nonzero end
//  ✓ Negative min size
//  ✓ Negative waste percent
//  ✓ Waste percent greater than 100
//  ✓ Negative waste buffer
//  ✓ Waste buffer equals 100
//  ✓ Empty slice
//  ✓ Used size below min size
//  ✓ Waste below threshold
//  ✓ Waste just below threshold
//  ✓ Standard reallocation
//  ✓ Min size boundary
//  ✓ Waste percent boundary
//  ✓ Min size zero with waste above threshold
//  ✓ Waste percent zero with any waste
//  ✓ High waste buffer value

// Verifies that Reallocate panics with appropriate error messages for invalid parameters
func TestReallocate_InvalidArgs(t *testing.T) {
	cases := []struct {
		name string
		data []int
		p    SliceReallocationParams
		want string
	}{
		{
			name: "negative_start_index",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    -1,
				UsedEnd:      3,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  80,
			},
			want: `"start index" must be >= 0, got -1`,
		},
		{
			name: "negative_end_index",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      -1,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  80,
			},
			want: `"end index" must be >= 0, got -1`,
		},
		{
			name: "start_index_greater_than_or_equal_to_end_index",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    2,
				UsedEnd:      2,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  80,
			},
			want: `"start index" must be < 2, got 2`,
		},
		{
			name: "end_index_greater_than_length",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      5,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  80,
			},
			want: `"end index" must be <= 3, got 5`,
		},
		{
			name: "empty_slice_with_nonzero_start",
			data: []int{},
			p: SliceReallocationParams{
				UsedStart:    1,
				UsedEnd:      0,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  80,
			},
			want: `"start index" must be == 0, got 1`,
		},
		{
			name: "empty_slice_with_nonzero_end",
			data: []int{},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      1,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  80,
			},
			want: `"end index" must be == 0, got 1`,
		},
		{
			name: "negative_min_size",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      3,
				MinSize:      -5,
				WastePercent: 50,
				WasteBuffer:  80,
			},
			want: `"min reallocation trigger size" must be >= 0, got -5`,
		},
		{
			name: "negative_waste_percent",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      3,
				MinSize:      1,
				WastePercent: -10,
				WasteBuffer:  80,
			},
			want: `"waste percent" must be >= 0, got -10`,
		},
		{
			name: "waste_percent_greater_than_100",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      3,
				MinSize:      1,
				WastePercent: 150,
				WasteBuffer:  80,
			},
			want: `"waste percent" must be <= 100, got 150`,
		},
		{
			name: "negative_waste_buffer",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      3,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  -10,
			},
			want: `"waste buffer" must be >= 0, got -10`,
		},
		{
			name: "waste_buffer_equals_100",
			data: []int{1, 2, 3},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      3,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  100,
			},
			want: `"waste buffer" must be <= 99, got 100`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			test.GotWantPanic(t, func() {
				Reallocate(c.data, c.p)
			}, c.want)
		})
	}
}

// Verifies that Reallocate returns unchanged data when reallocation conditions are not met
func TestReallocate_NotTriggered(t *testing.T) {
	cases := []struct {
		name string
		data []int
		p    SliceReallocationParams
	}{
		{
			name: "empty_slice",
			data: []int{},
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      0,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  80,
			},
		},
		{
			name: "used_size_below_min_size",
			// cap=20, len=10, used=2 (indices 3-5), waste=90%
			data: func() []int {
				data := make([]int, 10, 20)
				data[3] = 1
				data[4] = 2
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    3,
				UsedEnd:      5,
				MinSize:      5,  // ← Testing: 2 < 5
				WastePercent: 50, // ✓ 90% >= 50%
				WasteBuffer:  80,
			},
		},
		{
			name: "waste_below_threshold",
			// cap=10, len=10, used=7 (indices 0-7), waste=30%
			data: func() []int {
				data := make([]int, 10)
				for i := range 7 {
					data[i] = i + 1
				}
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      7,
				MinSize:      5,  // ✓ 7 >= 5
				WastePercent: 50, // ← Testing: 30% < 50%
				WasteBuffer:  80,
			},
		},
		{
			name: "waste_just_below_threshold",
			// cap=10, len=10, used=5 (indices 0-5), waste=50%
			data: func() []int {
				data := make([]int, 10)
				for i := range 5 {
					data[i] = i + 1
				}
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      5,
				MinSize:      1,  // ✓ 5 >= 1
				WastePercent: 51, // ← Testing: 50% < 51% (boundary)
				WasteBuffer:  80,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, start, end := Reallocate(c.data, c.p)
			test.GotWantSlice(t, data, c.data)
			test.GotWant(t, start, c.p.UsedStart)
			test.GotWant(t, end, c.p.UsedEnd)
		})
	}
}

// Verifies that Reallocate correctly shifts elements to the start and returns reallocated slice
func TestReallocate_Triggered(t *testing.T) {
	cases := []struct {
		name     string
		data     []int
		p        SliceReallocationParams
		wantData []int
		wantLen  int
		wantCap  int
	}{
		{
			name: "standard_reallocation",
			// cap=20, len=10, used=5 (indices 2-7), waste=75%
			data: func() []int {
				data := make([]int, 10, 20)
				for i := 2; i < 7; i++ {
					data[i] = i - 1
				}
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    2,
				UsedEnd:      7,
				MinSize:      3,  // ✓ 5 >= 3
				WastePercent: 50, // ✓ 75% >= 50%
				WasteBuffer:  80, // Target 40% waste
			},
			wantData: []int{1, 2, 3, 4, 5},
			wantLen:  5,
			wantCap:  10, // max(5*100/60, 10) = 10
		},
		{
			name: "min_size_boundary",
			// cap=20, len=10, used=3 (indices 0-3), waste=85%
			data: func() []int {
				data := make([]int, 10, 20)
				data[0] = 1
				data[1] = 2
				data[2] = 3
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      3,
				MinSize:      3,  // ← Testing: 3 >= 3 (boundary)
				WastePercent: 50, // ✓ 85% >= 50%
				WasteBuffer:  80,
			},
			wantData: []int{1, 2, 3},
			wantLen:  3,
			wantCap:  10, // max(3*100/60, 10) = 10
		},
		{
			name: "waste_percent_boundary",
			// cap=10, len=10, used=5 (indices 0-5), waste=50%
			data: func() []int {
				data := make([]int, 10)
				for i := range 5 {
					data[i] = i + 1
				}
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      5,
				MinSize:      1,
				WastePercent: 50, // ← Testing: 50% >= 50% (boundary)
				WasteBuffer:  80,
			},
			wantData: []int{1, 2, 3, 4, 5},
			wantLen:  5,
			wantCap:  10, // max(5*100/60, 10) = 10
		},
		{
			name: "min_size_zero_with_waste_above_threshold",
			// cap=20, len=5, used=1 (index 4), waste=95%
			data: func() []int {
				data := make([]int, 5, 20)
				data[4] = 1
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    4,
				UsedEnd:      5,
				MinSize:      0,  // ← Testing: 1 >= 0 (edge case)
				WastePercent: 50, // ✓ 95% >= 50%
				WasteBuffer:  80,
			},
			wantData: []int{1},
			wantLen:  1,
			wantCap:  10, // max(1*100/60, 10) = 10
		},
		{
			name: "waste_percent_zero_with_any_waste",
			// cap=10, len=5, used=4 (indices 1-5), waste=60%
			data: func() []int {
				data := make([]int, 5, 10)
				data[1] = 1
				data[2] = 2
				data[3] = 3
				data[4] = 4
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    1,
				UsedEnd:      5,
				MinSize:      1, // ✓ 4 >= 1
				WastePercent: 0, // ← Testing: any waste triggers (60% >= 0%)
				WasteBuffer:  80,
			},
			wantData: []int{1, 2, 3, 4},
			wantLen:  4,
			wantCap:  10, // 4*100/100 = 4, max(4, 10) = 10
		},
		{
			name: "high_waste_buffer_value",
			// cap=100, len=50, used=10 (indices 0-10), waste=90%
			data: func() []int {
				data := make([]int, 50, 100)
				for i := range 10 {
					data[i] = i + 1
				}
				return data
			}(),
			p: SliceReallocationParams{
				UsedStart:    0,
				UsedEnd:      10,
				MinSize:      1,
				WastePercent: 50,
				WasteBuffer:  99, // ← Testing high buffer (target 49.5% waste)
			},
			wantData: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantLen:  10,
			wantCap:  19, // max(10*100/51, 10) = 19
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotData, start, end := Reallocate(c.data, c.p)
			test.GotWantSlice(t, gotData, c.wantData)
			test.GotWant(t, start, 0)
			test.GotWant(t, end, c.wantLen)
			test.GotWant(t, cap(gotData), c.wantCap)
		})
	}
}
