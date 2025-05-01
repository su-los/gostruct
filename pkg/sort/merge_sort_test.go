package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMergeSortEmpty(t *testing.T) {
	out := MergeSort[int](nil, OrderAsc)
	require.Empty(t, out)
	out = MergeSort[int](nil, OrderDesc)
	require.Empty(t, out)

	out = MergeSort([]int{}, OrderAsc)
	require.Empty(t, out)
	out = MergeSort([]int{}, OrderDesc)
	require.Empty(t, out)
}

func TestMergeSortAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		out := MergeSort(arr, OrderAsc)
		require.True(t, CheckIsAsc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestMergeSortDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		out := MergeSort(arr, OrderDesc)
		require.True(t, CheckIsDesc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestMergeSortReAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		out := mergeSortRe(arr, OrderAsc)
		require.True(t, CheckIsAsc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestMergeSortReDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		out := mergeSortRe(arr, OrderDesc)
		require.True(t, CheckIsDesc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}
