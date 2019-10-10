/*
单源最短路径
思路: BFS
管理一个出度表
从 单源 出发 到各个节点, 更新各个节点的数 value 以及造成该值的前一个节点 previousNode
*/
package main

import (
	"fmt"
)

func main() {

}

type Graph struct {
	Vertices map[string][]Vertex
}

func (g *Graph) AddVertex(char string, vs []Vertex) {
	g.Vertices[char] = append(g.Vertices[char], vs...)
}

type Vertex struct {
	ID       string
	Distance int
}

func (vex *Vertex) String() string {
	return fmt.Sprintf(`Vertex [id=%s distance=%d]`, vex.ID, vex.Distance)
}

func compire(vex1, vex2 *Vertex) int {
	if vex1.Distance > vex2.Distance {
		return 1
	} else if vex1.Distance < vex2.Distance {
		return -1
	}
	return 0
}
