package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestInsertionSortEmpty(t *testing.T) {
	InsertionSort[int](nil, OrderAsc)
	InsertionSort[int](nil, OrderDesc)
	InsertionSort([]int{}, OrderAsc)
	InsertionSort([]int{}, OrderDesc)
}

func TestInserttionSortAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		InsertionSort(arr, OrderAsc)

		require.True(t, CheckIsAsc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestInsertionSortDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		InsertionSort(arr, OrderDesc)
		require.True(t, CheckIsDesc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}
