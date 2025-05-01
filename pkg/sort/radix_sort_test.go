package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRadixSortEmpty(t *testing.T) {
	RadixSort[int](nil, OrderAsc)
	RadixSort[int](nil, OrderDesc)
	RadixSort([]int{}, OrderAsc)
	RadixSort([]int{}, OrderDesc)
}

func TestRadixSortAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		res := RadixSort(arr, OrderAsc)

		require.True(t, CheckIsAsc(res))
		require.Len(t, res, len(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(res), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestRadixSortDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		res := RadixSort(arr, OrderDesc)
		require.True(t, CheckIsDesc(res))
		require.Len(t, res, len(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestRadixSortCountAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		res := RadixSortByCounting(arr, OrderAsc)
		require.True(t, CheckIsAsc(res))
		require.Len(t, res, len(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(res), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestRadixSortCountDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		res := RadixSortByCounting(arr, OrderDesc)
		require.True(t, CheckIsDesc(res))
		require.Len(t, res, len(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func BenchmarkRadixSortAsc(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		b.StopTimer()
		arr := GenNumList(10000, 5000)

		b.StartTimer()
		_ = RadixSort(arr, OrderAsc)
	}
}

func BenchmarkRadixSortCountAsc(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		b.StopTimer()
		arr := GenNumList(10000, 5000)

		b.StartTimer()
		_ = RadixSortByCounting(arr, OrderAsc)
	}
}
