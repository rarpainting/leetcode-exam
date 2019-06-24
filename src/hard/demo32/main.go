/*
最长有效括号
*/
package main

import (
	"flag"
	"fmt"
)

var (
	str = flag.String("str", "", "")
)

func main() {
	flag.Parse()

	fmt.Println("maxlen is", longestValidparentheses(*str))
}

// 最长有效括号
func longestValidparentheses(s string) (max int) {
	sbyte, start := []byte(s), 0
	stack := make([]int, 0)

	for i := range sbyte {
		lstack := len(stack)
		if sbyte[i] == byte(')') {
			if lstack == 0 {
				// 下一次
				start = i + 1
				continue
			}
			// pop
			// pre := stack[lstack-1]
			stack = stack[:lstack-1]

			if len(stack) == 0 {
				// 只修改 max
				// 包含 '(' ')' 两个字符
				if max < i-start+1 {
					max = i - start + 1
				}
			} else { // ')' 符号 && lstack != 0 && lstack > 1
				lstack = len(stack)

				// pop
				pre := stack[lstack-1]
				// stack = stack[:lstack-1]

				/*
					i-stack[lstack-1] 是什么 ?
					i - stack[lstack-1] 是 i 到倒数第二个 '(' 的 [距离]
					这个是 解决 '()(()' 的关键 ?
				*/
				if max < i-pre {
					max = i - pre
				}
			}
		} else {
			stack = append(stack, i)
		}
	}

	return
}

func longestValidParenthesesDP(s string) (maxLen int) {
	n := len(s)
	dp := make([]int, n) // 第 i 个字符前的最长括号串
	for i := 2; i <= n; i++ {
		j := i - 2 - dp[i-1]
		if s[i-1] == '(' || j < 0 || s[j] == ')' {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1] + 2 + dp[j]
			maxLen = max(maxLen, dp[i])
		}
	}

	return
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
