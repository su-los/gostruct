package segtree

import (
	"math/rand/v2"
	"testing"
	"time"
)

func BenchmarkQuery(b *testing.B) {
	rd := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	fArr := []AggFunc[int]{Sum[int], Max[int], Min[int]}

	b.ResetTimer()
	for i := range b.N {
		b.StopTimer()
		cnt := min((i+1)*100, 1000000)
		nums := GenNumList(cnt, 1000000)

		fidx := i % 3
		f := fArr[fidx]
		st := NewSegTree(nums, f)

		// 随机生成查询的区间
		start := rd.IntN(cnt)
		end := rd.IntN(cnt - start)
		end = min(end+start, cnt)

		b.StartTimer()
		st.Query(start, end)
	}
}
