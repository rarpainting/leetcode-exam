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
				// 结束
				start = i + 1
				continue
			}
			// pop
			// pre := stack[lstack-1]
			stack = stack[:lstack-1]
			if len(stack) == 0 {
				if max < i-start+1 {
					max = i - start + 1
				}
			} else {
				lstack = len(stack)
				pre := stack[lstack-1]
				stack = stack[:lstack-1]
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
