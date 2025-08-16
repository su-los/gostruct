// trie 树的实现
package trie

// TrieTree 前缀树 / 字典树
type TrieTree struct {
	root *trieNode
	// key 关键字，val 在全局数组中的索引
	// 用于快速判断某个单词是否已经插入过
	strs map[string]struct{}
}

// NewTrieTree 构建前缀树
//
// 根节点为 '/'，不存储任何数据
// size 为预分配的大小，用于减少 map 的扩容次数
func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: newTrieNode('/'),
		strs: make(map[string]struct{}),
	}
}

// Insert 插入单词
func (t *TrieTree) Insert(word string) {
	if t.root == nil {
		t.root = newTrieNode('/')
	}

	// 检查是否已经插入过
	if _, ok := t.strs[word]; ok {
		// 已经插入过，直接返回
		return
	}

	// 存储到 map 中，用于快速查找
	t.strs[word] = struct{}{}
	rwd := []rune(word)

	// 注意：遍历时的索引 i, char := range item.rwd
	// 就是表示当前节点匹配的是 rwd 的第几个字符，有需求可以记录
	cur := t.root
	for _, char := range rwd {
		child := cur.getChild(char)
		if child == nil {
			// 不存在则插入
			cur = cur.addChild(char)
			continue
		}

		cur = child
	}
	cur.isEnd = true
}

// Search 搜索单词
//
// 输出包含 word 前缀的所有单词
func (t *TrieTree) Search(word string) []string {
	if t.root == nil {
		return nil
	}

	rwd := []rune(word)
	cur := t.root
	for _, char := range rwd {
		child := cur.getChild(char)
		if child == nil {
			return nil
		}

		cur = child
	}

	// 遍历子节点，得到所有以 word 为前缀的字符串
	// 最后一个字符以 cur.char 结尾，所以这里传入 word[:len(word)-1]
	// 根据 cur.isEnd 判断是否需要将 word 追加进结果
	return t.preOrder(cur, rwd[:len(rwd)-1])
}

// preOrder 前序遍历获取每条路径上的字符串
//
// 递归实现
// func (t *TrieTree) preOrderRe(node *trieNode, word []rune) []string {
// 	if node == nil {
// 		return nil
// 	}

// 	var (
// 		buf = make([]rune, 0, len(word)+1)
// 		res = make([]string, 0)
// 	)
// 	buf = append(buf, word...)
// 	buf = append(buf, node.char)
// 	if node.isEnd {
// 		// 是单词结尾，直接返回
// 		res = append(res, string(buf))
// 	}

// 	for _, child := range node.children {
// 		res = append(res, t.preOrderRe(child, buf)...)
// 	}
// 	return res
// }

// preOrder 前序遍历获取每条路径上的字符串
//
// 非递归实现需要使用栈来存储节点
func (t *TrieTree) preOrder(node *trieNode, word []rune) []string {
	if node == nil {
		return nil
	}

	var (
		buf   = make([]rune, 0, len(word)+1)
		stack = []*trieNode{node}
		path  = make([]*trieNode, 0) // 存储路径
		res   = make([]string, 0)
	)
	buf = append(buf, word...)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		path = append(path, top)
		buf = append(buf, top.char)
		if top.isEnd {
			// 是单词结尾
			res = append(res, string(buf))
		}

		// 叶子节点
		if len(top.children) == 0 {
			// 栈顶分支
			if len(stack) == 0 {
				// 栈空了，top 也为叶子节点，说明所有路径已经遍历完毕
				break
			}
			pre := stack[len(stack)-1]
			// 回溯到当前栈顶的父节点
			for len(path) > 0 && path[len(path)-1].children[pre.char] != pre {
				path = path[:len(path)-1]
				buf = buf[:len(buf)-1]
			}
			continue
		}

		for _, child := range top.children {
			stack = append(stack, child)
		}
	}
	return res
}

// Delete 删除单词
func (t *TrieTree) Delete(word string) {
	if t.root == nil {
		return
	}

	// 检查删除的 word 是否存在
	_, ok := t.strs[word]
	if !ok {
		return
	}

	cur := t.root
	rwd := []rune(word)
	delStack := make([]*trieNode, 0, len(rwd))
	for _, char := range rwd {
		child := cur.getChild(char)
		if child == nil {
			return
		}
		delStack = append(delStack, child)
		cur = child
	}
	// 表示该节点不再是单词 word 的末尾
	cur.isEnd = false

	// 自底向上删除链路上的节点
	var delChild *trieNode
	for len(delStack) > 0 {
		top := delStack[len(delStack)-1]
		delStack = delStack[:len(delStack)-1]

		// 删除需要删除的子节点
		if delChild != nil && delChild != top {
			top.deleteChild(delChild.char)
			freeTrieNode(delChild)
		}

		if !top.isEnd && len(top.children) == 0 {
			// 没有子节点，也没有以该字符结尾的单词，可以删除
			delChild = top
		} else {
			break
		}
	}
}
