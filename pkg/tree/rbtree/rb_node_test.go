package rbtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyBlackHeightAndRedRules(t *testing.T) {
	//          2(B)
	//        /    \
	//     1(R)    3(R)
	//    /          \
	// 0(B)           4(R)

	var (
		n2 = newRBNode(2, black)
		n1 = newRBNode(1, red)
		n3 = newRBNode(3, red)
		n0 = newRBNode(0, black)
		n4 = newRBNode(4, red)
	)

	n2.left, n2.right = n1, n3
	n1.parent, n3.parent = n2, n2
	n3.right = n4
	n4.parent = n3
	n1.left = n0
	n0.parent = n1

	_, ok := n2.verifyBlackHeightAndRedRules()
	require.False(t, ok)

	_, ok = n1.verifyBlackHeightAndRedRules()
	require.False(t, ok)
	_, ok = n3.verifyBlackHeightAndRedRules()
	require.False(t, ok)

	var h int
	h, ok = n0.verifyBlackHeightAndRedRules()
	require.True(t, ok)
	require.Equal(t, 2, h)
	_, ok = n4.verifyBlackHeightAndRedRules()
	require.False(t, ok)
}

func TestGetGrandparent(t *testing.T) {
	node := newRBNode(1, red)
	node.right = newRBNode(2, red)
	node.right.parent = node
	node.right.right = newRBNode(3, red)
	node.right.right.parent = node.right
	require.Nil(t, node.getGrandparent())
	require.Nil(t, node.right.getGrandparent())
	require.NotNil(t, node.right.right.getGrandparent())
	require.Equal(t, node, node.right.right.getGrandparent())
}

func TestGetUncle(t *testing.T) {
	node := newRBNode(3, red)
	node.right = newRBNode(4, red)
	node.right.parent = node
	node.right.right = newRBNode(5, red)
	node.right.right.parent = node.right
	node.left = newRBNode(2, red)
	node.left.parent = node
	node.left.left = newRBNode(1, red)
	node.left.left.parent = node.left
	require.Nil(t, node.getUncle())
	require.Nil(t, node.right.getUncle())
	require.NotNil(t, node.right.right.getUncle())
	require.Equal(t, node.left, node.right.right.getUncle())
	require.NotNil(t, node.left.left.getUncle())
	require.Equal(t, node.right, node.left.left.getUncle())
}
