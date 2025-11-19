package structures

// SliceStackConfig controls memory optimization behavior for SliceStack.
//
// The stack supports one optional optimization strategy:
//
// Reallocation (Pop-time optimization):
//
// Shrinks the underlying slice capacity when waste exceeds a threshold,
// freeing memory for stacks that grow large then permanently shrink.
// Adds a one-time O(n) cost during the Pop operation that triggers
// reallocation.
type SliceStackConfig struct {
	// ReallocateOnPop enables slice reallocation after Pop operations.
	//
	// When enabled, the stack will reallocate its underlying slice when waste
	// exceeds ReallocateWastePercent and the used size is at least
	// MinOptimizationLength elements.
	//
	// This reduces memory usage for stacks that shrink significantly but adds
	// a one-time O(n) cost during the Pop that triggers reallocation.
	ReallocateOnPop bool

	// MinOptimizationLength represents the minimum stack size to trigger reallocation.
	//
	// Prevents expensive reallocations on small stacks where the overhead
	// outweighs the memory savings.
	//
	//   50-100:   General purpose
	//   500-1000: High-throughput systems (avoid optimization overhead)
	//   10-50:    Memory-constrained environments
	MinOptimizationLength int

	// ReallocateWastePercent represents the waste threshold to trigger reallocation (0-100).
	//
	// Reallocation occurs when:
	//   waste% = (capacity - size) / capacity >= ReallocateWastePercent
	//
	// Example: With 75%, a stack with capacity 100 and size 20 has 80% waste,
	// so reallocation will trigger.
	//
	// Lower values: More frequent reallocation, better memory reclamation, higher CPU
	// Higher values: Less frequent reallocation, slower memory reclamation, lower CPU
	//
	// Recommended values:
	//   70-80: Balanced (default: 75)
	//   60-70: Memory-constrained
	//   80-90: CPU-constrained
	ReallocateWastePercent int

	// ReallocateWasteBuffer controls target waste after reallocation.
	//
	// When reallocation triggers, the new capacity is sized to achieve waste
	// at WasteBuffer% of ReallocateWastePercent. This determines how much
	// headroom exists before the next reallocation trigger.
	//
	// Formula: target waste = ReallocateWastePercent * ReallocateWasteBuffer / 100
	//
	// Recommended values:
	//   - 80: Good balance - reasonable headroom
	//   - 50-70: Conservative - fewer reallocations, more memory usage
	//   - 90: Aggressive - lower memory usage, more reallocations
	//
	// Valid range: [0, 99]
	ReallocateWasteBuffer int
}
