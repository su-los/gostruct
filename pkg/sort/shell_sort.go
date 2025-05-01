// 希尔排序实现
//
// In-Place，空间复杂度 O(1)
// 不稳定
// 平均 [O(NlogN), O(N^2)]，最坏 O(N^2)，最好 O(NlogN)
package sort

import (
	"cmp"
)

func insertSort[T cmp.Ordered](arr []T, start, gap int, less func(i, j int) bool) {
	var (
		alen = len(arr)
	)

	for i := start + gap; i < alen; i += gap {
		cur := i // i 为待排序部分的第一个元素，i-gap 为已排序部分的最后一个元素
		// 升序，将小的交换到前面
		for prev := cur - gap; (prev >= 0) && (less(cur, prev)); prev -= gap {
			arr[cur], arr[prev] = arr[prev], arr[cur]
			cur = prev
		}
	}
}

func ShellSort[T cmp.Ordered](arr []T, ordType OrderType) {
	if len(arr) == 0 {
		return
	}

	var (
		alen  = len(arr)
		group = alen >> 1 // 初始有 n/2 组，每组两个元素，组内使用插入排序，每组元素相隔 group
	)

	for group >= 1 {
		// 对每一组使用插入排序
		for i := range group {
			// i 是每一组起始的元素
			// group 是每一组元素相隔的距离
			if ordType == OrderAsc {
				insertSort(arr, i, group, func(i, j int) bool {
					return arr[i] < arr[j]
				})
			} else if ordType == OrderDesc {
				insertSort(arr, i, group, func(i, j int) bool {
					return arr[i] > arr[j]
				})
			}
		}
		group >>= 1
	}
}
