/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

package tmp

// @lc code=start
func trap(height []int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	length := len(height)
	left, right, result := 0, length-1, 0

	// 两端往中间靠拢
	// 两边分别寻找沟位置, 找到后 result+=tmp
	// so on...
	for left < right {
		mn := min(height[left], height[right])
		if mn == height[left] {
			left++
			for left < right && height[left] < mn {
				// 存在沟, 则计算沟的数值
				result += mn - height[left]
				left++
			}
		} else {
			right--
			for left < right && height[right] < mn {
				// 存在沟, 则计算沟的数值
				result += mn - height[right]
				right--
			}
		}
	}

	return result
}

// @lc code=end

// DP
func trap1(height []int) int {
	return 0
}

// stack
func trap2(height []int) int {
	return 0
}
