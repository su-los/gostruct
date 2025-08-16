package segtree

import "testing"

func BenchmarkBuild(b *testing.B) {
	b.ResetTimer()

	for i := range b.N {
		b.StopTimer()
		cnt := min((i+1)*100, 100000)
		nums := GenNumList(cnt, 100000)
		b.StartTimer()
		build(nums, Sum)
	}
}

func BenchmarkBuildRecursive(b *testing.B) {
	b.ResetTimer()
	for i := range b.N {
		b.StopTimer()
		cnt := min((i+1)*100, 100000)
		nums := GenNumList(cnt, 100000)
		b.StartTimer()
		buildRecursive(nums, 0, len(nums), Sum)
	}
}

func BenchmarkBuildBottomUp(b *testing.B) {
	b.ResetTimer()
	for i := range b.N {
		b.StopTimer()

		cnt := min((i+1)*100, 100000)
		nums := GenNumList(cnt, 100000)
		b.StartTimer()
		buildBottomUp(nums, Sum)
	}
}
