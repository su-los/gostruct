// Package binary 二叉树定义
package binary

import "cmp"

// TreeNode 普通的二叉树节点定义.
type TreeNode[T cmp.Ordered] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}
