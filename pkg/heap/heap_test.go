package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/su-los/gostruct/pkg/sort"
)

func TestMaxHeap(t *testing.T) {
	// 随机生成非空序列，创建最大堆
	t.Run("Random", func(t *testing.T) {
		const testcnt = 100
		for i := range testcnt {
			list := GenList(1000, 10000)

			start := time.Now()
			h := NewMaxHeap(list)
			require.True(t, h.IsValidHeap())
			require.Equal(t, h.data[0], lo.Max(list))
			fmt.Println("[INFO] Success testcase", i, " in", time.Since(start).Milliseconds(), "ms")
		}
	})

	// 序列为空
	t.Run("Empty", func(t *testing.T) {
		h := NewMaxHeap([]int{})
		require.True(t, h.IsValidHeap())

		h = NewMaxHeap[int](nil)
		require.True(t, h.IsValidHeap())
	})
}

func TestMinHeap(t *testing.T) {
	// 随机生成序列，创建最小堆
	t.Run("Random", func(t *testing.T) {
		const testcnt = 100
		for i := range testcnt {
			list := GenList(1000, 10000)

			start := time.Now()
			h := NewMinHeap(list)
			require.True(t, h.IsValidHeap())
			require.Equal(t, h.data[0], lo.Min(list))
			fmt.Println("[INFO] Success testcase", i, " in", time.Since(start).Milliseconds(), "ms")
		}
	})

	// 序列为空
	t.Run("Empty", func(t *testing.T) {
		h := NewMinHeap([]int{})
		require.True(t, h.IsValidHeap())
		h = NewMinHeap[int](nil)
		require.True(t, h.IsValidHeap())
	})
}

func TestIsValidHeap(t *testing.T) {
	// 最小堆
	t.Run("MinHeap", func(t *testing.T) {
		h := &Heap[int]{
			data:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			htype: minHeap,
		}
		require.True(t, h.IsValidHeap())
	})

	// 最大堆
	t.Run("MaxHeap", func(t *testing.T) {
		h := &Heap[int]{
			data:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			htype: maxHeap,
		}
		require.True(t, h.IsValidHeap())
	})

	// 不是最小堆
	t.Run("NotMinHeap", func(t *testing.T) {
		h := &Heap[int]{
			data:  []int{1, 5, 3, 4, 2, 6, 7, 8, 9, 10},
			htype: minHeap,
		}
		require.False(t, h.IsValidHeap())

		h.data = []int{1, 2, 6, 4, 5, 7, 3, 8, 9, 10, 11}
		require.False(t, h.IsValidHeap())
	})

	// 不是最大堆
	t.Run("NotMaxHeap", func(t *testing.T) {
		h := &Heap[int]{
			data:  []int{10, 9, 3, 7, 6, 5, 4, 8, 2, 1},
			htype: maxHeap,
		}
		require.False(t, h.IsValidHeap())

		h.data = []int{10, 9, 3, 7, 6, 2, 4, 8, 5, 1, 0}
		require.False(t, h.IsValidHeap())
	})

	// 空序列
	t.Run("Empty", func(t *testing.T) {
		h := &Heap[int]{
			data:  []int{},
			htype: minHeap,
		}
		require.True(t, h.IsValidHeap())
		h.data = nil
		require.True(t, h.IsValidHeap())
	})
}

func TestPopPush(t *testing.T) {
	// 最大堆
	t.Run("MaxHeap", func(t *testing.T) {
		h := &Heap[int]{
			data:  nil,
			htype: maxHeap,
		}
		list := GenList(1000, 10000)
		for idx := range list {
			h.Push(list[idx])
		}
		require.True(t, h.IsValidHeap())

		res := sort.CountingSort(list, 10000, sort.OrderDesc)
		for idx := range list {
			val, ok := h.Pop()
			require.True(t, ok)
			require.True(t, h.IsValidHeap())
			require.Equal(t, val, res[idx])
		}
	})

	// 最小堆
	t.Run("MinHeap", func(t *testing.T) {
		h := &Heap[int]{
			data:  nil,
			htype: minHeap,
		}
		list := GenList(1000, 10000)
		for idx := range list {
			h.Push(list[idx])
		}
		require.True(t, h.IsValidHeap())
		res := sort.CountingSort(list, 10000, sort.OrderAsc)
		for idx := range list {
			val, ok := h.Pop()
			require.True(t, ok)
			require.True(t, h.IsValidHeap())
			require.Equal(t, val, res[idx])
		}
	})

	// 空序列
	t.Run("Empty", func(t *testing.T) {
		h := &Heap[int]{
			data:  nil,
			htype: minHeap,
		}
		val, ok := h.Pop()
		require.False(t, ok)
		require.Equal(t, val, 0)
	})
}

func TestUpdate(t *testing.T) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 最大堆
	t.Run("MaxHeap", func(t *testing.T) {
		list := GenList(1000, 10000)
		h := NewMaxHeap(list)
		require.True(t, h.IsValidHeap())

		// 随机更新
		for idx := range list {
			old := list[idx]
			list[idx] = rd.Intn(10000)
			require.True(t, h.Update(old, list[idx]))
			require.True(t, h.IsValidHeap())
			require.Equal(t, h.data[0], lo.Max(list))
		}
	})

	// 最小堆
	t.Run("MinHeap", func(t *testing.T) {
		list := GenList(1000, 10000)
		h := NewMinHeap(list)
		require.True(t, h.IsValidHeap())

		// 随机更新
		for idx := range list {
			old := list[idx]
			list[idx] = rd.Intn(10000)
			require.True(t, h.Update(old, list[idx]))
			require.True(t, h.IsValidHeap())
			require.Equal(t, h.data[0], lo.Min(list))
		}
	})

	// 更新的值不存在或没有变化
	t.Run("NotExist", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		h := NewMaxHeap(list)
		require.True(t, h.IsValidHeap())
		require.False(t, h.Update(11, 12))
		require.False(t, h.Update(0, 1))
		require.True(t, h.Update(10, 10))
	})
}
