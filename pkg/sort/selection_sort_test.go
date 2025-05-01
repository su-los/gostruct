package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSelectionSortAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		SelectionSort(arr, OrderAsc)
		require.True(t, CheckIsAsc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestSelectionSortDesc(t *testing.T) {
	// 空序列
	SelectionSort[int](nil, OrderAsc)
	SelectionSort([]int{}, OrderDesc)

	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		SelectionSort(arr, OrderDesc)
		require.True(t, CheckIsDesc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}
