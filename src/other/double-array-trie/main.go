package main

import (
	"fmt"
)

const (
	_BUF_SIZE  = 16384
	_UNIT_SIZE = 8 // size of int + int
)

func main() {
}

type DoubleArrayTrie struct {
	base  []rune
	check []rune
	used  []bool

	length  []int
	key     [][]rune
	value   []rune
	process int

	Node Node
}

func NewDoubelArrayTrie() *DoubleArrayTrie {
	dat := DoubleArrayTrie{
		base:  make([]rune, 65536*32),
		check: make([]rune, 65536*32),

		// 父节点无数据, 同时最后的节点必须挂一个空的节点
		Node: Node{
			Depth: 0,
		},
	}

	dat.base[0] = 1

	return &dat
}

func (dat *DoubleArrayTrie) Build(keys [][]rune, length []int, value []rune) error {
	if len(keys) == 0 {
		return nil
	}

	dat.key = keys
	dat.length = length
	dat.value = value
	dat.process = 0

	return nil
}

func (dat *DoubleArrayTrie) Fetch(n *Node) (silbings []Node, err error) {
	silbings = []Node{}
	prev := 0
	for i := n.Left; i < n.Right; i++ {
		length := 0
		if len(dat.length) != 0 {
			length = dat.length[i]
		} else {
			length = len(dat.key[i])
		}

		cur := 0

		if length < n.Depth {
			continue
		} else if length > n.Depth {
			cur = int(dat.key[i][n.Depth]) + 1
		}

		if prev > cur {
			return nil, fmt.Errorf("prev > cur")
		}

		if prev < cur || len(silbings) == 0 {
			node := Node{
				Depth: n.Depth,
				Code:  cur,
				Left:  i,
			}

			if len(silbings) != 0 {
				silbings[len(silbings)-1].Right = i
			}
			silbings = append(silbings, node)
		}

		cur = prev
	}

	if len(silbings) != 0 {
		silbings[len(silbings)-1].Right = n.Right
	}

	return silbings, nil
}

type Node struct {
	Code  int // 编码
	Depth int // 树深
	// Range [2]int // 范围
	Left  int // 范围
	Right int

	// Children map[rune]Node
}
