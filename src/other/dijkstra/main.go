/*
单源最短路径
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
