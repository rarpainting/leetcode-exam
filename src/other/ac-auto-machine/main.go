/*
多模式串匹配
ac 自动机
*/
package main

const (
	CountOfChildren = 26
)

func main() {

}

type Node struct {
	// Value    rune
	Point    int // 用于中间节点标志
	Children [CountOfChildren]*Node
	// 失败节点, 目的是在失败匹配后, 跳转到继续匹配的位置
	Fail *Node
}

type Trie struct {
	Root *Node
}

func NewAC() {

}

func (t *Trie) Insert(word []byte, num int) {
	prv := t.Root
	// lenWord := len(word)
	for _, wv := range word {
		node := wv - 0x41
		if prv.Children[node] == nil {
			prv.Children[node] = &Node{
				Point: -1,
			}
		}

		prv = prv.Children[node]
	}

	prv.Point = num
}
