// 前缀树 / 字典树的节点实现
//
// 难点在于，删除/更新时，如何快速更新各个节点的聚合值
// 这里采用了懒更新的方式，即只有在访问节点时才更新聚合值
package trie

import (
	"sync"
)

var trieNodePool = sync.Pool{
	New: func() any {
		return new(trieNode)
	},
}

// trieNode 前缀树节点定义（支持中文）
type trieNode struct {
	char rune

	children map[rune]*trieNode // 子节点
	isEnd    bool               // 是否是单词结尾
}

// newTrieNode 构建前缀树节点
func newTrieNode(char rune) *trieNode {
	node := trieNodePool.Get().(*trieNode)
	node.char = char
	node.children = make(map[rune]*trieNode)
	node.isEnd = false
	return node
}

// freeTrieNode 释放前缀树节点
func freeTrieNode(node *trieNode) {
	node.char = 0
	node.children = nil
	node.isEnd = false
	trieNodePool.Put(node)
}

// addChild 添加子节点
//
// idx 在全局数组中的索引
func (n *trieNode) addChild(char rune) *trieNode {
	child := newTrieNode(char)
	n.children[char] = child
	return child
}

// getChild 获取子节点
//
// 不存在返回 nil，外层需要判断 nil
func (n *trieNode) getChild(char rune) *trieNode {
	return n.children[char]
}

// deleteChild 删除子节点
func (n *trieNode) deleteChild(char rune) {
	child := n.children[char]
	if child == nil {
		return
	}
	delete(n.children, char)
}
