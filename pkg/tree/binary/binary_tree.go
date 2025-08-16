package binary

import "cmp"

// BinaryTree 二叉树.
type BinaryTree[T cmp.Ordered] struct {
	root *TreeNode[T]
}

func NewBinaryTree[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

// FromLevelOrder 根据层序遍历结果反序列化成一颗树
//
// 层序遍历结果要求带上空节点（使用空值替代）.
func FromLevelOrder[T cmp.Ordered](data []T) *BinaryTree[T] {
	size := len(data)
	if len(data) == 0 {
		return &BinaryTree[T]{
			root: nil,
		}
	}

	res := &BinaryTree[T]{
		root: &TreeNode[T]{
			Val:   data[0],
			Left:  nil,
			Right: nil,
		},
	}

	var (
		zero  T // 定义空值，或者用反射判断 reflect.ValueOf(v).IsZero()
		queue = make([]*TreeNode[T], 0, size)
	)
	queue = append(queue, res.root)
	for idx := 1; idx < size; {
		cur := queue[0]
		queue = queue[1:]

		// 构建左子树
		if data[idx] == zero {
			cur.Left = nil
		} else {
			cur.Left = &TreeNode[T]{
				Val:   data[idx],
				Left:  nil,
				Right: nil,
			}
			queue = append(queue, cur.Left)
		}

		if idx+1 >= size {
			cur.Right = nil

			break
		}

		// 构建右子树
		if data[idx+1] == zero {
			cur.Right = nil
		} else {
			cur.Right = &TreeNode[T]{
				Val:   data[idx+1],
				Left:  nil,
				Right: nil,
			}
			queue = append(queue, cur.Right)
		}
		idx += 2
	}

	return res
}

// Root 获取根节点.
func (t *BinaryTree[T]) Root() *TreeNode[T] {
	return t.root
}

// LevelOrder 层序遍历.
func (t *BinaryTree[T]) LevelOrder() []T {
	if t.root == nil {
		return []T{}
	}

	var (
		res   = make([]T, 0)
		queue = make([]*TreeNode[T], 0)
	)
	queue = append(queue, t.root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return res
}

// LevelOrderRecursive 层序遍历递归实现.
func (t *BinaryTree[T]) LevelOrderRecursive() []T {
	if t.root == nil {
		return []T{}
	}

	res := make([]T, 0)
	var helper func([]*TreeNode[T])
	helper = func(nodes []*TreeNode[T]) {
		if len(nodes) == 0 {
			return
		}

		next := make([]*TreeNode[T], 0, len(nodes)*2)
		for _, node := range nodes {
			res = append(res, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		helper(next)
	}

	helper([]*TreeNode[T]{t.root})

	return res
}
