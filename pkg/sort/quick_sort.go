// 快速排序实现
//
// In-Place 空间复杂度 O(1)
// 不稳定
// 平均/最好 O(nlogn)，最坏 O(n^2)
package sort

import (
	"cmp"
)

/*
// QuickSort 快速排序（相当于是一个前序遍历递归实现）
func QuickSort[T cmp.Ordered](arr []T, ordType OrderType) {
	alen := len(arr)
	if alen <= 1 {
		return
	}

	var (
		midx  = 0
		pivot = arr[alen-1]
	)
	for i := range arr {
		if ordType == OrderAsc && arr[i] <= pivot {
			// 升序：将小于基准值的交换到左边
			arr[i], arr[midx] = arr[midx], arr[i]
			midx++
		} else if ordType == OrderDesc && arr[i] >= pivot {
			// 降序：将大于等于基准值的交换到左边
			arr[i], arr[midx] = arr[midx], arr[i]
			midx++
		}
	}

	// 到了这一步，midx-1 处就是基准值，所以需要排除
	QuickSort(arr[0:midx-1], ordType)
	QuickSort(arr[midx:alen], ordType)
}
*/

// 快速排序的非递归实现

type quickNode struct {
	left, right int // 序列的左边与右边指针
}

// partition 对 arr[left:right+1] 进行一次快速排序
//
// return：基准值最后所在的索引（-1 则表示已经有序，无需基准值）
func partition[T cmp.Ordered](arr []T, left, right int, ordType OrderType) int {
	alen := right - left + 1
	if alen <= 1 {
		return -1
	}

	var (
		midx  = left
		pivot = arr[right] // 基准值
	)
	for i := left; i <= right; i++ {
		if ordType == OrderAsc && arr[i] <= pivot {
			// 升序：将小于等于基准值的元素交换到左边
			arr[i], arr[midx] = arr[midx], arr[i]
			midx++
		} else if ordType == OrderDesc && arr[i] >= pivot {
			// 降序，将大于等于基准值的元素交换到左边
			arr[i], arr[midx] = arr[midx], arr[i]
			midx++
		}
	}
	return midx - 1
}

// QuickSort 快速排序，非递归实现（前序遍历非递归实现）
//
// 空间复杂度 O(n)
func QuickSort[T cmp.Ordered](arr []T, ordType OrderType) {
	var (
		stack = []*quickNode{
			{
				left:  0,
				right: len(arr) - 1,
			},
		}
	)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		midx := partition(arr, top.left, top.right, ordType)

		if midx == -1 {
			continue
		}

		// 先加入右子树
		stack = append(stack, &quickNode{
			left:  midx + 1,
			right: top.right,
		}, &quickNode{
			left:  top.left,
			right: midx - 1,
		})
	}
}
