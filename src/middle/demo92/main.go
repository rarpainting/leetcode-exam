/*
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 *
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ln := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(&ln)

	newhead := reverseBetween(&ln, 2, 4)

	fmt.Println(newhead)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	strList := []string{}
	ln2 := ln
	for ; ln2 != nil; ln2 = ln2.Next {
		strList = append(strList, strconv.Itoa(ln2.Val))
	}
	return strings.Join(strList, " --> ")
}

// KEY:
// 1. 变换 n-m 个节点
// 2. 从 1 开始
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	dummy := ListNode{
		Next: head,
	}

	pre := &dummy
	for i := 1; i < m; pre, i = pre.Next, i+1 {
	}

	// pre 的 next 节点作为变换的尾节点
	tail := pre.Next
	// 替换 tail 和 tail.next
	for i := 0; i < n-m; i++ {
		preNext := pre.Next
		tail.Next, pre.Next = tail.Next.Next, tail.Next
		pre.Next.Next = preNext

		// tailNext := tail.Next
		// tail.Next = tailNext.Next
		// tailNext.Next = pre.Next
		// pre.Next = tailNext
	}

	return dummy.Next
}
