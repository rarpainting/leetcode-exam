package main

import (
	"flag"
	"fmt"
)

var (
	num = flag.Int("num", 1, "")
)

func main() {
	flag.Parse()
	out := generateParenthesis(*num)
	fmt.Println(out)
}

// DFS 遍历某深度的树
func generateParenthesis(n int) []string {
	out := generateParenthesisDFS(n, n, "")
	return out
}

func generateParenthesisDFS(left, right int, symbol string) (out []string) {
	// 要求必须先排了左符号, 再排右符号
	// if left > right {
	// 	return
	// }
	if left == 0 && right == 0 { // 在节点尾部
		out = append(out, symbol)
	} else {
		if left > 0 {
			out = append(out, generateParenthesisDFS(left-1, right, symbol+"(")...)
		}
		if right > 0 {
			out = append(out, generateParenthesisDFS(left, right-1, symbol+")")...)
		}
	}
	return
}
