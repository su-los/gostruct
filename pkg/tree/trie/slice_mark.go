// 实现延迟删除的 slice
package trie

type sliceWithMark[T any] struct {
	data    []T
	delIdx  []int  // 删除的元素的索引
	marks   []bool // 标记是否删除
	deleted int    // 删除的元素个数
}

func newSliceWithMark[T any](size int) *sliceWithMark[T] {
	return &sliceWithMark[T]{
		data:  make([]T, size),
		marks: make([]bool, size),
	}
}

// Len 返回未删除的元素个数
func (s *sliceWithMark[T]) Len() int {
	return len(s.data) - s.deleted
}

// Get 获取第 i 个元素
func (s *sliceWithMark[T]) Get(i int) (T, bool) {
	if s.marks[i] {
		var zero T
		return zero, false
	}
	return s.data[i], true
}

// Delete 删除第 i 个元素
func (s *sliceWithMark[T]) Delete(i int) {
	if s.marks[i] {
		return
	}

	s.delIdx = append(s.delIdx, i)
	s.marks[i] = true
	s.deleted++
}

// Put 添加元素
//
// 返回索引
func (s *sliceWithMark[T]) Put(v T) int {
	s.compact()

	if s.deleted > 0 {
		idx := s.delIdx[len(s.delIdx)-1]
		s.delIdx = s.delIdx[:len(s.delIdx)-1]
		s.data[idx] = v
		return idx
	} else {
		s.data = append(s.data, v)
		s.marks = append(s.marks, false)
		return len(s.data) - 1
	}
}

// Range 遍历未删除的元素
//
// i 为元素的索引
// v 为元素的值
// 返回 false 时停止遍历
func (s *sliceWithMark[T]) Range(f func(i int, v T) bool) {
	for i, v := range s.data {
		if !s.marks[i] {
			if !f(i, v) {
				break
			}
		}
	}
}

// compact 压缩空间
func (s *sliceWithMark[T]) compact() {
	const capSacle = 125 // 容量缩容的比例
	const threshold = 65 // 删除的数量占总长度的比例 >= 65% 时进行缩容

	if s.deleted*100 < threshold*len(s.data) {
		return
	}
	newSize := int((len(s.data) - s.deleted) * capSacle / 100)
	newData := make([]T, 0, newSize)
	newMarks := make([]bool, newSize)
	for i, v := range s.data {
		if !s.marks[i] {
			newData = append(newData, v)
			newMarks = append(newMarks, false)
		}
	}
	s.data = newData
	s.marks = newMarks
	s.deleted = 0
}
