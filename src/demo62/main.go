package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(5, 4))
}

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
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

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	} else if len(obstacleGrid[0]) == 0 {
		return 0
	} else if obstacleGrid[0][0] == 1 {
		return 0
	} /* else if len(obstacleGrid) == 1 || len(obstacleGrid[0]) == 1 {
		return 1
	} */

	lenT, lenS := len(obstacleGrid)+1, len(obstacleGrid[0])+1

	newObs := make([][]int, lenT) // m + 1
	for i := range newObs {
		newObs[i] = make([]int, lenS) // n + 1
	}
	newObs[0][1] = 1

	for i := 1; i < lenT; i++ {
		for j := 1; j < lenS; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			newObs[i][j] = newObs[i-1][j] + newObs[i][j-1]
		}
	}
	return newObs[lenT-1][lenS-1]
}
