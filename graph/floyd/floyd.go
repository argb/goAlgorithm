package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt32

func main() {
	vertexes := []string{"A", "B", "C", "D", "E"}
	var matrix = [][]int{
		/*A*//*B*//*C*//*D*//*E*/
		/*A*/ {   0,  8,   1, 2, INF},
		/*B*/ {   8,  0, INF, 3, INF},
		/*C*/ {   1,  3,  0,  2,   3},
		/*D*/ {   2, INF, 2,  0,   3},
		/*E*/ { INF, INF, 3,  3,   0}}
	dist := floyd(matrix)
	n := len(dist)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			fmt.Printf("%s-->%s, dist[%d,%d]:%d\n", vertexes[i], vertexes[j], i, j, matrix[i][j])
		}
	}
}

func floyd(matrix [][]int) [][]int {
	n := len(matrix)
	for k:=0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j :=0; j < n; j++ {
				matrix[i][j] = min(matrix[i][j], matrix[i][k] + matrix[k][j])
			}
		}
	}
	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}