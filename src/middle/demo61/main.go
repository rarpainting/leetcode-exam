package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln := ListNode{0, nil}
	ln.Next = &ListNode{1, nil}
	ln.Next.Next = &ListNode{2, nil}

	rotateRight(&ln, 4)
}

func rotateRight(head *ListNode, k int) *ListNode {
	// TODO: 排错
	if head == nil {
		return nil
	}

	tail, n := head, 0
	for ; tail != nil; tail, n = tail.Next, n+1 {
	}

	k %= n // 真正的位置
	fast, slow := head, head

	for i := 0; i < k; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}

	// 刚刚好是最后一个
	if fast == nil {
		return head
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	fast.Next, fast, slow.Next = head, slow.Next, nil

	return fast
}

// k --> [0, n-k%n], (n-k%n, n]
func rotateRight2(head *ListNode, k int) *ListNode {
	// TODO: 排错
	if head == nil {
		return nil
	}

	cur, n := head, 1
	for ; cur.Next != nil; cur, n = cur.Next, n+1 {
	}

	cur.Next = head // 首尾相接
	// m := k%n - 1
	tk := k % n
	if tk == 0 {
		tk = n
	} else {
		tk = tk - 1
	}
	// m := n - k%n // 结果的头节点的前一个节点 ??

	for i := 0; i < tk; cur, i = cur.Next, i+1 {
	}

	newHead := cur.Next
	cur.Next = nil
	return newHead
}
