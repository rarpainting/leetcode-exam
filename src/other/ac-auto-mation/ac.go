/*
多模式串匹配
ac 自动机
*/
package auautomation

const (
	CountOfChildren = 26
)

// Fail point 依赖于这个基础: Parent.Fail.Depth < Parent.Depth 恒成立
// 因此, Prev.Fail.Depth < Parent.Depth 同样恒成立
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

func NewAC() *ACAuto {
	return &ACAuto{Root: &Node{}}
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

func (t *ACAuto) InsertWithFail(word []byte, num int) {
	t.Insert(word, num)

	t.GenerateFail()
}

// 构建失配指针
//
// 对于一个节点 C, 标识字符 a
//
// 顺着 C 的父亲节点(parent)的失配指针(fail)走, 走到第一个儿子也是 a 的节点 T=parent.Children[i] , 那么 C 的失配指针就指向 T 的标识 a 的儿子节点
//
// 如果找不到这个节点，那么失配指针指向 Root
func (t *ACAuto) GenerateFail() {
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

func (t *ACAuto) Search(query string) (ret [][2]int) {
	lenQue, prv := len(query), t.Root
	for i := 0; i < lenQue; i++ {
		idx := query[i] - 'a'
		for prv.Children[idx] == nil && prv != t.Root {
			// Root 的 Fail 是它自身
			prv = prv.Fail
		}

		if prv.Children[idx] == nil {
			// 匹配失败
			continue
		}

		prv = prv.Children[idx]

		for tl := prv; tl != t.Root; tl = tl.Fail {
			if tl.Point != -1 {
				ret = append(ret, [2]int{tl.Point, int(idx)})
			}
		}
	}

	return
}
