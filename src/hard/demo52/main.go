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

func main() {

}

func totalNQueens(n int) int {
	count := 0

	for i := 0; i < n; i++ {
		if isVaild(i, n) {
			count++
		}
	}

	return count
}

func isVaild(side int, n int) bool {
	return true
}
