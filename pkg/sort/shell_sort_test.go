package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestShellSortEmpty(t *testing.T) {
	ShellSort[int](nil, OrderAsc)
	ShellSort[int](nil, OrderDesc)

	ShellSort([]int{}, OrderAsc)
	ShellSort([]int{}, OrderDesc)
}

func TestShellSortAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()

		ShellSort(arr, OrderAsc)
		require.True(t, CheckIsAsc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestShellSortDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		ShellSort(arr, OrderDesc)
		require.True(t, CheckIsDesc(arr))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}
