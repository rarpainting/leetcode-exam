package main

func main() {
}

func NewDoubelArrayTrie() *DoubleArrayTrie {
	return &DoubleArrayTrie{}
}

type DoubleArrayTrie struct {
	Check []rune
	Base  []rune

	Node
}

type Node struct {
	Code  int // 编码
	Depth int // 树深
	Left  int // 范围
	Right int
}
