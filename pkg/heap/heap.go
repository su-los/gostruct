package heap

import "cmp"

type heapType bool

const (
	minHeap heapType = true
	maxHeap heapType = false
)

// Heap 堆的数组实现
type Heap[T cmp.Ordered] struct {
	data  []T
	htype heapType // 堆的类型
}

// NewMaxHeap 新建一个大顶堆
func NewMaxHeap[T cmp.Ordered](arr []T) *Heap[T] {
	h := &Heap[T]{
		data:  make([]T, len(arr)),
		htype: maxHeap,
	}

	copy(h.data, arr)
	heapifyUpwards(h.data, h.htype)
	return h
}

// NewMinHeap 新建一个小顶堆
func NewMinHeap[T cmp.Ordered](arr []T) *Heap[T] {
	h := &Heap[T]{
		data:  make([]T, len(arr)),
		htype: minHeap,
	}
	copy(h.data, arr)
	heapifyUpwards(h.data, h.htype)
	return h
}

// IsValidHeap 判断是否是堆
func (h *Heap[T]) IsValidHeap() bool {
	arrLen := len(h.data)
	if arrLen == 0 {
		return true
	}

	for idx := range h.data {
		lt, rt := (idx<<1)+1, (idx<<1)+2
		if lt >= arrLen && rt >= arrLen {
			continue
		}

		if lt < arrLen {
			if h.htype == maxHeap && h.data[idx] < h.data[lt] {
				return false
			} else if h.htype == minHeap && h.data[idx] > h.data[lt] {
				return false
			}
		}

		if rt < arrLen {
			if h.htype == maxHeap && h.data[idx] < h.data[rt] {
				return false
			} else if h.htype == minHeap && h.data[idx] > h.data[rt] {
				return false
			}
		}
	}
	return true
}

// Push 往堆中存入一个元素
func (h *Heap[T]) Push(val T) {
	if len(h.data) == 0 {
		h.data = make([]T, 0)
	}

	// 添加到最后一个叶子节点
	h.data = append(h.data, val)
	// 向上调整
	up(h.data, len(h.data)-1, h.htype)
}

// Pop 弹出堆顶元素
func (h *Heap[T]) Pop() (T, bool) {
	var ret T
	if len(h.data) == 0 {
		return ret, false
	}

	// 将对顶元素与最后叶子节点交换再弹出
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	ret = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	// 向下调整
	down(h.data, 0, h.htype)
	return ret, true
}

// Update 修改某个节点的权值
func (h *Heap[T]) Update(val T, newVal T) bool {
	if val == newVal {
		return true
	}

	var (
		idx     = 0
		isexist = false
	)
	for idx = range h.data {
		if h.data[idx] == val {
			h.data[idx] = newVal
			isexist = true
			break
		}
	}

	if !isexist {
		return false // 没有找到
	}

	switch h.htype {
	case maxHeap:
		if newVal > val {
			// 增大了，可能破坏其上层的堆特性，因此需要向上调整
			up(h.data, idx, h.htype)
		} else {
			// 反之，可能破坏其下层的堆特性
			down(h.data, idx, h.htype)
		}
	case minHeap:
		if newVal > val {
			// 增大了，可能破坏最小堆下层特性，因此需要向下调整
			down(h.data, idx, h.htype)
		} else {
			up(h.data, idx, h.htype)
		}
	default:
		// 非法类型
		return false
	}
	return true
}
