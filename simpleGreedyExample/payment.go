package main

import (
	"fmt"
	"sort"
)


type Person struct {
	Name string
	wallet Wallet // key是币值，value是数量
}
type Wallet map[int]int

func main() {
	xiaoming := new(Person)
	xiaoming.Name = "小明"
	xiaoming.wallet = map[int]int{
		1: 5,
		5: 2,
		10: 2,
		50: 3,
		100: 5,
	}
	num, record := xiaoming.pay(456)
	fmt.Printf("minimum number:%d, combination:%v\n", num, record)
}

func (p *Person) pay(money int) (int,map[int]int) {
	var num int
	var record map[int]int
	//var kind = len(p.money)
	record = make(map[int]int)
	var keys []int
	for key, _ := range p.wallet {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for i := len(keys) -1; i >= 0; i -- {
		key := keys[i]
		value := p.wallet[key]

		//fmt.Println(key,value)
		count := min(money/key, value) // greedy strategy: 每次尽可能多的拿大面值
		record[key] = count
		money -= count * key
		num += count
	}

	return num, record
}

func min(a, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}
