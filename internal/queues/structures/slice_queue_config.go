package structures

// SliceQueueConfig controls memory optimization behavior for SliceQueue.
//
// The queue supports two independent optimization strategies that can be
// enabled or disabled based on workload characteristics:
//
// 1. Compaction (Enqueue-time optimization):
//
// Shifts active elements to the front of the slice when dead space exceeds
// a threshold, enabling capacity reuse and preventing unbounded growth.
//
// 2. Reallocation (Dequeue-time optimization):
//
// Shrinks the underlying slice capacity when waste becomes excessive,
// freeing memory for permanently shrinking queues.
//
// Default configuration (NewSliceQueue):
//
//	CompactOnEnqueue:       true   // prevent unbounded growth
//	ReallocateOnDequeue:    true   // enable memory reclamation
//	MinOptimizationLength:  100    // avoid optimizing tiny queues
//	CompactWastePercent:    50     // compact when 50%+ waste
//	ReallocateWastePercent: 75     // reallocate when 75%+ waste
//
// Example configurations:
//
//	// High-throughput server (favor speed)
//	config := SliceQueueConfig{
//	    CompactOnEnqueue:      true,
//	    ReallocateOnDequeue:   false,  // Skip reallocation overhead
//	    MinOptimizationLength: 1000,   // Only optimize large queues
//	    CompactWastePercent:   60,     // Tolerate more waste
//	}
//
//	// Memory-constrained environment
//	config := SliceQueueConfig{
//	    CompactOnEnqueue:       true,
//	    ReallocateOnDequeue:    true,   // Aggressively free memory
//	    MinOptimizationLength:  50,     // Optimize even small queues
//	    CompactWastePercent:    40,     // Lower waste tolerance
//	    ReallocateWastePercent: 60,     // Earlier reallocation
//	}
//
//	// Pure growth workload (disable all optimizations)
//	config := SliceQueueConfig{
//	    CompactOnEnqueue:    false,
//	    ReallocateOnDequeue: false,
//	}
type SliceQueueConfig struct {
	// CompactOnEnqueue enables compaction during enqueue operations.
	// When enabled, shifts active elements to the front of the slice
	// if waste exceeds CompactWastePercent.
	//
	// Cost: O(n) copy operation when triggered
	//
	// Benefit: Prevents unbounded growth, enables capacity reuse
	//
	// Triggers: Only when size >= MinOptimizationLength and waste > threshold
	CompactOnEnqueue bool

	// ReallocateOnDequeue enables capacity shrinking during dequeue operations.
	// When enabled, allocates a smaller slice and copies active elements
	// if waste exceeds ReallocateWastePercent.
	//
	// Cost: O(n) allocation + copy when triggered
	//
	// Benefit: Frees memory for permanently shrinking queues
	//
	// Triggers: Only when capacity >= MinOptimizationLength and waste > threshold
	ReallocateOnDequeue bool

	// MinOptimizationLength is the minimum queue capacity before optimizations
	// are considered. Prevents optimization overhead on small queues.
	//
	// Recommended values:
	//   50-100:   General purpose
	//   500-1000: High-throughput systems (avoid optimization overhead)
	//   10-50:    Memory-constrained environments
	MinOptimizationLength int

	// CompactWastePercent is the waste threshold (as percentage) that triggers
	// compaction during enqueue operations.
	//
	// Waste is calculated as: 100 * curr / len(data)
	// where curr is the index of the first active element.
	//
	// Lower values: More aggressive compaction, less memory waste, higher CPU
	// Higher values: Less frequent compaction, more memory waste, lower CPU
	//
	// Recommended values:
	//   40-50: Balanced (default: 50)
	//   30-40: Memory-constrained
	//   60-70: CPU-constrained
	CompactWastePercent int

	// ReallocateWastePercent is the waste threshold (as percentage) that triggers
	// reallocation during dequeue operations.
	//
	// Waste is calculated as: 100 * (1 - size/capacity)
	//
	// Lower values: More aggressive reallocation, better memory reclamation, higher CPU
	// Higher values: Less frequent reallocation, slower memory reclamation, lower CPU
	//
	// Recommended values:
	//   70-80: Balanced (default: 75)
	//   60-70: Memory-constrained
	//   80-90: CPU-constrained
	//
	// Note: Should be higher than CompactWastePercent to avoid conflicts
	ReallocateWastePercent int
}
