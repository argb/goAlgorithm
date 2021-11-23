package main

import (
	"algorithm/powerset"
	"fmt"
	"time"
)

/**
01 背包问题
 */
func main() {
	capacity := 10
	info := []Item {
		{"A", 2, 2},
		{"B", 3, 4},
		{"C", 5, 3},
		{"D", 5, 7},
	}

	// 暴力搜索所有情况求解
	start := time.Now()
	ViolentSearch(info, capacity)
	fmt.Println("time:", time.Now().Sub(start).Microseconds())

	// 动态规划方法求解
	start = time.Now()
	table := dp(info, capacity)
	maxValue := getMaxValue(table, len(info), capacity)
	fmt.Printf("The max Value: %d\n", maxValue)
	path := getPath(table, &info, capacity)
	fmt.Printf("The best combination: %v\n", path)
	fmt.Println("time:", time.Now().Sub(start).Microseconds())

	/*
	测试结论：
	问题规模小的时候比如，n=4，两者差别不是很大，当然还是dp快，几倍
	当问题规模n=12的时候，速度差22691/46=491倍，基本可以确定暴力搜索解法基本没什么实用价值，实在太慢了。
	 */
}

type Item struct {
	Name string
	Weight int
	Value int
}

/**
Dynamic programing solution
time: O(n*v) n是物品数量，v是背包容量
注意这里按照"标准时间复杂度"计算的话，这是个伪多项式时间，v是个指数级输入规模
这里有个讨论 https://www.zhihu.com/question/20013122
但是我们的输入规模并不是线性增长的，也就是说v的增长很慢，达不到线性增长的程度，如果v是线性增长，那么基于
输入规模的时间复杂度就是指数级的。
所以这个按照输入规模的定义的时间复杂度在这个问题上并没有太多实际意义，基于问题规模的传统时间复杂度能
更好的反应问题。

这里有个不错的讲解 https://blog.csdn.net/qq_21989927/article/details/108471314
 */
func dp(items[]Item, capacity int) [][]int{
	//capacity = 10
	n := len(items)
	var table [][] int

	table = make([][]int, n+1)
	for i:=0; i<=n; i++{
		table[i] = make([]int, capacity+1)
	}
	//printDpTable(table)
	for i := 1; i<=n; i++ {
		for v :=1; v <= capacity; v++ {
			currItem := items[i-1]
			// 递推公式
			if currItem.Weight <= v { // 选择当前物品
				newValue := currItem.Value + table[i-1][v-currItem.Weight]
				table[i][v] = max(newValue, table[i-1][v])
			}else { // 不选择
				table[i][v] = table[i-1][v]
			}
		}
	}

	printDpTable(table)
	return table
}

func getPath(table [][]int, items *[]Item, capacity int) []int {
	n := len(*items)
	v := capacity
	var path []int

	for i := n; i >= 1; i-- {
		if table[i][v] != table[i-1][v] {
			v = v - (*items)[i-1].Weight
			// 这里需要的是物品基于0的下表，但是之前计算dp table的时候多加了一行和一列，表示取0个物品和背包容量为0的情况
			// 所以这里的i值比物体的索引大1，所以这里将其减去
			path = append(path, i-1)
		}
	}

	return path
}

func getMaxValue(table [][]int, n int, v int) int {
	return table[n][v]
}

func printDpTable(table [][]int) {
	for i := 0; i< len(table); i ++ {
		fmt.Println(table[i])
	}
}


func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

// ViolentSearch /**
/*
暴力解法，但是这个求幂集的过程也不是很慢，所以总体上也还可以
 */
func ViolentSearch(items []Item, capacity int) {

	var names []string
	for i:=0; i<len(items); i++ {
		names = append(names, items[i].Name)
	}

	ps := powerset.PowerSet(names)
	fmt.Println(ps)
	bestCom := filtrate(ps, items, capacity)
	fmt.Println("the best combination:")
	fmt.Println(bestCom)
}

func filtrate(powerSet [][]int, items []Item, capacity int) []int {
	var maxIndex int
	var maxV int

	for i:=0; i < len(powerSet); i++ {
		var sumW int
		var sumV int
		subSet := powerSet[i] // 每个子集是一组可能的解
		for j:=0;j<len(subSet); j++ {
			itemIndex := subSet[j]
			sumW += items[itemIndex].Weight
			sumV += items[itemIndex].Value
		}
		if sumW <= capacity && sumV > maxV {
			maxV = sumV
			maxIndex = i
		}
	}
	return powerSet[maxIndex]
}