package structures

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/bench"
)

// Benchmark configurations representing different optimization strategies.
// Used across all benchmarks to compare performance characteristics.
var configs = map[string]SliceQueueConfig{
	// NoOptimizations: Baseline with all optimizations disabled.
	// Expected: Fastest for pure growth, slowest for mixed workloads.
	"NoOptimizations": {
		CompactOnEnqueue:    false,
		ReallocateOnDequeue: false,
	},

	// CompactOnly: Enables compaction to prevent unbounded growth.
	// Expected: Best for balanced and oscillating workloads.
	"CompactOnly": {
		CompactOnEnqueue:      true,
		ReallocateOnDequeue:   false,
		MinOptimizationLength: 100,
		CompactWastePercent:   50,
	},

	// ReallocateOnly: Enables reallocation to free memory.
	// Expected: Best memory reclamation for shrinking workloads.
	"ReallocateOnly": {
		CompactOnEnqueue:       false,
		ReallocateOnDequeue:    true,
		MinOptimizationLength:  100,
		ReallocateWastePercent: 75,
	},

	// BothOptimizations: Enables both strategies.
	// Expected: Balanced performance across all workloads.
	"BothOptimizations": {
		CompactOnEnqueue:       true,
		ReallocateOnDequeue:    true,
		MinOptimizationLength:  100,
		CompactWastePercent:    50,
		ReallocateWastePercent: 75,
	},
}

// BenchmarkSliceQueue_Balanced measures performance with equal enqueue/dequeue operations.
// Queue size remains constant. Tests steady-state performance without growth or shrinkage.
//
// Pattern: [Enqueue, Dequeue] × 500
// Expected winner: CompactOnly (~3x faster, 0 allocations)
func BenchmarkSliceQueue_Balanced(b *testing.B) {
	for name, config := range configs {
		b.Run(name, func(b *testing.B) {
			q := NewSliceQueueWithConfig[int](config)

			for i := range 10000 {
				q.Enqueue(i)
			}

			b.ResetTimer()

			for b.Loop() {
				for j := range 500 {
					q.Enqueue(j)
					q.Dequeue()
				}
			}
		})
	}
}

// BenchmarkSliceQueue_Oscilating measures performance with alternating growth/shrinkage.
// Creates significant waste (70%) then refills. Tests compaction effectiveness.
//
// Pattern: Create 70% waste → [Dequeue × 500, Enqueue × 500]
// Expected winner: CompactOnly (triggers compaction, reuses capacity)
func BenchmarkSliceQueue_Oscilating(b *testing.B) {
	for name, config := range configs {
		b.Run(name, func(b *testing.B) {
			q := NewSliceQueueWithConfig[int](config)

			for i := range 10000 {
				q.Enqueue(i)
			}

			for range 7000 {
				q.Dequeue() // Dequeue 70%, creates 70% waste!
			}

			b.ResetTimer()

			for b.Loop() {
				for range 500 {
					q.Dequeue()
				}

				for j := range 500 {
					q.Enqueue(j)
				}
			}
		})
	}
}

// BenchmarkSliceQueue_MostlyGrowing measures performance with net positive growth.
// 67% enqueue, 33% dequeue. Tests optimization overhead on growing queues.
//
// Pattern: [Enqueue, Enqueue, Dequeue] × 333
// Expected winner: NoOptimizations (compaction overhead hurts growth)
func BenchmarkSliceQueue_MostlyGrowing(b *testing.B) {
	for name, config := range configs {
		b.Run(name, func(b *testing.B) {
			q := NewSliceQueueWithConfig[int](config)

			b.ResetTimer()

			for b.Loop() {
				for j := range 1000 {
					if j%3 == 0 {
						q.Dequeue()
					} else {
						q.Enqueue(j)
					}
				}
			}
		})
	}
}

// BenchmarkSliceQueue_OnlyGrowing measures performance with pure insertion.
// Tests baseline enqueue performance without any dequeue operations.
//
// Pattern: [Enqueue] × 1000
// Expected: All configs similar (no optimization triggers)
func BenchmarkSliceQueue_OnlyGrowing(b *testing.B) {
	for name, config := range configs {
		b.Run(name, func(b *testing.B) {
			q := NewSliceQueueWithConfig[int](config)

			b.ResetTimer()

			for b.Loop() {
				for j := range 1000 {
					q.Enqueue(j)
				}
			}
		})
	}
}

// BenchmarkSliceQueue_MostlyShrinking measures performance with net negative growth.
// 67% dequeue, 33% enqueue. Tests reallocation effectiveness.
//
// Pattern: Start with 1M elements → [Dequeue, Dequeue, Enqueue] × 333
// Expected winner: CompactOnly or ReallocateOnly (depending on refill pattern)
func BenchmarkSliceQueue_MostlyShrinking(b *testing.B) {
	for name, config := range configs {
		b.Run(name, func(b *testing.B) {
			q := NewSliceQueueWithConfig[int](config)

			for i := range 1_000_000 {
				q.Enqueue(i)
			}

			b.ResetTimer()

			for b.Loop() {
				for j := range 1000 {
					if j%3 == 0 {
						q.Enqueue(j)
					} else {
						q.Dequeue()
					}
				}
			}
		})
	}
}

// BenchmarkSliceQueue_OnlyShrinking measures performance with pure removal.
// Tests baseline dequeue performance after large queue.
//
// Pattern: Start with 1M elements → [Dequeue] × 1000
// Expected: All configs similar (optimization in dequeue path rarely helps)
func BenchmarkSliceQueue_OnlyShrinking(b *testing.B) {
	for name, config := range configs {
		b.Run(name, func(b *testing.B) {
			q := NewSliceQueueWithConfig[int](config)

			for i := range 1_000_000 {
				q.Enqueue(i)
			}

			b.ResetTimer()

			for b.Loop() {
				for range 1000 {
					q.Dequeue()
				}
			}
		})
	}
}

// BenchmarkSliceQueue_TotalMemory measures total memory footprint (capacity)
// across different workload patterns. Reports custom metric "total-KB" showing
// actual memory held by the queue after operations.
//
// This benchmark demonstrates when reallocation provides value:
//   - OnlyDequeue: ReallocateOnly frees 97% memory
//   - MostlyDequeue: ReallocateOnly frees 99.97% memory
//   - OnlyEnqueue: All configs same
//   - MostlyEnqueue: CompactOnly prevents unbounded growth
func BenchmarkSliceQueue_TotalMemory(b *testing.B) {
	total := func(q *SliceQueue[int]) float64 {
		return bench.ToKiloBytes(cap(q.data), 8)
	}

	for name, config := range configs {
		q := NewSliceQueueWithConfig[int](config)

		b.Run(name+"/OnlyEnqueue", func(b *testing.B) {
			for i := range 1_000_000 {
				q.Enqueue(i)
			}

			b.ReportMetric(total(q), "total-KB")
		})

		b.Run(name+"/OnlyDequeue", func(b *testing.B) {
			for range 1_000_000 {
				q.Dequeue()
			}

			b.ReportMetric(total(q), "total-KB")
		})

		b.Run(name+"/MostlyEnqueue", func(b *testing.B) {
			for i := range 1_000_000 {
				if i%4 == 0 {
					q.Dequeue()
				} else {
					q.Enqueue(i)
				}
			}

			b.ReportMetric(total(q), "total-KB")
		})

		b.Run(name+"/MostlyDequeue", func(b *testing.B) {
			for i := range 1_000_000 {
				if i%4 == 0 {
					q.Enqueue(i)
				} else {
					q.Dequeue()
				}
			}

			b.ReportMetric(total(q), "total-KB")
		})
	}
}
