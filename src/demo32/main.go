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
	sbyte, curMax := []byte(s), 0
	stack := make([]byte, 0)
	for i := range sbyte {
		lstack := len(stack)
		if sbyte[i] == byte(')') {
			if lstack == 0 {
				// 结束
				if curMax > max {
					max, curMax = curMax, 0
				}
				continue
			}
			pre := stack[lstack-1]
			stack = stack[:lstack-1]
			if pre != byte('(') {
				// 结束
				if curMax > max {
					max, curMax = curMax, 0
				}
				continue
			} else {
				curMax += 2
			}
		} else {
			stack = append(stack, sbyte[i])
		}
	}
	if curMax > max {
		max = curMax
	}

	return
}
