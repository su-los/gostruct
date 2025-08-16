package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrieTreeASCII(t *testing.T) {
	trie := NewTrieTree()
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hello")

	require.ElementsMatch(t, []string{"hello"}, trie.Search("he"))
	require.ElementsMatch(t, []string{"world"}, trie.Search("worl"))
	require.ElementsMatch(t, []string{"world"}, trie.Search("world"))
	require.Empty(t, trie.Search("worldd"))
}

func TestTrieTreeChinese(t *testing.T) {
	trie := NewTrieTree()
	trie.Insert("你好")
	trie.Insert("世界")

	require.ElementsMatch(t, []string{"你好"}, trie.Search("你"))
	require.ElementsMatch(t, []string{"世界"}, trie.Search("世"))
	require.ElementsMatch(t, []string{"世界"}, trie.Search("世界"))
	require.Empty(t, trie.Search("世界d"))
}

func TestTrieTreeMix(t *testing.T) {
	trie := NewTrieTree()
	trie.Insert("你好, Hello")
	trie.Insert("世界")
	trie.Insert("世界, world")
	trie.Insert("hello")
	trie.Insert("world")

	require.ElementsMatch(t, []string{"你好, Hello"}, trie.Search("你好,"))
	require.ElementsMatch(t, []string{"世界, world"}, trie.Search("世界, w"))
	require.ElementsMatch(t, []string{"world"}, trie.Search("worl"))
}

func TestDelete(t *testing.T) {
	trie := NewTrieTree()
	trie.Insert("你好, Hello")
	trie.Insert("世界")
	trie.Insert("世界, world")
	trie.Insert("hello")
	trie.Insert("world")
	require.ElementsMatch(t, []string{"你好, Hello"}, trie.Search("你好,"))

	trie.Delete("你好, Hello")
	require.Empty(t, trie.Search("你好,"))

	// 删除不存在的 key
	trie.Delete("不存在")

	trie.Insert("ho")
	trie.Insert("how")
	require.ElementsMatch(t, []string{"ho", "how"}, trie.Search("ho"))

	trie.Delete("ho")
	require.ElementsMatch(t, []string{"how"}, trie.Search("ho"))
}

func TestTrieTree(t *testing.T) {
	initGlobalTrie(10, 10000)

	for i := range 10000 {
		str := gloStrs[i]
		res := gloTrie.Search(str)
		require.Contains(t, res, str)
	}

	for i := range 10000 {
		str := gloStrs[i]
		gloTrie.Delete(str)
	}
}
