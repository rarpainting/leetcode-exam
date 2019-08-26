/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 */

// Definition for a binary tree node.
package tmp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}

	for len(stack) != 0 {
		i, l, se := 0, len(stack), false
		for ; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			if se != true {
				res = append(res, node.Val)
				se = true
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return res
}
