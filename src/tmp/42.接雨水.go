/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

package tmp

// @lc code=start
func trap(height []int) int {

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
	stack := make([]int, 0, len(height))
	i, res, n := 0, 0, len(height)

LOOP:
	for i < n {
		// 栈为空 OR 当前(i)高度小于栈顶高度
		if lenOfStack := len(stack); lenOfStack == 0 || height[i] <= height[stack[lenOfStack-1]] {
			stack = append(stack, i)
			i++
		} else {
			st := stack[lenOfStack-1]
			stack = stack[:lenOfStack-1]

			lenOfStack := len(stack)
			if lenOfStack == 0 {
				continue LOOP
			}

			res += (min(height[i], height[stack[lenOfStack-1]]) - height[st]) * (i - stack[lenOfStack-1] - 1)
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
