package main

import (
	"algorithm/graph/dijkstra"
	"math"
)

const INF = math.MaxInt32

func main() {
	vertexNum := dijkstra.VertexNum
	vertexes := []string{"A", "B", "C", "D", "E", "F", "G"}
	var matrix = [][5]int{
				/*A*//*B*//*C*//*D*//*E*/
		/*A*/ {   0,  8,   1, 2, INF},
		/*B*/ {   8,  0, INF, 3, INF},
		/*C*/ {   1,  3,  0,  2,   3},
		/*D*/ {   2, INF, 2,  0,   3},
		/*E*/ { INF, INF, 3,  3,   0}}

	var dist []int
	dist = make([]int, vertexNum)

	g := new(dijkstra.Graph)
	g.InitTable(vertexes, len(vertexes), matrix)
	//g.Matrix = matrix
	g.Dijkstra(0, dist)
}
