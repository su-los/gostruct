// 冒泡排序实现
//
// - In-Place、O(1)
// - 稳定
// - 平均 O(N^2)，最好 O(N)，最坏 O(N^2)
package sort

import "cmp"

// BubbleSort 冒泡排序
func BubbleSort[T cmp.Ordered](arr []T, ordType OrderType) {
	arrLen := len(arr)
	for i := arrLen - 1; i >= 0; i-- {
		isSwapped := false
		// 对序列 [0, i] 进行一次冒泡
		for j := range arr {
			if j == i {
				break
			}

			if ordType == OrderAsc && arr[j] > arr[j+1] {
				// 升序：前面的大，则交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			} else if ordType == OrderDesc && arr[j] < arr[j+1] {
				// 降序：前面的小，则交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			}
		}

		// 一次都没有交换过，证明序列有序
		// 最好情况的时间复杂度为 O(N) 就体现在这里
		if !isSwapped {
			break
		}
	}
}
