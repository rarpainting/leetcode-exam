package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	cur := head
	for i := 0; i < n; i++ {
		cur = cur.Next
	}

	// 到结尾了
	if cur == nil {
		return head.Next
	}

	pre := head
	for cur.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}

	cur.Next = cur.Next.Next

	return head
}
