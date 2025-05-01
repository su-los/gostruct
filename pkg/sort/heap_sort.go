// 堆排序实现
//
// - In-Place 空间复杂度 O(1)
// - 不稳定，排序后相同元素的相对位置可能发生改变
// - 平均 O(NlogN)，最好 O(NlogN)，最坏 O(NlogN)
package sort

// MaxHeapSort 最大堆排序（升序）
func MaxHeapSort[T any](arr []T, greater func(lhs, rhs T) bool) {
	heapsort(arr, greater)
}

// MinHeapSort 最小堆排序（降序）
func MinHeapSort[T any](arr []T, less func(lhs, rhs T) bool) {
	heapsort(arr, less)
}

// down 将堆中位置为 x 的元素下沉
func down[T any](arr []T, x int, cmp func(lhs, rhs T) bool) {
	var (
		alen = len(arr)
		dpos = x // 需要下层的位置
	)

	for dpos < alen {
		pre := dpos // 下层前的节点
		// 左右孩子节点
		l, r := pre<<1+1, (pre+1)<<1

		// l 满足最大/最小堆的条件，就将 x 下层到 l
		if l < alen && cmp(arr[l], arr[dpos]) {
			dpos = l
		}

		// r 满足最大/最小堆的条件，就将 x 下层到 r
		if r < alen && cmp(arr[r], arr[dpos]) {
			dpos = r
		}

		// 因为是从最后一个非叶子节点调整上去的
		// 所以无需下沉了可以放心推出
		if dpos == pre {
			break
		}

		// 下沉
		arr[pre], arr[dpos] = arr[dpos], arr[pre]
	}
}

// heapify 堆化（自底向上构建 O(n)）
func heapify[T any](arr []T, cmp func(lhs, rhs T) bool) {
	if len(arr) == 0 {
		return
	}

	last := len(arr) - 1
	// 从最后一个非叶子节点开始
	for cur := (last - 1) >> 1; cur >= 0; cur-- {
		down(arr, cur, cmp)
	}
}

// heapsort 不断 Pop，对堆化后的数据进行排序
func heapsort[T any](arr []T, cmp func(lhs, rhs T) bool) {
	heapify(arr, cmp)
	var (
		alen = len(arr)
		last = alen - 1 // 堆最后一个元素
	)

	for last > 0 {
		// pop 操作：交换堆顶跟堆尾，堆大小减少 1，然后下沉新堆顶
		arr[0], arr[last] = arr[last], arr[0]
		last--
		down(arr[0:last+1], 0, cmp)
	}
}
