package main

import (
	"encoding/json"
	"fmt"
)

//结构体的序列化
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	Sal float64 `json:"sal"`
	Skill string `json:"skill"`
}

//将结构体序列化
func testStruct() {

	//演示
	monster := Monster{
		Name: "牛魔王",
		Age: 500,
		Birthday:"2011-11-011",
		Sal: 8000.00,
		Skill: "牛魔曲昂",
	}

	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误 err=%v", err)
	}

	fmt.Printf("monster序列化的结果=%v\n", string(data))
}


//将map序列化
func testMap() {

	//定义map
	var a map[string]interface{}
	//使用map，需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 200
	a["address"] = "洪崖洞"

	//序列化
	data, err := json.Marshal(a)

	if err != nil {
		fmt.Println("序列化错误 err=%v", err)
	}

	fmt.Printf("a map序列化的结果=%v\n", string(data))
}

//切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	//使用map前先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tome"
	m2["age"] = "25"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	//序列化
	data, err := json.Marshal(slice)

	if err != nil {
		fmt.Println("序列化错误 err=%v", err)
	}

	fmt.Printf("slice 序列化的结果=%v\n", string(data))
}

//普通数据类型序列化
func testBase() {

	//说明：对基本数据类型进行序列化的意义不大
	var num1 float64 = 144.11
	data, err := json.Marshal(num1)

	if err != nil {
		fmt.Println("序列化错误 err=%v", err)
	}

	fmt.Printf("num1 序列化的结果=%v\n", string(data))
}

func main() {

	//json序列化操作
	testStruct()
	testMap()
	testSlice()
	testBase()
}
