// 归并排序实现
//
// Out-Place 空间复杂度 $O(n)$
// 稳定
// 时间复杂度 $O(n\log n)$
package sort

import (
	"cmp"
)

// combine 合并两个相邻序列
//
// dst：合并数据存放的位置
// 调用方保证 len(dst) = len(left) + len(right)
func combine[T cmp.Ordered](dst []T, left, right []T, ordType OrderType) {
	var (
		lsize, rsize = len(left), len(right)
		lidx, ridx   = 0, 0
	)

	for i := range lsize + rsize {
		if lidx >= lsize {
			dst[i] = right[ridx]
			ridx++
			continue
		}

		if ridx >= rsize {
			dst[i] = left[lidx]
			lidx++
			continue
		}

		if ordType == OrderAsc {
			// 升序：将小的先复制过去
			if left[lidx] < right[ridx] {
				dst[i] = left[lidx]
				lidx++
			} else {
				dst[i] = right[ridx]
				ridx++
			}
		} else if ordType == OrderDesc {
			// 降序：将大的先复制过去
			if left[lidx] > right[ridx] {
				dst[i] = left[lidx]
				lidx++
			} else {
				dst[i] = right[ridx]
				ridx++
			}
		}
	}
}

// MergeSort 归并排序非递归实现
func MergeSort[T cmp.Ordered](arr []T, ordType OrderType) []T {
	var (
		alen = len(arr)
		// out 跟 merge 都是在合并过程存储合并后结果用
		// out 存储待合并元素，mer 存储合并后元素
		out, mer = make([]T, alen), make([]T, alen)
		// 每组元素个数
		gpsize = 1
	)
	copy(out, arr)

	for gpsize < alen {
		// 相邻两个组的起始位置以及结束位置，遍历所有组
		g1start, g2start := 0, gpsize
		for g1start < alen && g2start < alen {
			g2end := min(g2start+gpsize, alen)

			combine(mer[g1start:g2end], out[g1start:g2start], out[g2start:g2end], ordType)
			// 下一个相邻两个组
			g1start, g2start = g2end, g2end+gpsize
		}

		// 处理最后一组的情况
		if g1start < alen {
			copy(mer[g1start:], out[g1start:])
		}

		// 扩大分组
		out, mer = mer, out
		gpsize <<= 1
	}

	return out
}

// mergeSortRe 递归实现（自上而下）
func mergeSortRe[T cmp.Ordered](arr []T, ordType OrderType) []T {
	alen := len(arr)
	if alen == 0 || alen == 1 {
		return arr
	}

	var (
		mid         = alen >> 1
		left, right []T
	)

	if mid > 0 {
		left = arr[0:mid]
	}
	if mid < alen {
		right = arr[mid:alen]
	}

	lefsorted := mergeSortRe(left, ordType)
	rigsorted := mergeSortRe(right, ordType)

	out := make([]T, alen)
	combine(out[0:alen], lefsorted, rigsorted, ordType)
	return out
}
