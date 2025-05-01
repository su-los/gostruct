// 基数排序实现
//
// 稳定
// 空间复杂度：O(k)（Out-Place）
// 时间复杂度：O（n*k）
package sort

import "math"

// RadixSort 基数排序
func RadixSort[T Integer](arr []T, ordType OrderType) []T {
	if len(arr) == 0 {
		return []T{}
	}

	var (
		alen   = len(arr)
		res    = make([]T, alen) // 存储结果
		bkSize = 10
		bucket = make([][]T, bkSize) // 桶
		maxVal = arr[0]
	)
	for i := range arr {
		maxVal = max(maxVal, arr[i])
		res[i] = arr[i]
	}

	// 从最低位开始遍历
	for exp := T(1); (maxVal / exp) != 0; exp *= 10 {
		// 根据对应位分组
		for j := range res {
			dig := (res[j] / exp) % 10

			if ordType == OrderAsc {
				bucket[dig] = append(bucket[dig], res[j])
			} else if ordType == OrderDesc {
				// 降序：倒序放置（0 放到 9 号桶、1 放到 8 号桶...）
				idx := bkSize - int(dig) - 1
				bucket[idx] = append(bucket[idx], res[j])
			}
		}

		// 按顺序放入原序列
		start := 0
		for j := range bucket {
			if len(bucket[j]) == 0 {
				continue
			}
			copy(res[start:start+len(bucket[j])], bucket[j])
			start += len(bucket[j])
			// 注意：一定要重置 bucket
			bucket[j] = bucket[j][:0]
		}
	}
	return res
}

// -----------------------------------------------------------
// RadixSortByCounting 基数排序（内部使用计数排序进行排序）
func RadixSortByCounting[T Integer](arr []T, ordType OrderType) []T {
	if len(arr) == 0 {
		return []T{}
	}

	var (
		alen   = len(arr)
		res    = make([]T, alen)
		maxVal = arr[0]
	)
	for i := range arr {
		maxVal = max(maxVal, arr[i])
		res[i] = arr[i]
	}

	digitCnt := 0
	for exp := T(1); maxVal/exp != 0; exp *= 10 {
		digitCnt++
	}

	for i := range digitCnt {
		res = countingByDigit(res, i+1, ordType)
	}
	return res
}

// countingByDigit 对每位进行计数排序
//
// digit 位数，从 1 开始
func countingByDigit[T Integer](arr []T, digit int, ordType OrderType) []T {
	var (
		alen       = len(arr)
		exp        = T(math.Pow10(digit) / 10)
		res        = make([]T, alen)
		bucketSize = 10
		counting   = make([]int, bucketSize)
	)

	for i := range arr {
		idx := (arr[i] / exp) % 10
		counting[idx]++
	}

	// 计算前缀和
	if ordType == OrderAsc {
		for i := range counting {
			if i == 0 {
				continue
			}
			counting[i] += counting[i-1]
		}
	} else if ordType == OrderDesc {
		// 只能修改前缀和的定义，prefix[i] 表示大于等于 i 的数据的个数
		// 如果像普通 CountingSort 那样在放置的时候计算逆序，会导致以下情况无法处理：
		// digit = 4，对于 []{668,117}，count[0]=2，导致 668 跟 117 会放置错误
		for i := bucketSize - 1; i >= 0; i-- {
			if i == bucketSize-1 {
				continue
			}
			counting[i] += counting[i+1]
		}
	}

	// 按前缀和放置
	for i := alen - 1; i >= 0; i-- {
		idx := (arr[i] / exp) % 10
		res[counting[idx]-1] = arr[i]
		counting[idx]--
	}
	return res
}
