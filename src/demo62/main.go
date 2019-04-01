package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(5, 4))
}

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

func uniquePaths2(m int, n int) int {
	num, denom := int64(1), int64(1)
	small := 0
	if m > n {
		small = n
	} else {
		small = m
	}

	for i := 1; i <= small-1; i++ {
		num *= int64(m + n - 1 - i)
		denom *= int64(i)
	}

	return int(num / denom)
}

func helper(row int, col int, m int, n int) int {
	if row == m && col == n {
		return 1
	} else if row > m || col > n {
		return 0
	}
	return helper(row+1, col, m, n) + helper(row, col+1, m, n)
}
