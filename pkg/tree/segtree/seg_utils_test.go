package segtree

import (
	"cmp"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// checkArrayEqual 检查一维数组是否相等
func checkArrayEqual[T cmp.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// check2DArrayEqual 检查二维数组是否相等
func check2DArrayEqual[T cmp.Ordered](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !checkArrayEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestBuild(t *testing.T) {
	testcnt := 110
	for i := range testcnt {
		nums := GenNumList(1000, 10000)
		seg, cnt := build(nums, Min)
		segRe, cntRe := buildRecursive(nums, 0, len(nums), Min)
		require.Equal(t, cnt, cntRe)

		st, stRe := &SegTree[int]{
			root: seg,
			len:  cnt,
		}, &SegTree[int]{
			root: segRe,
			len:  cntRe,
		}
		levelOrder := st.LevelOrder()
		levelOrderRe := stRe.LevelOrder()
		require.True(t, check2DArrayEqual(levelOrder, levelOrderRe))
		require.Len(t, levelOrder, cnt)

		preOrder := st.PreOrder()
		preOrderRe := stRe.PreOrder()
		require.True(t, check2DArrayEqual(preOrder, preOrderRe))
		require.Len(t, preOrder, cnt)

		inOrder := st.InOrder()
		inOrderRe := stRe.InOrder()
		require.True(t, check2DArrayEqual(inOrder, inOrderRe))
		require.Len(t, inOrder, cnt)

		postOrder := st.PostOrder()
		postOrderRe := stRe.PostOrder()
		require.True(t, check2DArrayEqual(postOrder, postOrderRe))
		require.Len(t, postOrder, cnt)
		fmt.Println("[INFO] case ", i, " success", ", len: ", cnt)
	}
}

func TestBuild_Special(t *testing.T) {
	seg, cnt := build([]int{}, Sum)
	require.Nil(t, seg)
	require.Equal(t, 0, cnt)

	seg, cnt = build([]int{1}, Sum)
	require.NotNil(t, seg)
	require.Equal(t, 1, cnt)
	require.Equal(t, 1, len(seg.seg))
}

func TestBuildRecursive_Special(t *testing.T) {
	seg, cnt := buildRecursive([]int{}, 0, 0, Sum)
	require.Nil(t, seg)
	require.Equal(t, 0, cnt)

	seg, cnt = buildRecursive([]int{1}, 0, 1, Sum)
	require.NotNil(t, seg)
	require.Equal(t, 1, cnt)
	require.Equal(t, 1, len(seg.seg))
}

func TestBuildBottomUp(t *testing.T) {
	testcnt := 110
	for i := range testcnt {
		nums := GenNumList(1000, 10000)
		segBtUp, cntBtUp := buildBottomUp(nums, Max)
		_, cntRe := buildRecursive(nums, 0, len(nums), Max)
		require.NotNil(t, segBtUp)
		require.Equal(t, cntBtUp, cntRe)

		st := &SegTree[int]{
			root: segBtUp,
			len:  cntBtUp,
		}
		levelOrder := st.LevelOrder()
		require.Len(t, levelOrder, cntBtUp)
		fmt.Println("[INFO] case ", i, " success", ", len: ", cntBtUp)
	}
}
