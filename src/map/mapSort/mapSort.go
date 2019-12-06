package main

import (
	"fmt"
	"sort"
)

func main() {

	//学习map排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 24
	map1[3] = 35
	map1[4] = 53
	map1[6] = 33
	fmt.Println(map1)

	//按照map的key的顺序排序输出
	//1. 现将map的key放入切片中
	//2. 对切片排序
	//3. 遍历切片，然后按照key来输出map的值

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map[%v]=%v\t", k, map1[k])
	}
}

