/*
最近公共祖先

倍增法
1. 两节点深度相同

1.1. 预处理
1.1.1. 存储一棵树(邻接表法)
用邻接表存储一棵树(无特征树), 并用 from[] 数组记录各结点的 **父节点**, 其中没有父节点的就是 root .
1.1.2. 获取树各结点的上的深度(dfs或bfs)
1.1.3. 获取 2 次幂祖先的结点, 用 parents[maxn][20] 数组存储, 倍增法关键
1.1.4. 用倍增法查询 Lca

1.2. 查询

2. 两节点深度不同 -- 更深的节点上浮到两节点深度相同, 再通过 1 的方法解
*/
package main

import ()

type Tree struct {
	Root *Node
}

type Node struct {
	Val interface{}
	Son []*Node
}

func main() {

}
