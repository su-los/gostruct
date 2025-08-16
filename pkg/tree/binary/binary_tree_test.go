package binary

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFromLevelOrder(t *testing.T) {
	vals := []int{8, 17, 21, 18, 0, 0, 6}
	tree := FromLevelOrder(vals)

	require.Equal(t, 8, tree.root.Val)
	require.Equal(t, 17, tree.root.Left.Val)
	require.Equal(t, 21, tree.root.Right.Val)

	l := tree.root.Left
	require.Equal(t, 18, l.Left.Val)
	require.Nil(t, l.Right)

	r := tree.root.Right
	require.Nil(t, r.Left)
	require.Equal(t, 6, r.Right.Val)
}

func TestLevelOrder(t *testing.T) {
	const testcnt = 100
	for i := range testcnt {
		data := GenBFSList(1000, 1000)
		tree := FromLevelOrder(data)

		start := time.Now()
		res := tree.LevelOrder()
		resRe := tree.LevelOrderRecursive()
		require.ElementsMatch(t, res, resRe)
		fmt.Println("[INFO] Success testcase", i, "(len=", len(res), ") in", time.Since(start).Milliseconds(), "ms")
	}
}
