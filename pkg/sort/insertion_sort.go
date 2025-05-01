// 插入排序实现
//
// - In-Place，空间复杂度 O(1)
// - 稳定
// - 平均 O(N^2)、最好 O(N)，最坏 O(N^2)
package sort

import "cmp"

// InsertionSort 插入排序
func InsertionSort[T cmp.Ordered](arr []T, ordType OrderType) {
	arrLen := len(arr)
	// 遍历未排序部分
	// sorted 初始等于 0，表示已排序部分的最后一个元素
	for unsorted := 1; unsorted < arrLen; unsorted++ {
		cur := unsorted // 未排序部分第一个元素
		// 插入已排序部分（从后往前不断交换即可）
		for prev := cur - 1; prev >= 0; prev-- {
			if ordType == OrderAsc {
				// 升序，交换小的元素到前面
				if arr[cur] < arr[prev] {
					arr[cur], arr[prev] = arr[prev], arr[cur]
					cur = prev
				} else {
					// 已经无需交换
					break
				}
			} else if ordType == OrderDesc {
				// 降序，交换大的元素到前面
				if arr[cur] > arr[prev] {
					arr[cur], arr[prev] = arr[prev], arr[cur]
					cur = prev
				} else {
					// 已经无需交换
					break
				}
			}
		}
	}
}
