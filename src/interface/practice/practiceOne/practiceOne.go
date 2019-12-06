package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age int
}

//声明一个Hero结构体切片类型
type HeroSlice []Hero

//实现接口Len()方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less方法就是决定你使用什么标准进行排序
//1. 按Hero的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

//实现接口Swap()方法
func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {

	//接口的最佳实践一：实现对Hero结构体切片的排序 sort.Sort(data)
	//var intSlice = []int{0, -1, 21, 4, 76}
	//sort.Ints(intSlice)
	//
	//fmt.Println(intSlice)

	//对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}

		heroes = append(heroes, hero)
	}

	//排序前
	for _, v := range heroes {
		fmt.Println(v)
	}

	//调用sort.Sort()
	sort.Sort(heroes)
	fmt.Println()
	for _, v := range heroes {
		fmt.Println(v)
	}
}
