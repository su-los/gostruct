// 选择排序的实现
//
// In-Place，空间复杂度 O(1)
// 不稳定
// 平均、最好、最坏都是 O(N^2)
package sort

import "cmp"

func SelectionSort[T cmp.Ordered](arr []T, ordType OrderType) {
	var (
		sorted = -1 // 已排序部分的最后一个元素
		arrLen = len(arr)
	)

	// 已排序部分最后一个元素等于数组最后一个元素了
	// 说明数组全部排序完毕
	for sorted < arrLen-1 {
		mostIdx := sorted + 1
		for i := sorted + 1; i < arrLen; i++ {
			if ordType == OrderAsc && arr[i] < arr[mostIdx] {
				// 升序：则需要找到最小值，放到左边的已排序部分
				mostIdx = i
			} else if ordType == OrderDesc && arr[i] > arr[mostIdx] {
				// 降序：则需要找到最大值，放到左边的已排序部分
				mostIdx = i
			}
		}

		arr[sorted+1], arr[mostIdx] = arr[mostIdx], arr[sorted+1]
		sorted++
	}
}
