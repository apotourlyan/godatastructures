package structures

import (
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/bench"
)

var configs = map[string]SliceQueueConfig{
	"NoOptimizations": {
		CompactOnEnqueue:    false,
		ReallocateOnDequeue: false,
	},
	"CompactOnly": {
		CompactOnEnqueue:      true,
		ReallocateOnDequeue:   false,
		MinOptimizationLength: 100,
		CompactWastePercent:   50,
	},
	"ReallocateOnly": {
		CompactOnEnqueue:       false,
		ReallocateOnDequeue:    true,
		MinOptimizationLength:  100,
		ReallocateWastePercent: 75,
	},
	"BothOptimizations": {
		CompactOnEnqueue:       true,
		ReallocateOnDequeue:    true,
		MinOptimizationLength:  100,
		CompactWastePercent:    50,
		ReallocateWastePercent: 75,
	},
}

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
