/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

package tmp

// 提示: 如果中间存在某点无法通过, 那么当前假设的起点到这个点都不能作为起点
func canCompleteCircuit(gas []int, cost []int) int {
	total, sum, start := 0, 0, 0
	for i := range gas {
		remain := gas[i] - cost[i]
		total += remain
		sum += remain

		// 出现无法继续的点
		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}

	if total < 0 {
		return -1
	} else {
		return start
	}
}

// ??
func canCompleteCircuit2(gas []int, cost []int) int {
	total, mx, start := 0, -1, 0
	for i := len(gas) - 1; i > 0; i-- {
		remain := gas[i] - cost[i]
		total += remain

		// 出现可能作为起点的点
		if total > mx {
			start = i
			mx = total
		}
	}

	if total < 0 {
		return -1
	} else {
		return start
	}
}
