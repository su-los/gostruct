// 线段树定义
package segtree

import (
	"cmp"
	"fmt"
)

// SegTree 线段树定义
type SegTree[T cmp.Ordered] struct {
	root *segNode[T]
	f    AggFunc[T] // 聚合函数
	len  int
}

// NewSegTree 构建线段树
func NewSegTree[T cmp.Ordered](seg []T, f AggFunc[T]) *SegTree[T] {
	st := &SegTree[T]{
		f: f,
	}
	st.root, st.len = build(seg, f)
	return st
}

// Query 线段树的区间查询
func (st *SegTree[T]) Query(l, r int) (T, error) {
	if l < 0 || r >= st.len || l > r {
		return *new(T), fmt.Errorf("invalid range [%d, %d)", l, r)
	}
	return query(st.root, l, r, st.f)
}

// LevelOrder 层序遍历的结果
func (st *SegTree[T]) LevelOrder() [][]T {
	if st.root == nil {
		return nil
	}

	var (
		queue = []*segNode[T]{st.root}
		res   = make([][]T, 0, st.len)
	)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		res = append(res, top.Segment())
		if top.left != nil {
			queue = append(queue, top.left)
		}

		if top.right != nil {
			queue = append(queue, top.right)
		}
	}
	return res
}

// PreOrder 前序遍历的结果
func (st *SegTree[T]) PreOrder() [][]T {
	if st.root == nil {
		return nil
	}

	var (
		stack = make([]*segNode[T], 1, st.len)
		res   = make([][]T, 0, st.len)
	)
	stack[0] = st.root

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Segment())

		if top.right != nil {
			stack = append(stack, top.right)
		}

		if top.left != nil {
			stack = append(stack, top.left)
		}
	}
	return res
}

// InOrder 中序遍历的结果
func (st *SegTree[T]) InOrder() [][]T {
	if st.root == nil {
		return nil
	}

	var (
		stack = make([]*segNode[T], 0, st.len)
		res   = make([][]T, 0, st.len)
		cur   = st.root
	)

	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.left
			continue
		}

		// 回溯到根节点，直接访问即可
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Segment())
		cur = top.right
	}
	return res
}

// PostOrder 后序遍历的结果
func (st *SegTree[T]) PostOrder() [][]T {
	if st.root == nil {
		return nil
	}

	var (
		stack                 = make([]*segNode[T], 0, st.len)
		res                   = make([][]T, 0, st.len)
		cur, prev *segNode[T] = st.root, nil
	)

	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.left
			continue
		}

		// 回溯到根节点
		peek := stack[len(stack)-1]
		if peek.right != nil && peek.right != prev {
			// 不是从右子树回溯的，先访问右子树
			cur = peek.right
		} else {
			// 从右子树回溯的，直接访问即可
			res = append(res, peek.Segment())
			stack = stack[:len(stack)-1]

			// 记录状态，peek 与其左右子树都已经访问完成
			// 设置 cur = nil，触发访问 peek 的父节点
			prev = peek
			cur = nil
		}
	}
	return res
}
