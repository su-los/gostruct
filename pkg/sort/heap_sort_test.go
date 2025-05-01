package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHeapSortEmpty(t *testing.T) {
	MaxHeapSort(nil, func(lhs, rhs int) bool {
		return lhs > rhs
	})
	MinHeapSort(nil, func(lhs, rhs int) bool {
		return lhs < rhs
	})
}

func TestMaxHeapSort(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()

		MaxHeapSort(arr, func(lhs, rhs int) bool {
			return lhs > rhs
		})
		require.True(t, CheckIsAsc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestMinHeapSort(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		MinHeapSort(arr, func(lhs, rhs int) bool {
			return lhs < rhs
		})

		require.True(t, CheckIsDesc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}
