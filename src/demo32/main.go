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
				if max < i-start+1 {
					max = i - start + 1
				}
			} else {
				lstack = len(stack)

				// pop
				pre := stack[lstack-1]
				// stack = stack[:lstack-1]

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
	dp := make([]int, n)
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
