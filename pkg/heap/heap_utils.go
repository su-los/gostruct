package heap

import "cmp"

// up 以索引 x 的节点为起点向上调整
func up[T cmp.Ordered](arr []T, x int, htype heapType) {
	arrLen := len(arr)
	if arrLen == 0 || x >= arrLen {
		return
	}

	cur := x
	for cur > 0 {
		parent := (cur - 1) >> 1
		if htype == maxHeap && arr[cur] > arr[parent] {
			arr[cur], arr[parent] = arr[parent], arr[cur]
		} else if htype == minHeap && arr[cur] < arr[parent] {
			arr[cur], arr[parent] = arr[parent], arr[cur]
		} else {
			cur--
		}
	}
}

// down 以索引 x 节点为起点的向下调整
func down[T cmp.Ordered](arr []T, x int, htype heapType) {
	arrLen := len(arr)
	if arrLen == 0 || x >= arrLen {
		return
	}

	cur := x
	for cur < arrLen {
		left, right := (cur<<1)+1, (cur+1)<<1
		submin, submax, ok := subCompare(arr, left, right)
		if !ok {
			// 没有子节点了，无需下移
			break
		}

		if htype == maxHeap && arr[cur] < arr[submax] {
			arr[cur], arr[submax] = arr[submax], arr[cur]
			cur = submax
		} else if htype == minHeap && arr[cur] > arr[submin] {
			arr[cur], arr[submin] = arr[submin], arr[cur]
			cur = submin
		} else {
			// 子树已经满足堆条件了，结束移动
			break
		}
	}
}

// subCompare 找出左右子树的最大值以及最小值
func subCompare[T cmp.Ordered](arr []T, left, right int) (min, max int, ok bool) {
	arrLen := len(arr)
	if left >= arrLen && right >= arrLen {
		return -1, -1, false
	}

	if left < arrLen && right < arrLen {
		if arr[left] > arr[right] {
			min, max = right, left
		} else {
			min, max = left, right
		}
	} else if left < arrLen && right >= arrLen {
		min, max = left, left
	} else {
		min, max = right, right
	}
	return min, max, true
}

// heapifyUpwards 从底向上构建堆
func heapifyUpwards[T cmp.Ordered](arr []T, htype heapType) {
	arrLen := len(arr)
	if arrLen == 0 {
		return
	}

	cur := (arrLen >> 1) - 1 // 最后一个非叶子节点
	for ; cur >= 0; cur-- {
		down(arr, cur, htype)
	}
}
