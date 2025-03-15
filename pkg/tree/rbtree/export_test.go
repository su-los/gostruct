package rbtree

import (
	"math/rand"
	"time"
)

// GenBFSList 生成一个随机的层序遍历的序列
func GenBFSList() []int {
	var (
		rd      = rand.New(rand.NewSource(int64(time.Now().UnixNano())))
		len     = rd.Intn(1000) + 1 // 序列长度
		res     = make([]int, 0, len)
		isfirst = true
	)

	// 生成不重复的数字序列
	uniq := make(map[int]struct{}, 0)
	for i := 0; i < len; {
		num := rd.Intn(1000)
		_, ok := uniq[num]
		if num != 0 && ok {
			continue
		}

		// 防止第一个元素为 0
		if isfirst {
			res = append(res, num+1)
			uniq[num+1] = struct{}{}
			isfirst = false
		} else {
			res = append(res, num)
			uniq[num] = struct{}{}
		}
		i++
	}
	return res
}
