# SliceQueue Benchmarking Results

AI analysis by Claude Sonnet 4.5 (verification is advised)

## System Information
- **CPU**: AMD Ryzen 9 7900X3D 12-Core Processor
- **OS**: Linux (amd64)
- **Go Package**: github.com/apotourlyan/godatastructures/internal/queues/structures
- **Date**: 2025-11-06

## Data
[slice-queue-bench-results.txt](slice-queue-bench-results.txt)

## Methodology

All benchmarks run with:
- `-bench=^BenchmarkSliceQueue`: Only SliceQueue benchmarks
- `-benchmem`: Memory allocation tracking
- `-benchtime=100000x`: Fixed 100,000 iterations per benchmark
- `-count=10`: 10 runs for statistical validity

## Test Scenarios

### Balanced
Equal enqueue/dequeue operations maintaining constant queue size.
- Setup: 10,000 initial elements
- Pattern: [Enqueue, Dequeue] × 500 per iteration

### Oscillating  
Alternating growth and shrinkage with high waste.
- Setup: 10,000 elements, dequeue 7,000 (70% waste)
- Pattern: [Dequeue × 500, Enqueue × 500] per iteration

### MostlyGrowing
Net positive growth pattern.
- Pattern: 67% enqueue, 33% dequeue

### OnlyGrowing
Pure insertion workload.
- Pattern: Enqueue × 1000 per iteration

### MostlyShrinking
Net negative growth pattern.
- Setup: 1,000,000 initial elements
- Pattern: 67% dequeue, 33% enqueue

### OnlyShrinking
Pure removal workload.
- Setup: 1,000,000 initial elements
- Pattern: Dequeue × 1000 per iteration

### TotalMemory
Measures actual memory footprint (capacity) after different workload patterns.

## Performance Results Summary

### Balanced Workload

| Configuration | Time (ns/op) | Memory (B/op) | Speed vs NoOpt | Memory vs NoOpt |
|---------------|--------------|---------------|----------------|-----------------|
| **CompactOnly** | **1,413** | **3** | **2.1x faster** ✅ | **7,829x less** ✅ |
| BothOptimizations | 1,593 | 3 | 1.9x faster | 7,829x less |
| NoOptimizations | 3,088 | 23,487 | Baseline | Baseline |
| ReallocateOnly | 4,905 | 21,226 | 0.6x (slower) ❌ | 1.1x less |

**Winner: CompactOnly** - Fastest with near-zero memory allocations

### Oscillating Workload

| Configuration | Time (ns/op) | Memory (B/op) | Speed vs NoOpt | Memory vs NoOpt |
|---------------|--------------|---------------|----------------|-----------------|
| **CompactOnly** | **1,434** | **0** | **2.2x faster** ✅ | **∞ (0 allocs)** ✅ |
| BothOptimizations | 1,618 | 0 | 2.0x faster | ∞ (0 allocs) |
| NoOptimizations | 3,241 | 23,487 | Baseline | Baseline |
| ReallocateOnly | 3,673 | 14,767 | 0.9x (slower) | 1.6x less |

**Winner: CompactOnly** - Fastest with perfect memory reuse (0 allocations)

### MostlyGrowing Workload

| Configuration | Time (ns/op) | Memory (B/op) | Notes |
|---------------|--------------|---------------|-------|
| NoOptimizations | 5,743 | 29,685 | Baseline |
| CompactOnly | 12,686 | 22,173 | **2.2x slower** ❌ |
| ReallocateOnly | 6,294 | 31,458 | Similar to baseline |
| BothOptimizations | 12,824 | 22,173 | **2.2x slower** ❌ |

**Winner: NoOptimizations** - Optimizations hurt growth workloads!

### OnlyGrowing Workload

| Configuration | Time (ns/op) | Memory (B/op) | Notes |
|---------------|--------------|---------------|-------|
| All configs | ~3,200-3,600 | ~16,000-16,200 | Similar performance |

**Result: All similar** - No optimization triggers in pure growth

### MostlyShrinking Workload

| Configuration | Time (ns/op) | Memory (B/op) | Notes |
|---------------|--------------|---------------|-------|
| CompactOnly | 5,870 | 5,291 | Good memory, decent speed |
| BothOptimizations | 6,113 | 5,291 | Similar to CompactOnly |
| NoOptimizations | 6,223 | 21,093 | 4x more memory ❌ |
| ReallocateOnly | 7,475 | 13,553 | Slower, moderate memory |

**Winner: CompactOnly** - Best balance of speed and memory

### OnlyShrinking Workload

| Configuration | Time (ns/op) | Memory (B/op) | Notes |
|---------------|--------------|---------------|-------|
| All configs | ~6,600-7,400 | ~7,800-9,400 | Similar performance |

**Result: All similar** - Dequeue path rarely triggers optimizations

## Memory Footprint Analysis (TotalMemory Benchmarks)

### After OnlyEnqueue (1M enqueues)

| Configuration | Final Capacity |
|---------------|----------------|
| All configs | ~150-188 MB |

**Result: All similar** - Pure growth, no optimization helps

### After OnlyDequeue (dequeue entire queue)

| Configuration | Final Capacity | Memory Freed |
|---------------|----------------|--------------|
| NoOptimizations | **188 MB** | **0%** ❌ |
| CompactOnly | **188 MB** | **0%** ❌ |
| **ReallocateOnly** | **1.4 KB** | **99.999%** ✅ |
| **BothOptimizations** | **0.09 KB** | **99.9999%** ✅ |

**HUGE FINDING:** Reallocation frees essentially all memory!

### After MostlyEnqueue (net growth)

| Configuration | Final Capacity |
|---------------|----------------|
| NoOptimizations | 294 MB (unbounded growth!) ❌ |
| CompactOnly | 129 MB ✅ |
| ReallocateOnly | 210 MB |
| BothOptimizations | 129 MB ✅ |

**Finding:** Compaction prevents unbounded memory growth (2.3x less memory)

### After MostlyDequeue (net shrinkage)

| Configuration | Final Capacity | Memory Freed |
|---------------|----------------|--------------|
| NoOptimizations | **367 MB** | 0% (keeps growing!) ❌ |
| CompactOnly | 129 MB | 65% ✅ |
| **ReallocateOnly** | **16-64 KB** | **99.98%** ✅ |
| **BothOptimizations** | **0.09 KB** | **99.9999%** ✅ |

**Finding:** Reallocation achieves massive memory reclamation!

## Key Findings

### 1. CompactOnly Dominates Balanced/Oscillating Workloads

**Speed improvement:**
- Balanced: **2.1x faster** (3,088 → 1,413 ns/op)
- Oscillating: **2.2x faster** (3,241 → 1,434 ns/op)

**Memory improvement:**
- Balanced: **7,829x less** (23,487 → 3 B/op)
- Oscillating: **Perfect reuse** (0 B/op)

**Why:** Compaction resets the `curr` pointer, allowing capacity reuse without allocation.

### 2. Reallocation Critical for Shrinkage

**Memory freed:**
- OnlyDequeue: **99.9999%** (188 MB → 0.09 KB)
- MostlyDequeue: **99.9999%** (367 MB → 0.09 KB)

**Why:** CompactOnly can't shrink capacity - only reuse it. Reallocation is essential for memory-constrained environments.

### 3. Optimizations Hurt Pure Growth

**MostlyGrowing:**
- NoOptimizations: 5,743 ns/op (fastest)
- CompactOnly: 12,686 ns/op (**2.2x slower**)

**Why:** Compaction overhead > benefit when queue continuously grows.

### 4. BothOptimizations = Safe Default

Handles all patterns reasonably:
- ✅ Balanced: Nearly as fast as CompactOnly (1,593 vs 1,413 ns)
- ✅ Oscillating: Nearly as fast as CompactOnly (1,618 vs 1,434 ns)
- ✅ Shrinking: Maximum memory reclamation (99.9999%)
- ❌ Growing: Slower like CompactOnly (optimization overhead)

## Recommendations

### Choose Configuration Based on Workload:

| Workload Pattern | Recommended Config | Reason |
|------------------|-------------------|--------|
| **Balanced ops** | CompactOnly | 2.1x faster, 7,829x less memory |
| **Oscillating size** | CompactOnly | 2.2x faster, 0 allocations |
| **Pure growth** | NoOptimizations | Fastest (avoid optimization overhead) |
| **Permanent shrinkage** | BothOptimizations or ReallocateOnly | 99.9999% memory freed |
| **Unknown/Mixed** | **BothOptimizations** | Safe default for all patterns |

### Default Configuration

```go
// Recommended default (BothOptimizations)
config := SliceQueueConfig{
    CompactOnEnqueue:       true,
    ReallocateOnDequeue:    true,
    MinOptimizationLength:  100,
    CompactWastePercent:    50,
    ReallocateWastePercent: 75,
}
```

## Reproduction

```bash
# Run complete benchmark suite
go test -bench=^BenchmarkSliceQueue -benchmem -benchtime=100000x -count=10 \
  ./internal/queues/structures > slice_queue_bench_results.txt

# Compare configurations with benchstat
go install golang.org/x/perf/cmd/benchstat@latest
benchstat results_before.txt results_after.txt
```

## Statistical Analysis

With 10 runs per benchmark, results have statistical validity:
- Variance is typically < 5% for speed measurements
- Memory allocations are consistent across runs
- Large performance differences (>2x) are highly significant
```

## Conclusion

The benchmarks prove clear performance characteristics:

1. **CompactOnly is optimal for balanced/oscillating workloads** (2x faster, near-zero allocations)
2. **Reallocation is essential for memory reclamation** (99.9999% memory freed after shrinkage)
3. **Optimizations have costs** - they hurt pure growth workloads (2x slower)
4. **BothOptimizations is the safest default** - handles all patterns reasonably

These results demonstrate workload-dependent optimization tradeoffs and validate the configurable design approach.