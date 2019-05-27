/*
一般用于 关键词匹配
由于 关键词匹配 的关键词列表一般比需要匹配的字段大得多
因此一般用 trie 保存关键词
*/
package main

import (
	"flag"
	"fmt"
)

var (
	source = flag.String("src", "ab", "")
)

func main() {
	flag.Parse()

	f := NewFilter()

	f.Trie.Insert([]rune("abc"))
	f.Trie.Insert([]rune("ab"))

	fmt.Println(f)
	lists, ok := f.Search([]rune(*source))
	fmt.Printf("[%s] [%v] [%v]", *source, lists, ok)
}

type Node struct {
	// Value    rune
	Mark     bool // 用于中间节点标志
	Children map[rune]Node
}

type Trie struct {
	Root Node // 根节点
}

func (t *Trie) Insert(word []rune) {
	prv := t.Root
	lenWord := len(word)
	for i, wv := range word {
		v, ok := prv.Children[wv]
		if !ok {
			v = Node{
				Children: make(map[rune]Node),
			}
		}
		if i == lenWord-1 {
			v.Mark = true
		}

		// 因为没有用指针, 所以要再复制一遍才能更新 v
		prv.Children[wv] = v

		// 更新 prv
		prv = v
	}
}

func (t *Trie) Search(word []rune) bool {
	prv := t.Root
	// lenWord := len(word)

	for _, wv := range word {
		v, ok := prv.Children[wv]
		if !ok {
			return false
		}
		prv = v
	}

	// lastNode := prv.Children[word[len(word)-1]]
	lastNode := prv
	// 最后一个 || mark==true
	if len(lastNode.Children) == 0 || lastNode.Mark {
		return true
	}
	return false
}

type Filter struct {
	Trie
}

func NewFilter() *Filter {
	return &Filter{
		Trie: Trie{
			Root: Node{
				Children: make(map[rune]Node),
			},
		},
	}
}

// 找出 word 中的关键词
func (f *Filter) Search(word []rune) ([][2]int, bool) {
	lists := [][2]int{}
	lenWord := len(word)
	for i := 0; i < lenWord; i++ {
		for j := i + 1; j < lenWord+1; j++ {
			if ok := f.Trie.Search(word[i:j]); ok {
				list := [2]int{i, j}
				lists = append(lists, list)
			}
		}
	}

	if len(lists) != 0 {
		return lists, true
	}

	return lists, false
}
