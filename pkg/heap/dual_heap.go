// 对顶堆的实现
package heap

import (
	"cmp"
	"errors"
)

// DualHeap 对顶堆
type DualHeap[T cmp.Ordered] struct {
	k       int
	maxHeap *Heap[T] // 大根堆，存放剩余的元素
	minHeap *Heap[T] // 小根堆，存放前 k 个比其他大的元素
}

// NewDualHeap 新建对顶堆
func NewDualHeap[T cmp.Ordered](k int) (*DualHeap[T], error) {
	if k <= 0 {
		return nil, errors.New("invalid k input")
	}

	return &DualHeap[T]{
		k:       k,
		maxHeap: NewMaxHeap[T](nil),
		minHeap: NewMinHeap[T](nil),
	}, nil
}

// Push 存入一个元素到对顶堆中
func (dh *DualHeap[T]) Push(val T) {
	dh.minHeap.Push(val)
	dh.balance()
}

// Pop 取出第 k 大的元素
func (dh *DualHeap[T]) Pop() (T, bool) {
	val, ok := dh.minHeap.Pop()
	dh.balance()
	return val, ok
}

// TopK 读取第 k 大的元素
func (dh *DualHeap[T]) TopK() (T, bool) {
	var ret T
	if len(dh.minHeap.data) == 0 {
		return ret, false
	}

	ret = dh.minHeap.data[0]
	return ret, true
}

// UpdateK 更新 k 值
func (dh *DualHeap[T]) UpdateK(newK int) error {
	if newK <= 0 {
		return errors.New("invalid k input")
	}
	dh.k = newK
	dh.balance()
	return nil
}

// balance 维护对顶堆
func (dh *DualHeap[T]) balance() {
	// 小根堆中的元素小于 k，需要从大根堆中调取
	for len(dh.minHeap.data) < dh.k {
		if val, ok := dh.maxHeap.Pop(); ok {
			dh.minHeap.Push(val)
		} else {
			// 大根堆中也没有元素了
			break
		}
	}

	// 小根堆中的元素大于了 k，需要转移元素至大根堆
	for len(dh.minHeap.data) > dh.k {
		if val, ok := dh.minHeap.Pop(); ok {
			dh.maxHeap.Push(val)
		} else {
			// 小根堆中没有元素了（不可能发生）
			break
		}
	}
}
