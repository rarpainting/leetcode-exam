/*
 * 我们有一系列公交路线。每一条路线 routes[i] 上都有一辆公交车在上面循环行驶。例如，有一条路线 routes[0] = [1, 5,
 * 7]，表示第一辆 (下标为0) 公交车会一直按照 1->5->7->1->5->7->1->... 的车站路线行驶。
 *
 * 假设我们从 S 车站开始（初始时不在公交车上），要去往 T 站。 期间仅可乘坐公交车，求出最少乘坐的公交车数量。返回 -1 表示不可能到达终点车站。
 *
 *
 * 示例:
 * 输入:
 * routes = [[1, 2, 7], [3, 6, 7]]
 * S = 1
 * T = 6
 * 输出: 2
 * 解释:
 * 最优策略是先乘坐第一辆公交车到达车站 7, 然后换乘第二辆公交车到车站 6。
 *
 *
 * 说明:
 *
 *
 * 1 <= len(routes) <= 500.
 * 1 <= len(routes[i]) <= 500.
 * 0 <= routes[i][j] < 10 ^ 6.
 *
 *
 */
package main

func main() {

}

type Bus struct {
	BusId   int
	BusStop []int
}

// Dijkstra
// 由位置 A 到位置 B 的路途中
// 求出最少乘坐的公交车数量, 而不是最短的路径...
// routes[i] -- 公交车索引
// route[i][j] -- 该公交车 i 可到的站台号
func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}

	// 能到达该站点的公交车索引
	// 即 站台 busStop[i] 为 结点, 公交车 routes[j] 为线的最短路径
	busStop := make(map[int][]Bus)
	for i, v := range routes {
		for si, sv := range v {
			busStop[sv] = append(busStop[sv], Bus{
				BusId:   i,
				BusStop: append(v[:si], v[si+1:]...),
			})
		}
	}

	// 一开始就有序, 我排序你 mua
	stack := []int{S}                // [站点]栈
	visitedSet := make(map[int]bool) // 已遍历的[站点]
	res := 0

	for {
		if len(stack) == 0 {
			return -1
		}

		res++

		for i := 0; i < len(stack); i++ {
			// pop
			t := stack[0]
			stack = stack[1:]

			for _, v := range busStop[t] {
				for _, stopv := range routes[v.BusId] {
					if _, ok := visitedSet[stopv]; ok {
						continue
					}
					visitedSet[stopv] = true
					if stopv == T {
						return res
					}
					stack = append(stack, stopv)
				}
			}
		}
	}

	// return -1
}

// 排序
func sort(src []int) []int {
	if len(src) < 2 {
		return src
	} else if len(src) < 3 {
		if src[0] > src[1] {
			return []int{src[1], src[0]}
		}
		return src
	}

	dist, o := []int{}, src[0]
	l, g := []int{}, []int{}
	for _, v := range src[1:] {
		if v < o {
			l = append(l, v)
		} else {
			g = append(g, v)
		}
	}
	dist = append(sort(l), o)
	return append(dist, sort(g)...)
}
