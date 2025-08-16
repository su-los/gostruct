package segtree

import (
	"cmp"
	"container/list"
	"errors"
	"fmt"
)

// buildRecursive 递归构建线段树
//
// seg 表示原始序列
// start, end 表示区间 [start, end)，左闭右开
// cnt 表示节点的数量
func buildRecursive[T cmp.Ordered](seg []T, start, end int, f AggFunc[T]) (root *segNode[T], cnt int) {
	if start >= end {
		return nil, 0
	}

	root = newSegNode(seg, start, end, f)
	if end-start == 1 {
		// 叶子节点
		return root, 1
	}

	l, r := start, end
	mid := (l + r) >> 1

	lcnt, rcnt := 0, 0
	// 只有两个元素的场景 [1, 2] l = 0, r = 2, mid = 1
	// 此时应该分成 [0,1] 和 [1,2]，所以左半部分需要是 [l, mid+1]
	if mid-l == 0 || r-mid == 0 {
		// 对于长度大于 2 的序列的划分，这里不应该出现划分出来的区间长度为 0 的情况
		panic("should not happen")
	}
	root.left, lcnt = buildRecursive(seg, l, mid, f)
	root.right, rcnt = buildRecursive(seg, mid, r, f)
	return root, lcnt + rcnt + 1
}

// build 非递归构建线段树
//
// 相当于通过层序遍历构建线段树
func build[T cmp.Ordered](seg []T, f AggFunc[T]) (root *segNode[T], cnt int) {
	segLen := len(seg)
	if segLen == 0 {
		return nil, 0
	}

	root = newSegNode(seg, 0, segLen, f)
	if segLen == 1 {
		// 叶子节点
		return root, 1
	}

	queue := []*segNode[T]{root}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		cnt++
		if (top.end - top.start) <= 1 {
			// 叶子节点
			continue
		}

		l, r := top.start, top.end
		mid := (l + r) >> 1
		// 对于元素只有两个的情况：[1,2]，l = 0, r = 2, mid = 1
		// 此时应该分成 [0,1] 和 [1,2]，所以左半部分需要是 [l, mid]
		if mid-l == 0 || r-mid == 0 {
			// 对于长度大于 2 的序列的划分，这里不应该出现划分出来的区间长度为 0 的情况
			panic("should not happen")
		}
		top.left = newSegNode(top.seg, l, mid, f)
		top.right = newSegNode(top.seg, mid, r, f)
		queue = append(queue, top.left, top.right)
	}
	return root, cnt
}

// buildBottomUp 自底向上构建线段树
//
// cnt 表示节点的数量
func buildBottomUp[T cmp.Ordered](seg []T, f AggFunc[T]) (root *segNode[T], cnt int) {
	queue, next := list.New(), list.New()
	// 初始化叶子节点
	for i := range seg {
		queue.PushBack(newSegNode(seg, i, i+1, f))
		cnt++
	}

	// 初始化非叶子节点
	for queue.Len() > 1 {
		l, r := queue.Front(), queue.Front().Next()
		lVal, rVal := l.Value.(*segNode[T]), r.Value.(*segNode[T])
		if lVal.end != rVal.start && rVal.end != lVal.start {
			// 理论上不会出现
			panic(fmt.Sprintf("invalid segNode: %v, %v", lVal, rVal))
		}

		queue.Remove(l)
		queue.Remove(r)
		start, end := min(lVal.start, rVal.start), max(lVal.end, rVal.end)
		peek := newSegNode(seg, start, end, f)
		peek.left = l.Value.(*segNode[T])
		peek.right = r.Value.(*segNode[T])
		// 合并后存储到下一层
		next.PushBack(peek)
		cnt++

		if queue.Len() == 1 {
			// 提升一层，进行合并
			next.PushBack(queue.Front().Value.(*segNode[T]))
			queue.Remove(queue.Front())
			queue, next = next, queue
		} else if queue.Len() == 0 {
			queue, next = next, queue
		}
	}
	return queue.Front().Value.(*segNode[T]), cnt
}

// ErrNotInRang 表示查询区间不在当前序列的范围内
var ErrNotInRang = errors.New("not in range")

// query 线段树的区间查询（递归）
//
// [l, r) 表示查询区间，左闭右开
// 返回区间 [l, r) 的和、最大值、最小值
func query[T cmp.Ordered](root *segNode[T], l, r int, f AggFunc[T]) (T, error) {
	if root == nil {
		return *new(T), errors.New("invaliTTegNode")
	}

	// 情况 1：无交集
	if root.start >= r || root.end <= l {
		return *new(T), ErrNotInRang
	}

	// 情况 2：[start, end) 完全包含于 [l, r)
	if root.start >= l && root.end <= r {
		return root.aggVal, nil
	}

	var (
		lVal, rVal T
		lErr, rErr error
	)
	// 情况 3：处理左右孩子有交集的情况
	if root.left != nil {
		var start, end int
		if start, end, lErr = getIntersect(root.left, l, r); lErr == nil {
			lVal, lErr = query(root.left, start, end, f)
		}
	}

	if root.right != nil {
		var start, end int
		if start, end, rErr = getIntersect(root.right, l, r); rErr == nil {
			rVal, rErr = query(root.right, start, end, f)
		}
	}

	if lErr != nil {
		return rVal, rErr
	} else if rErr != nil {
		return lVal, lErr
	} else {
		return f([]T{lVal, rVal}), nil
	}
}

// 获取交集区间
func getIntersect[T cmp.Ordered](root *segNode[T], l, r int) (int, int, error) {
	if root.end > l || root.start < r {
		return max(root.start, l), min(root.end, r), nil
	}
	return -1, -1, errors.New("no intersect")
}
