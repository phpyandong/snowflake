package flake

import (
	"testing"
)

//////////////////////////////////////////////////////////////////////////////

// Benchmarks Presence Update event with fake data.
func BenchmarkGenerate(b *testing.B) {

	node, _ := NewFlakeNode(1)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, _ = node.Generate()
	}

}

// Benchmarks Presence Update event with fake data.
func BenchmarkGenerateLocks(b *testing.B) {

	node, _ := NewFlakeNode(1)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, _ = node.LockedGenerate()
	}

}

// Benchmarks Presence Update event with fake data.
func BenchmarkGenerateLocksParallel(b *testing.B) {

	node, _ := NewFlakeNode(1)

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = node.LockedGenerate()
		}
	})
}
