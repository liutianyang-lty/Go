package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"name"` //这就是结构体的标签
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	//注意点4：struct 的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
	//创建一个monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~~"}

	//将monster变量序列化为 json格式的字符串
	jsonStr, err := json.Marshal(monster) //原理是json.Marshal函数在底层使用了反射
	if err != nil {
		fmt.Println("json处理错误：", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
