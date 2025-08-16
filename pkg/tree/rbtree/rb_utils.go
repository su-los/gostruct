package rbtree

import "cmp"

// FixFunc 修复红黑树的函数定义
type FixFunc[T cmp.Ordered] func(*rbNode[T]) *rbNode[T]

// rotateRight 右旋
//
// 返回旋转后新的根节点(需要将原先的父节点的孩子设置成新返回的节点)
func rotateRight[T cmp.Ordered](root *rbNode[T]) *rbNode[T] {
	newRoot := root.left
	root.left = newRoot.right
	if newRoot.right != nil {
		newRoot.right.parent = root
	}
	newRoot.parent = root.parent
	newRoot.right = root
	root.parent = newRoot
	return newRoot
}

// rotateLeft 左旋
//
// 返回旋转后新的根节点(需要将原先的父节点的孩子设置成新返回的节点)
func rotateLeft[T cmp.Ordered](root *rbNode[T]) *rbNode[T] {
	newRoot := root.right
	root.right = newRoot.left
	if newRoot.left != nil {
		newRoot.left.parent = root
	}
	newRoot.parent = root.parent
	newRoot.left = root
	root.parent = newRoot
	return newRoot
}

// isBlack 是否是黑色
func isBlack[T cmp.Ordered](node *rbNode[T]) bool {
	if node == nil {
		return true
	}
	return node.color == black
}

// fixInsertionCase3LL 处理情况3：父节点为红色，叔父节点为黑色；插入节点是 LL 型
//
//	     G(B)                        P(B)
//	    /    \      右旋             /    \
//	 P(R)    U(B) --------->      z(R)    G(R)
//	 /                                        \
//	z(R)                                      U(B)
//
// 1. 将父节点变成黑色，祖父节点变成红色
// 2. 将祖父节点右旋
// 返回旋转后新的根节点
func fixInsertionCase3LL[T cmp.Ordered](node *rbNode[T]) *rbNode[T] {
	// 情况 3 一定会有祖父节点，因为父节点是红色不可能作为根节点
	var (
		parent = node.parent
		grand  = parent.parent
	)
	parent.color = black
	grand.color = red
	return rotateRight(grand)
}

// fixInsertionCase3RR 处理情况3：父节点为红色，叔父节点为黑色；插入节点是 RR 型
//
//	    G(B)                        P(B)
//	   /    \      左旋             /    \
//	U(B)    P(R) --------->      G(R)    z(R)
//	          \                   /
//	          z(R)             U(B)
//
// 1. 将父节点变成黑色，祖父节点变成红色
// 2. 将祖父节点左旋
// 返回旋转后新的根节点
func fixInsertionCase3RR[T cmp.Ordered](node *rbNode[T]) *rbNode[T] {
	// 情况 3 一定会有祖父节点，因为父节点是红色不可能作为根节点
	var (
		parent = node.parent
		grand  = node.getGrandparent()
	)
	parent.color = black
	grand.color = red
	return rotateLeft(grand)
}

// fixInsertionCase4LR 处理情况4: 父节点为红色，叔父节点为黑色；插入节点是 LR 型
//
//	    G(B)                        G(B)
//	   /    \    左旋转换成LL        /    \
//	P(R)    U(B) --------->      z(R)    U(B)
//	  \                          /
//	  z(R)                     P(R)
//
// 1. 对父节点左旋
// 2. 处理节点变成父节点 -> 转换成情况 3 LL
// 返回旋转后的新根节点
func fixInsertionCase4LR[T cmp.Ordered](node *rbNode[T]) *rbNode[T] {
	var (
		parent = node.parent
		grand  = parent.parent
	)
	if grand.left == parent {
		grand.left = rotateLeft(parent)
	} else {
		grand.right = rotateLeft(parent)
	}
	return fixInsertionCase3LL(parent)
}

// fixInsertionCase4RL 处理情况4: 父节点为红色，叔父节点为黑色；插入节点是 RL 型
//
// 1. 对父节点右旋
// 2. 处理节点变成父节点 -> 转换成情况 3 RR
// 返回旋转后的新根节点
func fixInsertionCase4RL[T cmp.Ordered](node *rbNode[T]) *rbNode[T] {
	var (
		parent = node.parent
		grand  = parent.parent
	)
	if grand.left == parent {
		grand.left = rotateRight(parent)
	} else {
		grand.right = rotateRight(parent)
	}
	return fixInsertionCase3RR(parent)
}

// --------------------------------------------------------
// 删除后的平衡
// --------------------------------------------------------
// getBrother 获取 n 的兄弟节点（n 可能为空）
//
// 返回兄弟节点，以及兄弟节点在左还是右
func getBrother[T cmp.Ordered](n, p *rbNode[T]) (bro *rbNode[T], isLeft bool) {
	if p == nil {
		return nil, false
	}

	isLeft = false
	if p.right == n {
		isLeft = true
		return p.left, isLeft
	}
	return p.right, isLeft
}

// fixDeletionCase2 处理情况2：兄弟节点为红色
//
//	  P(B)                   S(B)
//	  /  \     变色+旋转      /  \
//	N(?) S(R)   ---->      P(R) SR(B)
//	     / \               / \
//	  SL(B) SR(B)        N(?) SL(B)
//
// 返回旋转后的新的根节点
func fixDeletionCase2[T cmp.Ordered](n, p *rbNode[T]) (newp *rbNode[T]) {
	bro, isLeft := getBrother(n, p) // 不可能为空
	bro.color = black
	p.color = red
	if isLeft {
		newp = rotateRight(p)
	} else {
		newp = rotateLeft(p)
	}
	return newp
}

// fixDeletionCase3_1 处理情况 3.1：兄弟节点为黑色，P、SL、SR也为黑色
//
//	  P(B)                        P(B)
//	  /  \    变色+向上调整         /  \
//	N(?) S(B)   ------>         N(?) S(R)
//	     / \                         / \
//	  SL(B) SR(B)                 SL(B) SR(B)
//
// 返回新的替代节点
func fixDeletionCase3_1[T cmp.Ordered](n, p *rbNode[T]) (newn *rbNode[T]) {
	bro, _ := getBrother(n, p)
	bro.color = red
	newn = p
	return newn
}

// fixDeletionCase3_2 处理情况 3.2：兄弟节点为黑色，SL、SR 为黑色；P 为红色。
//
//	  P(R)                        P(B)
//	  /  \      交换颜色           /  \
//	N(B) S(B)   ------>         N(B) S(R)
//	     / \                         / \
//	  SL(B) SR(B)                 SL(B) SR(B)
func fixDeletionCase3_2[T cmp.Ordered](n, p *rbNode[T]) {
	bro, _ := getBrother(n, p)
	bro.color = red
	p.color = black
}

// fixDeletionCase3_3 处理情况 3.3：兄弟节点为黑色，近侄子为红色，远侄子为黑色；P随意
//
//	  P(?)                        P(?)                     P(?)
//	  /  \     交换S跟SL颜色       /  \         旋转S        /  \
//	N(?) S(B)   ------>         N(?) S(R)    ----->      N(?) SL(B)
//	     / \                         / \                       \
//	  SL(R) SR(B)                 SL(B) SR(B)                  S(R)
//	                                                            \
//	                                                            SR(B)
//
// 转换成情况 3.4（还有 N(B)位于右子树的情况，旋转方向改变即可）
func fixDeletionCase3_3[T cmp.Ordered](n, p *rbNode[T]) (newp *rbNode[T]) {
	bro, isLeft := getBrother(n, p)
	bro.color = red
	if isLeft {
		bro.right.color = black
		p.left = rotateLeft(bro)
	} else {
		bro.left.color = black
		p.right = rotateRight(bro)
	}
	return fixDeletionCase3_4(n, p)
}

// fixDeletionCase3_4 处理情况 3.4：兄弟节点为黑色，远侄子为红色，P跟近侄子随意
//
//	  P(?)                        P(B)                      S(?)
//	  /  \  交换P跟S颜色+SR变色      /  \       旋转P          /  \
//	N(?) S(B)   ------>         N(?) S(?)    ------>      P(B) SR(B)
//	     / \                         / \                  / \
//	  SL(?) SR(R)                 SL(?) SR(B)           N(?) SL(?)
//
// 返回旋转后的新根节点（还有 N(B) 位于右子树的情况，旋转方向改变即可）
func fixDeletionCase3_4[T cmp.Ordered](n, p *rbNode[T]) (newp *rbNode[T]) {
	bro, isLeft := getBrother(n, p)
	bro.color = p.color
	p.color = black

	// 远侄子变黑色
	if isLeft {
		bro.left.color = black
		return rotateRight(p)
	} else {
		bro.right.color = black
		return rotateLeft(p)
	}
}
