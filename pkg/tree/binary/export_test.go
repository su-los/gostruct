package binary

import (
	"math/rand"
	"time"
)

// GenBFSList 生成一个随机的层序遍历的序列
func GenBFSList(size, maxVal int) []int {
	var (
		rd  = rand.New(rand.NewSource(int64(time.Now().UnixNano())))
		len = rd.Intn(size) + 1 // 序列长度
		res = make([]int, 0, len)
	)

	// 生成不重复的数字序列
	uniq := make(map[int]struct{}, 0)
	for i := 0; i < len; {
		num := rd.Intn(maxVal)
		if num == 0 {
			continue
		}

		_, ok := uniq[num]
		if ok {
			continue
		}

		res = append(res, num)
		uniq[num] = struct{}{}
		i++
	}
	return res
}
