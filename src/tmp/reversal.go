/*
智商有问题的记得要构建新链表
原链表翻转好痛苦
*/
package tmp

import (
	"bytes"
	"fmt"
)

type Node struct {
	Val  interface{}
	Next *Node
}

func reversal(first *Node) *Node {
	if first == nil {
		panic(first)
	}
	res := &Node{Val: -1}
	res.Next = first

	for prv := first.Next; prv != nil; {
		tmp := prv.Next
		prv.Next = res.Next
		res.Next = prv
		prv = tmp
	}

	first.Next = nil

	return res.Next
}

func (n *Node) String() string {
	buf := bytes.NewBuffer([]byte{})
	for n != nil {
		buf.Write([]byte(fmt.Sprintf("%v ->", n.Val)))
		n = n.Next
	}

	return buf.String()
}

func main() {
	first := &Node{
		Val: 1,
	}

	fmt.Println(first)

	new := reversal(first)

	fmt.Println(new)
}
