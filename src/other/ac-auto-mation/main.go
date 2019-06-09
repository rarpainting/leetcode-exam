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
	Father   *Node
	Point    int // 用于中间节点标志
	Children [CountOfChildren]*Node
	// 失败节点, 目的是在失败匹配后, 跳转到继续匹配的位置
	Fail *Node
}

type ACAuto struct {
	Root *Node
}

func NewAC() {

}

func (t *ACAuto) Insert(word []byte, num int) {
	prv := t.Root

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

func (t *ACAuto) InsertWithFailPoint(word []byte, num int) {
	t.Insert(word, num)

	t.GenerateFailPoint()
}

// 构建失配指针
// 对于一个节点 C, 标识字符 a,
// 顺着 C 的父亲节点的失配指针走, 走到第一个有儿子也是 a 的节点 T, 那么 C 的失配指针就指向 T 的标识 a 的儿子节点
// 如果找不到这个节点，那么失配指针指向 Root
func (t *ACAuto) GenerateFailPoint() {
	root := t.Root
	root.Fail = root
	stack := []*Node{}
	for _, v := range root.Children {
		if v == nil {
			continue
		}
		v.Fail = root
		stack = append(stack, v)
	}

	for {
		if len(stack) == 0 {
			return
		}

		lenStack := len(stack)
		for i := 0; i < lenStack; i++ {
			// pop parent node
			pNode := stack[0]
			stack = stack[1:]

			// 只操作 children
		loop:
			for si, sv := range pNode.Children {
				if sv == nil {
					continue
				}

				// parent->FailPoint->Children[si]
				var canFail *Node
				for failNode := pNode.Fail; ; failNode, canFail = failNode.Fail, nil {
					if failNode == nil {
						// 找不到可行的 failNode
						canFail = root
					} else if failNode.Children[si] != nil {
						canFail = failNode.Children[si]
					}

					if canFail != nil {
						sv.Fail = canFail
						stack = append(stack, sv)
						continue loop
					}
				}
			}
		}
	}
}
