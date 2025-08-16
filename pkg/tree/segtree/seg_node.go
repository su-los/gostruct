// 线段树节点定义
package segtree

import "cmp"

// AggFunc 提供聚合功能的函数
type AggFunc[T any] func([]T) T

// Sum 求和
func Sum[T cmp.Ordered](seg []T) T {
	var sum T
	for _, v := range seg {
		sum += v
	}
	return sum
}

// Max 求最大值
func Max[T cmp.Ordered](seg []T) T {
	var max T
	for _, v := range seg {
		if v > max {
			max = v
		}
	}
	return max
}

// Min 求最小值
func Min[T cmp.Ordered](seg []T) T {
	var min T
	for _, v := range seg {
		if v < min {
			min = v
		}
	}
	return min
}

// 也可以采用堆式存储来构建，节点编号从 1 开始
// 对于编号为 i 的节点，左孩子编号为 2i，右孩子为 2i+1
type segNode[T cmp.Ordered] struct {
	seg        []T // 原始序列
	start, end int // 区间 [start, end)，左闭右开
	aggVal     T   // 聚合值，表示区间和、区间最大值、区间最小值等，由聚合函数决定
	left       *segNode[T]
	right      *segNode[T]
}

// newSegNode 构建线段树节点
//
// seg 表示原始序列
// l, r 表示区间 [l, r)，左闭右开
func newSegNode[T cmp.Ordered](seg []T, l, r int, f AggFunc[T]) *segNode[T] {
	return &segNode[T]{
		seg:    seg,
		start:  l,
		end:    r,
		aggVal: f(seg[l:r]),
	}
}

// Segment 返回当前节点的区间
func (s *segNode[T]) Segment() []T {
	return s.seg[s.start:s.end]
}
