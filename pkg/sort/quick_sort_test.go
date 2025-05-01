package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQuickSortEmqty(t *testing.T) {
	QuickSort[int](nil, OrderAsc)
	QuickSort[int](nil, OrderDesc)

	QuickSort([]int{}, OrderAsc)
	QuickSort([]int{}, OrderDesc)
}

func TestQuickSortAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		QuickSort(arr, OrderAsc)
		require.True(t, CheckIsAsc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestQuickSortDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		QuickSort(arr, OrderDesc)
		require.True(t, CheckIsDesc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}
