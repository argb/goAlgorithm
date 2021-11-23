package dijkstra

import (
	"testing"
)

const INF = -1

func TestDijkstra(t *testing.T)  {
	vertexes := []int{1, 2, 3, 4, 5}
	var matrix = [][9]int{
		/*A*//*B*//*C*//*D*//*E*//*F*//*G*/
		/*A*/ {   0,  12, INF, INF, INF,  16,  14},
		/*B*/ {  12,   0,  10, INF, INF,   7, INF},
		/*C*/ { INF,  10,   0,   3,   5,   6, INF},
		/*D*/ { INF, INF,   3,   0,   4, INF, INF},
		/*E*/ { INF, INF,   5,   4,   0,   2,   8},
		/*F*/ {  16,   7,   6, INF,   2,   0,   9},
		/*G*/ {  14, INF, INF, INF,   8,   9,   0}}

	var prev []int
	var dist []int
	prev = make([]int, VertexNum)
	dist = make([]int, VertexNum)

	g := new(Graph)
	g.InitTable(vertexes, len(vertexes), matrix)
	g.Dijkstra(0, prev, dist)
}
