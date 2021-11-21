package powerset

import "fmt"

func main() {
	var initial []string
	initial = []string{
		"a", "b", "c", "d", "e",
	}

	result := PowerSet(initial)

	result3 := getCombinationN(3, result)
	transformedResult3 := transform(initial, result3)


	fmt.Println(transformedResult3)
}

func transform(src []string, target [][]int)[][]string{
	var result [][]string
	for i:=0; i<len(target); i++{
		t := target[i]
		var temp []string
		for j:=0; j<len(t); j++{
			temp = append(temp, src[t[j]])
		}
		result = append(result, temp)
	}

	return result
}

/**
某个长度的所有可能组合
 */
func getCombinationN(n int, set [][]int) [][]int {
	var result [][] int
	for i:=0; i< len(set); i++ {
		if len(set[i]) == n {
			result = append(result, set[i])
		}
	}
	fmt.Printf("length N(%d):\n", n)
	fmt.Println(result)
	return result
}

// PowerSet /**
/*
幂集，也就是所有可能的组合(子集）,2^n个，n是初始集合的元素个数
 */
func PowerSet(arr []string) [][] int{
	var result [][]int
	var first []int
	result = append(result, first)

	for i:=0; i< len(arr); i++ {
		var newArr [][]int
		newArr = deepCopy(newArr, result)
		for j :=0; j < len(newArr); j++{
			newArr[j] = append(newArr[j], i)
		}
		result = arrayMerge(result, newArr)
	}
	fmt.Println("All:")
	fmt.Println(result) // 问题规模变大的时候这个数组会变得非常大，因为枚举了所有可能的组合情况，指数级的增长速度2^n。

	return result
}

func deepCopy(dst, src [][]int) [][]int{
	for _, item := range src {
		dst = append(dst, item)
	}

	return dst
}

func arrayMerge(dst, src [][]int) [][]int {
	dst = append(dst, src...)
	return dst
}