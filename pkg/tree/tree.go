package tree

type Tree interface {
	Insert(val int)
	Delete(val int)
	Find(val int) bool
	InOrder() []int
	PreOrder() []int
	PostOrder() []int
	LevelOrder() []int
	Min() int
	Max() int
}
