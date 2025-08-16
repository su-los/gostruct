package segtree

import (
	"math/rand/v2"
	"time"
)

// GenNumList 生成一个随机的层序遍历的序列
func GenNumList(size, maxVal int) []int {
	var (
		rd  = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
		res = make([]int, 0, size)
	)

	for range size {
		num := rd.IntN(maxVal)
		res = append(res, num)
	}
	return res
}
