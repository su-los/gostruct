package segtree

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewSegTree(t *testing.T) {
	st := NewSegTree([]int{1, 2}, Sum)
	require.NotNil(t, st)
	require.Equal(t, 3, st.len)

	st = NewSegTree[int](nil, Sum)
	require.NotNil(t, st)
	require.Equal(t, 0, st.len)
	require.Nil(t, st.root)

	st = NewSegTree([]int{}, Sum)
	require.NotNil(t, st)
	require.Equal(t, 0, st.len)
	require.Nil(t, st.root)
}

func TestPostOrder(t *testing.T) {
	seg := NewSegTree([]int{}, Sum)
	require.Nil(t, seg.PostOrder())
}

func TestPreOrder(t *testing.T) {
	seg := NewSegTree([]int{}, Sum)
	require.Nil(t, seg.PreOrder())
}

func TestLevelOrder(t *testing.T) {
	seg := NewSegTree([]int{}, Sum)
	require.Nil(t, seg.LevelOrder())
}

func TestInOrder(t *testing.T) {
	seg := NewSegTree([]int{}, Sum)
	require.Nil(t, seg.InOrder())
}

func TestQuery(t *testing.T) {
	testcnt := 1100
	rd := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	fArr := []AggFunc[int]{Sum[int], Max[int], Min[int]}
	for i := range testcnt {
		cnt := min((i+1)*10, 10000)
		nums := GenNumList(cnt, 100000)

		fidx := i % 3
		f := fArr[fidx]
		st := NewSegTree(nums, f)

		// 随机生成查询的区间
		start := rd.IntN(cnt)
		end := rd.IntN(cnt - start)
		end = min(end+start, cnt)

		begin := time.Now()
		// 计算聚合值，与查询的值进行对比
		ans := f(nums[start:end])
		res, err := st.Query(start, end)

		if start < end {
			require.NoError(t, err)
		} else {
			require.ErrorIs(t, ErrNotInRang, err)
		}
		require.Equal(t, ans, res)
		fmt.Println("[INFO] test case:", i, "success, cost=", time.Since(begin).Milliseconds(), "ms, cnt=", cnt)
	}
}
