package main

import (
	"fmt"
)

func main() {

	//学习map切片:切片的数据类型是map，即切片中存放的是map类型的数据，这样map就可以动态的添加到切片中

	//声明一个map切片
	monsters := make([]map[string]string, 2)
	//向map切片添加数据
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "狐狸精"
		monsters[1]["age"] = "300"
	}

	//使用切片的append函数，实现动态增加
	//1. 先定义monster信息
	newMonster := map[string]string{
		"name" : "新妖怪",
		"age" : "200",
	}
	monsters = append(monsters,newMonster)
	fmt.Println(monsters)
}
