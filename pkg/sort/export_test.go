package sort

import (
	"cmp"
	"math/rand"
	"time"
)

// GenNumList 生成一个随机的层序遍历的序列
func GenNumList(size, maxVal int) []int {
	var (
		rd  = rand.New(rand.NewSource(int64(time.Now().UnixNano())))
		len = rd.Intn(size) + 1 // 序列长度)
		res = make([]int, 0, len)
	)

	for range len {
		num := rd.Intn(maxVal)
		res = append(res, num)
	}
	return res
}

func CheckIsAsc[T cmp.Ordered](arr []T) bool {
	for i := range arr {
		if i == len(arr)-1 {
			break
		}
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func CheckIsDesc[T cmp.Ordered](arr []T) bool {
	for i := range arr {
		if i == len(arr)-1 {
			break
		}
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}
