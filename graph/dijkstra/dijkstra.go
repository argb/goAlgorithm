package dijkstra

import (
	"fmt"
	"math"
)

const VertexNum = 5
const INF = math.MaxInt32 // infinity
type Graph struct {
	vertexes []string // 顶点集合
	vertexNum int
	edgeNum int
	matrix [][VertexNum]int
}

func (g *Graph) InitTable(vertexes []string, len int, matrix [][5]int) *Graph{
	g.vertexes = append(g.vertexes, vertexes...)
	g.vertexNum = len
	g.matrix = matrix

	return g
}

func (g *Graph) Dijkstra(vs int, dist []int,) {
	var visited[VertexNum]bool // 用于标记已经获得最短路径的点
	var minDist int
	var minDistIndex int
	var tmp int
	// 这里存的是前驱节点的[索引], 当前节点i，那么如果path[i]的值是j的话，j就是最短路径上i前面那个节点，也就是说
	//vs->...->n 是一条最短路径的话，那么j->i一定在这条路径上，而且j是i前面的邻居,the previous neighbour.
	//所以可以通过这个数组找出任意点开始到任意点结束的路径。
	var path []int

	path = make([]int, VertexNum)
	matrix := g.matrix
	// 根据权重初始化每个点距离起点vs的初始最短路径
	//相邻的就是边的权值，不相邻的就是无穷大，自己到自己就是0
	for i:=0; i<VertexNum; i++ {
		dist[i] = matrix[vs][i]
		visited[i] = false
		path[i] = vs
	}
	// 记录路径，这里是起点


	dist[vs] = 0
	visited[vs] = true // 把自己排除掉，否则最短路径永远是自己到自己，原地不动

	// N-1轮，因为除了顶点还剩下n-1个点，每一轮确定某个点的到vs的最短路径，结果计入dist[i]
	for i := 1; i<VertexNum; i++ { // 每一轮的vertex[i]是目标点，也就是终点end
		minDist = math.MaxInt32
		// 遍历剩余的点，找出可以使得vertex[vs]->vertex[i]最短的点
		for j := 0; j<VertexNum; j++ {
			if visited[j] == false && dist[j] < minDist {
				minDist = dist[j]
				minDistIndex = j
			}
		}
		visited[minDistIndex] = true // 标记为已经获得最短路径

		//修正所有点距离起点的vs的最短路径
		for j :=0; j< VertexNum; j++ {
			//因为这是个无向图，所以matrix[j][minDistIndex]跟matrix[minDistIndex][j]其实是相等的，整个矩阵是沿对角线对称的
			if matrix[minDistIndex][j] != INF {
				tmp = matrix[minDistIndex][j] + minDist
			}else {
				tmp = INF
			}
			if visited[j] == false && tmp < dist[j] {
				dist[j] = tmp
				path[j] = minDistIndex
			}
		}
	}
	fmt.Println("visited:", visited)
	fmt.Println("dijkstra")
	fmt.Printf("vertexes: %v\n", g.vertexes)
	for i:=0; i< VertexNum; i++ {
		if i == vs {
			// 不能起点到起点
			continue
		}
		t := i
		var pathStr string
		for t != vs {
			pathStr += g.vertexes[t] + "<--"
			t = path[t]
		}
		pathStr += g.vertexes[vs]

		fmt.Printf("shortestDistance(%v, %v) = %d\n",g.vertexes[vs], g.vertexes[i], dist[i])
		fmt.Printf("path:%s\n\n", pathStr)
	}
	//fmt.Println(dist)
	//fmt.Println("path:", path)
}
