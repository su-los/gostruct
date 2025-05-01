package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/su-los/gostruct/pkg/sort"
)

func TestNewDualHeap(t *testing.T) {
	// k <= 0
	t.Run("KInvalid", func(t *testing.T) {
		_, err := NewDualHeap[int](0)
		require.Error(t, err)
	})

	// k > 0
	t.Run("KValid", func(t *testing.T) {
		dh, err := NewDualHeap[int](1)
		require.NoError(t, err)
		require.NotNil(t, dh)
	})
}

func TestPushPop(t *testing.T) {
	k := 100
	dh, err := NewDualHeap[int](k)
	require.NoError(t, err)

	list := GenList(1000, 10000)
	for idx := range list {
		dh.Push(list[idx])
	}

	sorted := sort.CountingSort(list, 10000, sort.OrderDesc)
	var (
		topk, popval int
		ok           bool
	)
	for idx := k - 1; idx < len(sorted); idx++ {
		topk, ok = dh.TopK()
		require.True(t, ok)
		popval, ok = dh.Pop()
		require.True(t, ok)

		require.Equal(t, topk, popval)
		require.Equal(t, sorted[idx], topk)
	}
	require.Len(t, dh.maxHeap.data, 0)

	for idx := k - 2; idx >= 0; idx-- {
		topk, ok = dh.TopK()
		require.True(t, ok)
		popval, ok = dh.Pop()
		require.True(t, ok)

		require.Equal(t, topk, popval)
		require.Equal(t, sorted[idx], popval)
	}
	require.Len(t, dh.maxHeap.data, 0)
	require.Len(t, dh.minHeap.data, 0)

	_, ok = dh.TopK()
	require.False(t, ok)
	_, ok = dh.Pop()
	require.False(t, ok)
}

func TestUpdateK(t *testing.T) {
	k := 100
	dh, err := NewDualHeap[int](k)
	require.NoError(t, err)

	list := GenList(1000, 10000)
	for idx := range list {
		dh.Push(list[idx])
	}

	sorted := sort.CountingSort(list, 10000, sort.OrderDesc)
	require.NoError(t, dh.UpdateK(k+100))
	require.Equal(t, sorted[k-1+100], dh.minHeap.data[0])

	require.NoError(t, dh.UpdateK(k-50))
	require.Equal(t, sorted[k-1-50], dh.minHeap.data[0])
	require.Error(t, dh.UpdateK(k-100))
}
