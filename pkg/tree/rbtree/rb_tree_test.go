package rbtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFixInsertionCase1(t *testing.T) {
	// 情况 1: 父节点为黑色
	//      4(B)
	//     /   \
	//  3(R)    5(R)
	var (
		n4 = newRBNode(4, black)
		n3 = newRBNode(3, red)
		n5 = newRBNode(5, red)
	)

	n4.left, n4.right = n3, n5
	n3.parent, n5.parent = n4, n4
	rb := &RBTree[int]{root: n4}
	require.NoError(t, rb.fixInsertion(n3))

	require.Equal(t, n4, rb.root)
	require.Equal(t, 4, n4.val)
	require.Equal(t, black, n4.color)
	require.Equal(t, n3, n4.left)
	require.Equal(t, n5, n4.right)

	require.Equal(t, n4, n3.parent)
	require.Equal(t, 3, n3.val)
	require.Equal(t, red, n3.color)
	require.Nil(t, n3.left)
	require.Nil(t, n3.right)

	require.Equal(t, n4, n5.parent)
	require.Equal(t, 5, n5.val)
	require.Equal(t, red, n5.color)
	require.Nil(t, n5.left)
	require.Nil(t, n5.right)
}

func TestFixInsertionCase2(t *testing.T) {
	// 情况 2: 父节点为红色，叔父节点为红色
	//       10(B)
	//        /
	//      6(B)
	//     /   \
	//  5(R)   9(R)
	//   /
	// 3(R)
	var (
		n10 = newRBNode(10, black)
		n6  = newRBNode(6, black)
		n5  = newRBNode(5, red)
		n9  = newRBNode(9, red)
		n3  = newRBNode(3, red)
	)

	n10.left = n6
	n6.parent = n10
	n6.left, n6.right = n5, n9
	n5.parent, n9.parent = n6, n6
	n5.left = n3
	n3.parent = n5

	rb := &RBTree[int]{root: n10}
	require.NoError(t, rb.fixInsertion(n3))
	require.Equal(t, n10, rb.root)

	require.Equal(t, 10, n10.val)
	require.Equal(t, black, n10.color)
	require.Equal(t, n6, n10.left)
	require.Nil(t, n10.right)

	require.Equal(t, 6, n6.val)
	require.Equal(t, red, n6.color)
	require.Equal(t, n5, n6.left)
	require.Equal(t, n9, n6.right)
	require.Equal(t, n10, n6.parent)

	require.Equal(t, 5, n5.val)
	require.Equal(t, black, n5.color)
	require.Equal(t, n3, n5.left)
	require.Nil(t, n5.right)
	require.Equal(t, n6, n5.parent)

	require.Equal(t, 9, n9.val)
	require.Equal(t, black, n9.color)
	require.Nil(t, n9.left)
	require.Nil(t, n9.right)
	require.Equal(t, n6, n9.parent)

	require.Equal(t, 3, n3.val)
	require.Equal(t, red, n3.color)
	require.Nil(t, n3.left)
	require.Nil(t, n3.right)
	require.Equal(t, n5, n3.parent)
}

func TestFixInsertionCase3LL(t *testing.T) {
	// 情况 3: 父节点是红色，叔父节点是黑色；插入节点 LL
	//        10(B)                   10(B)
	//       /                        /
	//      7(B)                     5(B)
	//     /   \                   /    \
	//  5(R)   9(B)    ------>   3(R)   7(R)
	// /   \                            /   \
	//3(R) 6(B)                       6(B)  9(B)
	var (
		n10 = newRBNode(10, black)
		n7  = newRBNode(7, black)
		n5  = newRBNode(5, red)
		n9  = newRBNode(9, black)
		n3  = newRBNode(3, red)
		n6  = newRBNode(6, black)
	)
	n10.left = n7
	n7.parent = n10
	n7.left, n7.right = n5, n9
	n5.parent, n9.parent = n7, n7
	n5.left, n5.right = n3, n6
	n3.parent, n6.parent = n5, n5

	rb := &RBTree[int]{root: n10}
	require.NoError(t, rb.fixInsertion(n3))
	require.Equal(t, n5, n10.left)
	// 校验节点 n5
	require.Equal(t, 5, n5.val)
	require.Equal(t, black, n5.color)
	require.Equal(t, n10, n5.parent)
	require.Equal(t, n3, n5.left)
	require.Equal(t, n7, n5.right)
	// 校验节点 n7
	require.Equal(t, 7, n7.val)
	require.Equal(t, red, n7.color)
	require.Equal(t, n5, n7.parent)
	require.Equal(t, n6, n7.left)
	require.Equal(t, n9, n7.right)
	// 校验节点 n3
	require.Equal(t, 3, n3.val)
	require.Equal(t, red, n3.color)
	require.Equal(t, n5, n3.parent)
	require.Nil(t, n3.left)
	require.Nil(t, n3.right)
	// 校验节点 n9
	require.Equal(t, 9, n9.val)
	require.Equal(t, black, n9.color)
	require.Equal(t, n7, n9.parent)
	require.Nil(t, n9.left)
	require.Nil(t, n9.right)
	// 校验节点 n6
	require.Equal(t, 6, n6.val)
	require.Equal(t, black, n6.color)
	require.Equal(t, n7, n6.parent)
	require.Nil(t, n6.left)
	require.Nil(t, n6.right)
}

func TestFixInsertion(t *testing.T) {
	// 情况 4: 父节点是红色，叔父节点是黑色；插入节点 LR
	t.Run("PUBlackLR", func(t *testing.T) {
		//      7(B)                7(B)               6(B)
		//   5(R)  9(B) ---->    6(R)  9(B)         5(R)   7(R)
		// 3(B) 6(R)          5(R)                3(B)        9(B)
		//                 3(B)
		var (
			n7 = newRBNode(7, black)
			n5 = newRBNode(5, red)
			n9 = newRBNode(9, black)
			n6 = newRBNode(6, red)
			n3 = newRBNode(3, black)
		)
		n7.left, n7.right = n5, n9
		n5.parent, n9.parent = n7, n7
		n5.right, n5.left = n6, n3
		n6.parent, n3.parent = n5, n5
		rb := &RBTree[int]{root: n7}
		require.NoError(t, rb.fixInsertion(n6))
		// 校验 n6
		require.Equal(t, 6, n6.val)
		require.Equal(t, black, n6.color)
		require.Nil(t, n6.parent)
		require.Equal(t, n5, n6.left)
		require.Equal(t, n7, n6.right)
		// 校验节点 n5
		require.Equal(t, 5, n5.val)
		require.Equal(t, red, n5.color)
		require.Equal(t, n6, n5.parent)
		require.Equal(t, n3, n5.left)
		require.Nil(t, n5.right)
		// 校验节点 n7
		require.Equal(t, 7, n7.val)
		require.Equal(t, red, n7.color)
		require.Equal(t, n6, n7.parent)
		require.Nil(t, n7.left)
		require.Equal(t, n9, n7.right)
		// 校验节点 n3
		require.Equal(t, 3, n3.val)
		require.Equal(t, black, n3.color)
		require.Equal(t, n5, n3.parent)
		require.Nil(t, n3.left)
		require.Nil(t, n3.right)
		// 校验节点 n9
		require.Equal(t, 9, n9.val)
		require.Equal(t, black, n9.color)
		require.Equal(t, n7, n9.parent)
		require.Nil(t, n9.left)
		require.Nil(t, n9.right)
	})
	// 情况 4: 父节点是红色，叔父节点是黑色；插入节点 RR
	t.Run("PUBlackRR", func(t *testing.T) {
		//      7(B)                9(B)
		//   5(B)  9(R) ---->    7(R)  10(R)
		//      8(B) 10(R)     5(B) 8(B)
		var (
			n7  = newRBNode(7, black)
			n9  = newRBNode(9, red)
			n5  = newRBNode(5, black)
			n10 = newRBNode(10, red)
			n8  = newRBNode(8, black)
		)
		n7.left, n7.right = n5, n9
		n5.parent, n9.parent = n7, n7
		n9.right, n9.left = n10, n8
		n8.parent, n10.parent = n9, n9
		rb := &RBTree[int]{root: n7}
		require.NoError(t, rb.fixInsertion(n10))
		// 校验节点 n9
		require.Equal(t, 9, n9.val)
		require.Equal(t, black, n9.color)
		require.Nil(t, n9.parent)
		require.Equal(t, n7, n9.left)
		require.Equal(t, n10, n9.right)
		// 校验节点 n7
		require.Equal(t, 7, n7.val)
		require.Equal(t, red, n7.color)
		require.Equal(t, n9, n7.parent)
		require.Equal(t, n5, n7.left)
		require.Equal(t, n8, n7.right)
		// 校验节点 n10
		require.Equal(t, 10, n10.val)
		require.Equal(t, red, n10.color)
		require.Equal(t, n9, n10.parent)
		require.Nil(t, n10.left)
		require.Nil(t, n10.right)
		// 校验节点 n5
		require.Equal(t, 5, n5.val)
		require.Equal(t, black, n5.color)
		require.Equal(t, n7, n5.parent)
		require.Nil(t, n5.left)
		require.Nil(t, n5.right)
		// 校验节点 n8
		require.Equal(t, 8, n8.val)
		require.Equal(t, black, n8.color)
		require.Equal(t, n7, n8.parent)
		require.Nil(t, n8.left)
		require.Nil(t, n8.right)
	})
	// 情况 4: 父节点是红色，叔父节点是黑色；插入节点 RL
	t.Run("PUBlackLR", func(t *testing.T) {
		//      7(B)                7(B)               8(B)
		//   5(B)  9(R) ---->    5(B)  8(R)         7(R)   9(R)
		//      8(R) 10(B)               9(R)     5(B)       10(B)
		//                                  10(B)
		var (
			n7  = newRBNode(7, black)
			n9  = newRBNode(9, red)
			n5  = newRBNode(5, black)
			n8  = newRBNode(8, red)
			n10 = newRBNode(10, black)
		)
		n7.left, n7.right = n5, n9
		n5.parent, n9.parent = n7, n7
		n9.right, n9.left = n10, n8
		n8.parent, n10.parent = n9, n9
		rb := &RBTree[int]{root: n7}
		require.NoError(t, rb.fixInsertion(n8))
		// 校验节点 n8
		require.Equal(t, 8, n8.val)
		require.Equal(t, black, n8.color)
		require.Nil(t, n8.parent)
		require.Equal(t, n7, n8.left)
		require.Equal(t, n9, n8.right)
		// 校验节点 n7
		require.Equal(t, 7, n7.val)
		require.Equal(t, red, n7.color)
		require.Equal(t, n8, n7.parent)
		require.Equal(t, n5, n7.left)
		require.Nil(t, n7.right)
		// 校验节点 n9
		require.Equal(t, 9, n9.val)
		require.Equal(t, red, n9.color)
		require.Equal(t, n8, n9.parent)
		require.Nil(t, n9.left)
		require.Equal(t, n10, n9.right)
		// 校验节点 n10
		require.Equal(t, 10, n10.val)
		require.Equal(t, black, n10.color)
		require.Equal(t, n9, n10.parent)
		require.Nil(t, n10.left)
		require.Nil(t, n10.right)
		// 校验节点 n5
		require.Equal(t, 5, n5.val)
		require.Equal(t, black, n5.color)
		require.Equal(t, n7, n5.parent)
		require.Nil(t, n5.left)
		require.Nil(t, n5.right)
	})
}

func TestInsert(t *testing.T) {
	const count = 1000
	for range count {
		data := GenBFSList()
		rb := NewRBTree[int]()
		for idx := range data {
			require.NoError(t, rb.Insert(data[idx]))
		}
		_, ok := rb.root.verifyBlackHeightAndRedRules()
		require.True(t, ok)
		require.True(t, rb.IsValid())
	}
}

func TestDeletionCase2_BroRight(t *testing.T) {
	// 处理情况2：兄弟节点为红色
	//
	//  GP(B,2)                    GP(B,2)
	//      \                          \
	//	   P(B,6)                     S(B,10)
	//	   /  \                       /     \
	//	N(B,4) S(R,10)   ---->     P(B,6)   SR(B,12)
	//	       / \                  /  \         /  \
	//	  SL(B,9) SR(B,12)      N(B,4) SL(R,9)  SRL(B) SRR(B)
	//     / \      /  \               /  \
	// SLL(B)SLR(B)SRL(B) SRR(B)   SLL(B) SLR(B)
	var (
		// gp2 仅用来验证旋转后的父节点指针的正确性能；子树 P 用来验证平衡性
		gp2                = newRBNode(2, black)
		p6                 = newRBNode(6, black)
		n4                 = newRBNode(4, black)
		s10                = newRBNode(10, red)
		sl9                = newRBNode(9, black)
		sr12               = newRBNode(12, black)
		sll, slr, srl, srr = newRBNode(0, black), newRBNode(0, black), newRBNode(0, black), newRBNode(0, black)
	)
	gp2.right = p6
	p6.parent = gp2
	p6.left, p6.right = n4, s10
	n4.parent, s10.parent = p6, p6
	s10.left, s10.right = sl9, sr12
	sl9.parent, sr12.parent = s10, s10
	sl9.left, sl9.right = sll, slr
	sll.parent, slr.parent = sl9, sl9
	sr12.left, sr12.right = srl, srr
	srl.parent, srr.parent = sr12, sr12

	rb := &RBTree[int]{
		root: gp2,
	}
	require.NoError(t, rb.fixDeletion(n4, p6))

	// 验证
	require.Equal(t, gp2, rb.root)
	require.Equal(t, black, s10.color)
	require.Equal(t, gp2, s10.parent)
	require.Equal(t, p6, s10.left)
	require.Equal(t, sr12, s10.right)

	require.Equal(t, black, p6.color)
	require.Equal(t, s10, p6.parent)
	require.Equal(t, n4, p6.left)
	require.Equal(t, sl9, p6.right)

	require.Equal(t, black, sr12.color)
	require.Equal(t, s10, sr12.parent)
	require.Equal(t, srl, sr12.left)
	require.Equal(t, srr, sr12.right)

	require.Equal(t, black, n4.color)
	require.Equal(t, p6, n4.parent)
	require.Nil(t, n4.left)
	require.Nil(t, n4.right)

	require.Equal(t, red, sl9.color)
	require.Equal(t, p6, sl9.parent)
	require.Equal(t, sll, sl9.left)
	require.Equal(t, slr, sl9.right)
}

func TestDeletionCase2_BroLeft(t *testing.T) {
	// 处理情况2：兄弟节点为红色
	//
	//     GP(B,2)                    GP(B,2)
	//        \                           \
	//	     P(B,6)                       S(B,4)
	//	      /  \        变色+旋转        /     \
	//	   S(R,4) N(B,9)   ---->     SL(B,3)    P(B,6)
	//	     / \                    / \          /  \
	// SL(B,3) SR(B,5)          SLR(B) SLL(B)SR(R,5) N(B,9)
	// / \     /  \                          /  \
	// SLL(B)SLR(B)SRL(B) SRR(B)          SRL(B) SRR(B)
	var (
		// gp2 仅用来验证旋转后的父节点指针的正确性能；子树 P 用来验证平衡性
		gp2                = newRBNode(2, black)
		p6                 = newRBNode(6, black)
		n9                 = newRBNode(9, black)
		s4                 = newRBNode(4, red)
		sl3                = newRBNode(3, black)
		sr5                = newRBNode(5, black)
		sll, slr, srl, srr = newRBNode(0, black), newRBNode(0, black), newRBNode(0, black), newRBNode(0, black)
	)
	gp2.right = p6
	p6.parent = gp2
	p6.left, p6.right = s4, n9
	s4.parent, n9.parent = p6, p6
	s4.left, s4.right = sl3, sr5
	sl3.parent, sr5.parent = s4, s4
	sl3.left, sl3.right = sll, slr
	sll.parent, slr.parent = sl3, sl3
	sr5.left, sr5.right = srl, srr
	srl.parent, srr.parent = sr5, sr5

	rb := &RBTree[int]{root: gp2}
	require.NoError(t, rb.fixDeletion(n9, p6))

	// 验证
	require.Equal(t, gp2, rb.root)
	require.Equal(t, black, s4.color)
	require.Equal(t, gp2, s4.parent)
	require.Equal(t, sl3, s4.left)
	require.Equal(t, p6, s4.right)

	require.Equal(t, black, p6.color)
	require.Equal(t, s4, p6.parent)
	require.Equal(t, sr5, p6.left)
	require.Equal(t, n9, p6.right)

	require.Equal(t, black, sl3.color)
	require.Equal(t, s4, sl3.parent)
	require.Equal(t, sll, sl3.left)
	require.Equal(t, slr, sl3.right)

	require.Equal(t, black, n9.color)
	require.Equal(t, p6, n9.parent)
	require.Nil(t, n9.left)
	require.Nil(t, n9.right)

	require.Equal(t, red, sr5.color)
	require.Equal(t, p6, sr5.parent)
	require.Equal(t, srl, sr5.left)
	require.Equal(t, srr, sr5.right)
}

func TestDeletionCase3_1(t *testing.T) {
	// 处理情况 3.1：兄弟节点为黑色，P、SL、SR也为黑色
	//
	//      GP(B,21)                    GP(B,21)
	//        /                           /
	//	   P(B,6)                      P(B,6)
	//	   /  \      变色+向上调整        /  \
	//	N(B,4) S(B,10)   ---->     N(B,4) S(R,10)
	//	       / \                        /  \
	//	  SL(B,9) SR(B,12)            SL(B,9) SR(B,12)
	var (
		gp21 = newRBNode(21, black)
		p6   = newRBNode(6, black)
		n4   = newRBNode(4, black)
		s10  = newRBNode(10, black)
		sl9  = newRBNode(9, black)
		sr12 = newRBNode(12, black)
	)
	gp21.left = p6
	p6.parent = gp21
	p6.left, p6.right = n4, s10
	n4.parent, s10.parent = p6, p6
	s10.left, s10.right = sl9, sr12
	sl9.parent, sr12.parent = s10, s10
	newn := fixDeletionCase3_1(n4, p6)

	// 验证
	require.Equal(t, p6, newn)
	require.Equal(t, black, p6.color)
	require.Equal(t, gp21, p6.parent)
	require.Equal(t, n4, p6.left)
	require.Equal(t, s10, p6.right)

	require.Equal(t, red, s10.color)
	require.Equal(t, p6, s10.parent)
	require.Equal(t, sl9, s10.left)
	require.Equal(t, sr12, s10.right)

	require.Equal(t, black, sr12.color)
	require.Equal(t, s10, sr12.parent)
	require.Nil(t, sr12.left)
	require.Nil(t, sr12.right)

	require.Equal(t, black, n4.color)
	require.Equal(t, p6, n4.parent)
	require.Nil(t, n4.left)
	require.Nil(t, n4.right)

	require.Equal(t, black, sl9.color)
	require.Equal(t, s10, sl9.parent)
	require.Nil(t, sl9.left)
	require.Nil(t, sl9.right)
}

func TestDelete(t *testing.T) {
	const count = 1000
	rd := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	for range count {
		data := GenBFSList()
		rb := NewRBTree[int]()
		for idx := range data {
			require.NoError(t, rb.Insert(data[idx]))
		}

		delCnt := len(data) / 10
		for range delCnt {
			i := rd.Intn(delCnt)
			require.NoError(t, rb.Delete(data[i]))
		}

		_, ok := rb.root.verifyBlackHeightAndRedRules()
		require.True(t, ok)
		require.True(t, rb.IsValid())
	}
}
