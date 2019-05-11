package main

import (
	"flag"
	"fmt"
)

var (
	target = flag.Int("target", 100, "")
)

func main() {
	flag.Parse()

	fmt.Println(racecar(*target))
}

var (
	INT_MAX = 10000
)

/*
DP

1. 对于目标 i 步数, 有两种可能
  1.1. 到达 i 之前回头两次
	1.1.1. 正向次数为 cnt1 , 反向次数为 cnt2
	1.1.2. dp[i]: i 的最优次数; j=正向步数; k=反向步数
  2. 超过 i 之后回头
*/
func racecar(target int) int {
	dp := make([]int, target)
	for i := 1; i < target; i++ {
		dp[i] = INT_MAX
		j, cnt1 := 1, 1
		for ; j < i; j, cnt1 = 1<<uint(cnt1+1)-1, cnt1+1 {
			for k, cnt2 := 0, 0; k < j; k, cnt2 = 1<<uint(cnt2+1)-1, cnt2+1 {
				// 到达 i 之前, 往返两次
				// 两次 +1 是反向操作的代价
				dp[i] = min(dp[i], cnt1+1+cnt2+1+dp[i-(j-k)])
			}
		}

		// 超过 i 后回头
		// res: 返回后剩下的步数
		res := 0
		if i != j {
			res = 1 + dp[j-i]
		}
		dp[i] = min(dp[i], cnt1+res)
	}
	return dp[target]
}

func commandA(pos, spd int) (nextPos, nextSpd int) {
	return pos + spd, spd * 2
}

func commandR(pos, spd int) (nextPos, nextSpd int) {
	if spd > 0 {
		return pos, -1
	} else {
		return pos, 1
	}
}

func resolve(src uint32) (pos, spd uint16) {
	return uint16(src >> 16), uint16(src)
}

func combine(pos, spd uint16) (res uint32) {
	return uint32(pos)<<16 | uint32(spd)
}

func min(a int, b int) int {
	if b > a {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
