package main

import (
	"fmt"
)

func main() {

	//学习数据类型map
	//map的声明和注意事项
	//注意：
	// 1. 这里声明之后是没有内存空间的，需要先make，make的作用是给map分配数据空间
	// 2. map中的key是不能重复的
	// 3. map的key-value是无序的
	a := make(map[string]string, 10)

	a["no1"] = "松江"
	a["no2"] = "无用"
	a["no1"] = "武松"
	a["no3"] = "武松"
	fmt.Println(a)

	//map声明的第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"
	fmt.Println(cities)

	//第三种：声明、直接赋值
	heroes := map[string]string{
		"heroes1" : "宋江",
		"heroes2" : "卢俊义",
	}
	fmt.Println(heroes)
}

