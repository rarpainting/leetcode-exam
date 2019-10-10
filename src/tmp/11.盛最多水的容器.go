/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package tmp

// @lc code=start
func maxArea(height []int) int {
	i, j, result := 0, len(height)-1, 0
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	min := func(i, j int) int {
		if i < j {
			return i
		} else {
			return j
		}
	}

	for i < j {
		result = max((j-i)*min(height[i], height[j]), result)

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return result
}

// @lc code=end
