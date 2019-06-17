package main

func main() {
}

func NewDoubelArrayTrie() *DoubleArrayTrie {
	return &DoubleArrayTrie{
		Base:  append([]rune{}, 1),
		Check: append([]rune{}, 0),

		// 父节点无数据, 同时最后的节点必须挂一个空的节点
		Node: Node{
			Depth: 0,
			Children: map[rune]Node{
				0: Node{
					Depth: 1,
				},
			},
		},
	}
}

func (dat *DoubleArrayTrie) Build(keys [][]rune) error {
	if len(keys) == 0 {
		return nil
	}

	return nil
}

type DoubleArrayTrie struct {
	Base  []rune
	Check []rune

	Node
}

type Node struct {
	// Code  int    // 编码
	Depth int    // 树深
	Range [2]int // 范围
	// Left  int // 范围
	// Right int

	Children map[rune]Node
}
