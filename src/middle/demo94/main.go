package main

import (
	"fmt"
)

func main() {
	tn := TreeNode{
		Val: 34,
		Left: &TreeNode{
			Val: 22,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 12,
		},
	}

	fmt.Println(inorderTraversal(&tn))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	cur := root

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// pop
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}
