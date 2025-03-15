package rbtree

import (
	"cmp"
	"errors"
)

// RBTree 红黑树
type RBTree[T cmp.Ordered] struct {
	root *rbNode[T]
}

// NewRBTree 创建红黑树
func NewRBTree[T cmp.Ordered]() *RBTree[T] {
	return &RBTree[T]{
		root: nil,
	}
}

// IsValid 验证红黑树所有性质：
//
// 1. 根节点为黑色
// 2. 红色节点不能连续出现
// 3. 所有路径黑高一致
// 4. 叶子节点（NIL）视为黑色
func (rb *RBTree[T]) IsValid() bool {
	if rb.root == nil {
		return true
	}
	if rb.root.color != black {
		return false
	}
	_, ok := rb.root.verifyBlackHeightAndRedRules()
	return ok
}

// Insert 插入
func (rb *RBTree[T]) Insert(val T) error {
	newNode := newRBNode(val, black)
	if rb.root == nil {
		rb.root = newNode
		return nil
	}
	var (
		cur    = rb.root
		parent *rbNode[T]
	)
	for cur != nil {
		switch cmp.Compare(val, cur.val) {
		case -1:
			parent = cur
			cur = cur.left
		case 0:
			// 重复元素，无需插入
			return nil
		case +1:
			parent = cur
			cur = cur.right
		}
	}
	newNode.color = red
	newNode.parent = parent
	res := cmp.Compare(newNode.val, parent.val)
	if res == -1 {
		parent.left = newNode
	} else if res == +1 {
		parent.right = newNode
	} else {
		return errors.ErrUnsupported
	}
	return rb.fixInsertion(newNode)
}

// Find 查找
func (rb *RBTree[T]) Find(val T) *rbNode[T] {
	cur := rb.root
	for cur != nil {
		switch cmp.Compare(val, cur.val) {
		case -1:
			cur = cur.left
		case 0:
			return cur
		case +1:
			cur = cur.right
		}
	}
	return nil
}

// Delete 删除
func (rb *RBTree[T]) Delete(val T) error {
	if rb.root == nil {
		return nil
	}

	del := rb.Find(val)
	if del == nil {
		return nil
	}

	var (
		// delColor 删除节点的颜色
		delColor = del.color
		// n 是替代节点，p 是替代节点的父节点
		n, p *rbNode[T]
	)

	if del.left == nil {
		n = del.right
		p = del.parent
		rb.transplant(del, del.right)
	} else if del.right == nil {
		n = del.left
		p = del.parent
		rb.transplant(del, del.left)
	} else {
		minRight := del.right.minSubNode()
		del.val = minRight.val // 注意：只改变值，不改变颜色

		// 替代为删除右子树最小的节点
		delColor = minRight.color
		n = minRight.right
		p = minRight.parent
		rb.transplant(minRight, minRight.right)
	}

	if delColor == black {
		return rb.fixDeletion(n, p)
	}
	return nil
}

// -------------------------------------------------------------------
// Internal method
// -------------------------------------------------------------------

// transplant 使用 target 节点替代 src 节点
//
// src 的子节点不会在这个函数中继承给 target
func (rb *RBTree[T]) transplant(src, target *rbNode[T]) {
	// 父节点绑定给 target
	if src.parent == nil {
		rb.root = target
	} else if src.parent.left == src {
		src.parent.left = target
	} else {
		src.parent.right = target
	}

	// 设置 target 的父节点指针
	if target != nil {
		target.parent = src.parent
	}
}

// fixInsertion 处理插入后不平衡情况
func (rb *RBTree[T]) fixInsertion(node *rbNode[T]) error {
	cur := node
	if cur.parent == nil {
		return nil
	} else if cur.parent.color == black {
		return nil
	}

	// parent 是红色的话，一定会有祖父节点，因为红色无法作为根节点
	for cur.parent != nil && cur.parent.color == red {
		var (
			gp           = cur.getGrandparent()
			parent       = cur.parent
			uncle        = cur.getUncle()
			isUncleBlack = isBlack(uncle)
			fixFunc      FixFunc[T]
		)
		switch {
		case !isUncleBlack:
			// 情况2 父节点跟叔父节点都是红色
			parent.color = black
			uncle.color = black
			gp.color = red
			cur = gp
			continue
		case isUncleBlack && gp.left == parent && parent.left == cur:
			// 情况 3：父节点为红色，叔父节点为黑色，插入节点是 LL
			fixFunc = fixInsertionCase3LL
		case isUncleBlack && gp.left == parent && parent.right == cur:
			// 情况 4：父节点为红色，叔父节点为黑色，插入节点是 LR
			fixFunc = fixInsertionCase4LR
		case isUncleBlack && gp.right == parent && parent.right == cur:
			// 情况 3：父节点为红色，叔父节点为黑色，插入节点是 RR
			fixFunc = fixInsertionCase3RR
		case isUncleBlack && gp.right == parent && parent.left == cur:
			// 情况 4：父节点为红色，叔父节点为黑色，插入节点是 RL
			fixFunc = fixInsertionCase4RL
		default:
			return errors.New("invalid case")
		}
		// 注意点 1：先判断再调用 Fix 函数，否则调用 Fix 函数后，节点的关系就变了 gp 的 parent 就不是以前的 parent 了
		// 注意点 2：这里还要先缓存 gp.parent，因为gp.parent.right = fixFunc(cur) 是先执行 fixFunc 在赋值，此时 gp.parent 可能已经变了
		ggp := gp.parent
		if ggp == nil {
			rb.root = fixFunc(cur)
		} else if ggp.left == gp {
			ggp.left = fixFunc(cur)
		} else if ggp.right == gp {
			ggp.right = fixFunc(cur)
		} else {
			return errors.New("gp parent is invalid")
		}
		break
	}
	// 注意：情况 2 可能把根节点染成红色，所以这里要处理下
	rb.root.color = black
	return nil
}

// fixDeletion 处理插入后不平衡的情况。
//
// @params n 表示删除后的替代节点（可能为 nil），p 表示 n 的父节点
func (rb *RBTree[T]) fixDeletion(n, p *rbNode[T]) error {
	var (
		bro, isBroLeft = getBrother(n, p)
		parent         = p // 替代节点的父节点
		cur            = n // 当前的替代节点
	)

	for parent != nil {
		var (
			gp       = parent.parent // 祖父
			fixNode  *rbNode[T]
			isGPLeft = (gp != nil && gp.left == parent)
		)

		if bro == nil {
			// 除了根节点，黑色节点，不可能没有兄弟节点
			return errors.New("not brother node, invalid")
		} else if bro.color == red {
			// 情况 2:兄弟节点为红色
			if gp == nil {
				rb.root = fixDeletionCase2(cur, parent)
			} else if isGPLeft {
				gp.left = fixDeletionCase2(cur, parent)
			} else {
				gp.right = fixDeletionCase2(cur, parent)
			}
			// ⚠️注意点：旋转后替代节点跟父节点虽然没有变，但是兄弟节点改变了！
			bro, isBroLeft = getBrother(cur, parent)
			continue
		} else {
			slBlack, srBlack, pBalck := isBlack(bro.left), isBlack(bro.right), isBlack(parent)
			if slBlack && srBlack && pBalck {
				// 情况 3.1 全黑
				cur = fixDeletionCase3_1(cur, parent)
				parent = cur.parent
				bro, isBroLeft = getBrother(cur, parent)
				continue
			} else if slBlack && srBlack && !pBalck {
				// 情况 3.2 p 红，其他黑
				fixDeletionCase3_2(cur, parent)
				break
			} else if (isBroLeft && slBlack && !srBlack) ||
				(!isBroLeft && srBlack && !slBlack) {
				// 情况 3.3 远侄子黑，近侄子红
				fixNode = fixDeletionCase3_3(cur, parent)
			} else if (isBroLeft && !slBlack) || (!isBroLeft && !srBlack) {
				// 情况 3.4 远侄子红，其他随意
				fixNode = fixDeletionCase3_4(cur, parent)
			}

			if gp == nil {
				rb.root = fixNode
			} else if isGPLeft {
				gp.left = fixNode
			} else {
				gp.right = fixNode
			}
			break
		}
	}

	if rb.root != nil {
		rb.root.color = black
	}
	return nil
}
