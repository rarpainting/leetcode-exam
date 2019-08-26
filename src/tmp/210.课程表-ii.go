/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

package tmp

func findOrder(numCourses int, prerequisites [][]int) []int {
	if len(prerequisites) < 1 {
		if numCourses != 0 {
			res := []int{}
			for i := 0; i < numCourses; i++ {
				res = append(res, i)
			}
			return res
		}
	}
	if len(prerequisites) == 1 {
		if prerequisites[0][0] == prerequisites[0][1] {
			return []int{}
		} else {
			res := []int{prerequisites[0][1], prerequisites[0][0]}
			if numCourses > 2 {
				hasOK := make([]bool, numCourses)
				hasOK[prerequisites[0][1]] = true
				hasOK[prerequisites[0][0]] = true
				for i := range hasOK {
					if hasOK[i] != true {
						res = append(res, i)
					}
				}
			}
			return res
		}
	}
	// prerequisites[i][0] // 目标
	// prerequisites[i][1] // 先修
	// 入度 & init
	res := []int{}
	enterCount := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		enterCount[i] = 0
	}
	// 无前驱的点 index
	stack := []int{}
	// 构建入度数组
	for i := range prerequisites {
		enterCount[prerequisites[i][0]]++
	}

	// 第一次入栈
	for k := range enterCount {
		if enterCount[k] == 0 {
			stack = append(stack, k)
		}
	}

	for len(stack) != 0 {
		si, sl := 0, len(stack)
		// newpq := append([][]int{}, prerequisites...)

		for ; si < sl; si++ {
			// pop
			ec := stack[0]
			stack = stack[1:]

			delete(enterCount, ec)
			res = append(res, ec)

			lenpq := len(prerequisites)
			for pi := 0; pi < lenpq; pi++ {
				if prerequisites[pi][1] == ec {
					// 从 入度 map 和 prerequisites 中删除 ec 相关的内容
					enterCount[prerequisites[pi][0]]--

					if enterCount[prerequisites[pi][0]] == 0 {
						stack = append(stack, prerequisites[pi][0])
					}

					prerequisites = append(prerequisites[:pi], prerequisites[pi+1:]...)
					pi--
				}
				lenpq = len(prerequisites)
			}
		}
	}

	if len(enterCount) != 0 {
		return []int{}
	}

	return res
}
