package rbtree

import "cmp"

type rbColor bool

const (
	red   rbColor = true
	black rbColor = false
)

// rbNode 红黑树节点
type rbNode[T cmp.Ordered] struct {
	val    T
	color  rbColor
	left   *rbNode[T]
	right  *rbNode[T]
	parent *rbNode[T]
}

// newRBNode 创建新的红黑树节点
func newRBNode[T cmp.Ordered](val T, color rbColor) *rbNode[T] {
	return &rbNode[T]{
		val:   val,
		color: color,
	}
}

// isRedViolation 判断节点的前后是否存在连续的红色
//
// true：节点的前后存在连续的红色；false: 符合红黑树定义；
func (rb *rbNode[T]) isRedViolation() bool {
	if rb.color == black {
		return false
	}
	if rb.parent != nil && rb.parent.color == red {
		return true
	}
	return (rb.left != nil && rb.left.color == red) ||
		(rb.right != nil && rb.right.color == red)
}

// verifyBlackHeightAndRedRules 迭代验证红黑树性质：
//
// 1. 黑高一致性检查
// 2. 红色节点连续性检查
//
// 返回：黑高，是否符合红黑树定义（true 表示符合）
func (rb *rbNode[T]) verifyBlackHeightAndRedRules() (int, bool) {
	var (
		cur         = rb
		lastVisited *rbNode[T]
		stack       = make([]*rbNode[T], 0)
		mpBlack     = make(map[*rbNode[T]]int, 0)
	)
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.left
			continue
		}
		peek := stack[len(stack)-1]
		if peek.right != nil && peek.right != lastVisited {
			cur = peek.right
			continue
		}
		// 访问节点
		stack = stack[:len(stack)-1]
		lastVisited = peek
		cur = nil
		// 检查红节点连续性
		if peek.isRedViolation() {
			return 0, false
		}
		selfIncrement := 0
		if peek.color == black {
			selfIncrement = 1
		}
		bhLeft, bhRight := 1, 1
		if peek.left != nil {
			bh, ok := mpBlack[peek.left]
			if !ok {
				return 0, false
			}
			bhLeft = bh
		}
		if peek.right != nil {
			bh, ok := mpBlack[peek.right]
			if !ok {
				return 0, false
			}
			bhRight = bh
		}
		// 检查黑高一致性
		if bhLeft != bhRight {
			return 0, false
		}
		mpBlack[peek] = bhLeft + selfIncrement
	}
	bh, ok := mpBlack[rb]
	return bh, ok
}

// minSubNode 求最小节点
func (rb *rbNode[T]) minSubNode() *rbNode[T] {
	cur := rb
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// getGrandparent 获取祖父节点
func (rb *rbNode[T]) getGrandparent() *rbNode[T] {
	if rb.parent != nil {
		return rb.parent.parent
	}
	return nil
}

// getUncle 获取叔父节点
func (rb *rbNode[T]) getUncle() *rbNode[T] {
	gp := rb.getGrandparent()
	if gp == nil {
		return nil
	}
	if rb.parent == gp.left {
		return gp.right
	}
	return gp.left
}
