/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */
// 关键词: 拓扑排序 BFS
package tmp

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) < 1 {
		return true
	}
	if len(prerequisites) == 1 {
		if prerequisites[0][0] == prerequisites[0][1] {
			return false
		} else {
			return true
		}
	}
	// prerequisites[i][0] // 目标
	// prerequisites[i][1] // 先修
	// 入度 & init
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
		return false
	}

	return true
}
