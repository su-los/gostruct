// 计数排序实现
//
// - Out-Place，空间复杂度 O(k)，k 是序列最大元素大小
// - 稳定
// - 平均 O(n + k)、最好 O(n + k)、最坏 O(n + k)
package sort

type OrderType byte

const (
	OrderAsc  OrderType = iota // 升序
	OrderDesc                  // 降序
)

type Integer interface {
	SignedInteger | UnsignedInteger
}

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// CountingSort 计数排序
func CountingSort[T Integer](arr []T, maxVal T, order OrderType) []T {
	arrLen := len(arr)
	if arrLen == 0 {
		return []T{}
	}

	// 计数
	bucket := make([]int, maxVal+1)
	for i := range arr {
		bucket[arr[i]]++
	}

	// 计算前缀和
	for i := 1; i < len(bucket); i++ {
		bucket[i] += bucket[i-1]
	}

	// 逆序遍历原始数组，将对应元素移动到指定位置
	out := make([]T, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		val := arr[i]
		idx := bucket[val] - 1
		// 降序则倒转下设置的索引即可
		if order == OrderDesc {
			idx = arrLen - idx - 1
		}

		// 前面有 prefix 个元素小于等于它（包含它自己）
		// 所以它就在第 prefix - 1 的位置（从 0 开始计算）
		out[idx] = val
		bucket[val]--
	}
	return out
}

// CountingSortUnsigned 计数排序，对于正数数组，减少取值范围
func CountingSortUnsigned[T UnsignedInteger](arr []T, order OrderType) []T {
	arrLen := len(arr)
	if arrLen == 0 {
		return []T{}
	}

	minVal, maxVal := getMinAndMax(arr)
	for idx := range arr {
		arr[idx] -= minVal
	}
	out := CountingSort(arr, maxVal-minVal, order)

	for idx := range out {
		out[idx] += minVal
	}
	return out
}

// CountingSortSigned 计数排序，支持处理负数
func CountingSortSigned[T SignedInteger](arr []T, order OrderType) []T {
	arrLen := len(arr)
	if arrLen == 0 {
		return []T{}
	}

	minVal, maxVal := getMinAndMax(arr)
	if minVal < 0 {
		minVal *= -1
	}
	for idx := range arr {
		arr[idx] += minVal
	}
	out := CountingSort(arr, maxVal+minVal, order)

	for idx := range out {
		out[idx] -= minVal
	}
	return out
}

// ----------------------------------------------------
// Internal method
// ----------------------------------------------------

// getMinAndMax 求最大最小值
func getMinAndMax[T Integer](arr []T) (minVal, maxVal T) {
	minVal, maxVal = arr[0], arr[0]
	for idx := range arr {
		if minVal > arr[idx] {
			minVal = arr[idx]
		}

		if maxVal < arr[idx] {
			maxVal = arr[idx]
		}
	}
	return minVal, maxVal
}
