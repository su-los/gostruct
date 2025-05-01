package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCountingSortAsc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		out := CountingSort(arr, 5000, OrderAsc)

		require.True(t, CheckIsAsc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestCountingSortDesc(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(10000, 5000)
		start := time.Now()
		out := CountingSort(arr, 5000, OrderDesc)
		require.True(t, CheckIsDesc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds(), "ms")
	}
}

func TestCountingSortUnsigned(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(100000, 50000)
		usigarr := make([]uint, len(arr))
		for idx, val := range arr {
			usigarr[idx] = uint(val)
		}
		start := time.Now()
		out := CountingSortUnsigned(usigarr, OrderAsc)
		require.True(t, CheckIsAsc(out))

		out = CountingSortUnsigned(usigarr, OrderDesc)
		require.True(t, CheckIsDesc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds()/2, "ms")
	}
}

func TestCountingSortSigned(t *testing.T) {
	rd := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	const testcnt = 100
	for i := range testcnt {
		arr := GenNumList(100000, 50000)
		sigarr := make([]int, len(arr))
		for idx, val := range arr {
			r := rd.Intn(2)
			if r == 0 {
				val = -val
			}
			sigarr[idx] = int(val)
		}
		start := time.Now()
		out := CountingSortSigned(sigarr, OrderAsc)
		require.True(t, CheckIsAsc(out))

		out = CountingSortSigned(sigarr, OrderDesc)
		require.True(t, CheckIsDesc(out))
		fmt.Println("[INFO] Success testcase", i, "(len=", len(arr), ") in", time.Since(start).Milliseconds()/2, "ms")
	}
}
