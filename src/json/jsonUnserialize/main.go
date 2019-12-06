package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name string
	Age int
	Birthday string
	Sla float64
	Skill string
}

//将json字符串，反序列化struct
func unmarshalStruct() {

	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-01\",\"Sla\":8000.00,\"Skill\":\"牛魔曲昂\"}"

	//定义一个monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarshal err=%v", err)
	}

	fmt.Printf("反序列化后 monster=%v monster.Name=%v\n", monster, monster.Name)
}

//将json字符串反序列化为map
func unmarshalMap() {

	str := "{\"address\":\"洪崖洞\",\"age\":200,\"name\":\"红孩儿\"}"

	//定义map
	var a map[string]interface{}

	//反序列化
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err = %v", err)
	}
	fmt.Printf("反序列化后a=%v\n", a)
}

//将json字符串反序列化为slice
func unmarshalSlice() {

	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"25\",\"name\":\"tome\"}]"

	//定义slice
	var slice []map[string]interface{}

	//反序列化
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err = %v", err)
	}
	fmt.Printf("反序列化后slice=%v\n", slice)
}

func main() {

	//json反序列化
	//unmarshalStruct()
	//unmarshalMap()
	unmarshalSlice()

	//反序列化说明
	//1. 在反序列化一个json字符串时，要确保反序列化后的数据类型和原来序列化前的数据类型一致
	//2. 如果json字符串时通过程序渠道的，则不需要再对 双引号" 进行转义处理
}
