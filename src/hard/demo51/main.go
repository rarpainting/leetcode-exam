/*
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 *
 *
 * 上图为 8 皇后问题的一种解法。
 *
 * 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
 */
package main

import (
	"flag"
	"fmt"
)

var (
	nQueue = flag.Int("queue", 0, "")
)

func main() {
	flag.Parse()
	res := solveNQueens1(*nQueue)
	for _, v := range res {
		fmt.Println(v)
	}
	fmt.Println(len(res))
}

// DFS
func solveNQueens1(n int) (res [][]string) {
	if n == 0 {
		return
	}

	stack := []int{}

	// push
	stack = append(stack, 0)
	// 初始化栈的首行的列号
	lastN := 0

	for {
		if len(stack) == 0 && lastN == n-1 {
			return
		} else if len(stack) == 0 {
			lastN++
			stack = append(stack, lastN)
		}

		// fmt.Printf("[stack-len]: %d \t[stack]: %d\n", len(stack), stack)

		// 栈顶
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 保证 last 和 前面的都不冲突
		ok := true
		for i, v := range stack {
			if last == v || last-v == len(stack)-i || last-v == i-len(stack) {
				ok = false
				break
			}
		}

		if !ok || (ok && len(stack) == n-1) {
			if ok && len(stack) == n-1 {
				res = append(res, printQueue(n, append(stack, last)))
			} /*  else { // !ok

			} */

			// 如果 last == n-1 , 则选择
			for v := last; ; {
				if v == n-1 {
					if len(stack) == 0 {
						return
					}
					v = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				} else {
					if len(stack) == 0 {
						lastN = v + 1
					}
					stack = append(stack, v+1)
					break
				}
			}
		} else {
			// ok && len(stack) < n-1
			stack = append(stack, last, 0)
		}
	}
}

// 动态规划
// 1. 把每个点归到 竖/左斜/右斜 的三大类中
// 2. 把布尔操作改为 位操作

// shu, pie, na := make([]bool, n), make([]bool, 2*n-2), make([]bool, 2*n-2)
// 检查冲突的标准
// pie 和 na 可以对于 行/列 对称
// shu[col]
// pie[row+col]
// na[n-1-row+col]
func solveNQueens2(n int) (res [][]string) {
	if n == 0 {
		return
	}

	if 2*n-1 > 64 { // 未作位扩展
		return
	}
	shu, pie, na := uint(0), uint(0), uint(0)

	stack := []int{}

	// push
	stack = append(stack, 0)
	// 初始化栈的首行的列号
	lastN := uint(0)

	for {
		if len(stack) == 0 && lastN == uint(n)-1 {
			return
		} else if len(stack) == 0 {
			lastN++
			stack = append(stack, int(lastN))
		}

		// fmt.Printf("[stack-len]: %d \t[stack]: %d\n", len(stack), stack)

		// 栈顶
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 保证 last 和 前面的都不冲突
		ok := true
		if j, k := uint(len(stack)+last), uint(n)-1-uint(len(stack))+uint(last); (shu>>uint(last)|pie>>j|na>>k)&1 == 1 {
			ok = false
		}

		if !ok || (ok && len(stack) == n-1) {
			if ok && len(stack) == n-1 {
				res = append(res, printQueue(n, append(stack, last)))
			} /*  else { // !ok

			} */

			// 如果 last == n-1 , 则选择
			for v := last; ; {
				if v == n-1 {
					if len(stack) == 0 {
						return
					}
					v = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				} else {
					if len(stack) == 0 {
						lastN = uint(v + 1)
					}
					stack = append(stack, v+1)
					break
				}
			}
		} else {
			// ok && len(stack) < n-1
			stack = append(stack, last, 0)
		}
	}

	return
}

func printQueue(n int, sid []int) (res []string) {
	for i := 0; i < n; i++ {
		oneSid := sid[i]
		oneRes := ""
		for j := 0; j < n; j++ {
			if oneSid == j {
				oneRes = oneRes + "Q"
			} else {
				oneRes = oneRes + "."
			}
		}
		res = append(res, oneRes)
	}
	return
}
