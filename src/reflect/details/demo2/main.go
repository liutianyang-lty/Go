package main

import (
	"fmt"
	"reflect"
)

//定义一个结构体
type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex string
}

//返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//显示s的值
func (s Monster) Print() {
	fmt.Println("---start~---")
	fmt.Println(s)
	fmt.Println("---end~---")
}

//使用反射
func TestStruct(a interface{}) {

	//使用反射操作结构体

	//获取reflect.Type类型
	typ := reflect.TypeOf(&a)

	//获取reflect.Value类型
	val := reflect.ValueOf(&a)

	//获取到a对应的类别
	kd := val.Elem().Kind()
	fmt.Println(kd)
	//如果传入的不是struct就退出
	if kd != reflect.Struct {  //这里的比较一定要注意
		fmt.Println("expect struct")
		return
	}

	//获取该结构体有几个字段
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("犀牛精")
	fmt.Printf("struct has %d fields\n", num)

	//遍历结构体的所有字段
	for i:=0; i<num; i++ {

		fmt.Printf("Field %d: 值为%v\n", i, val.Field(i))

		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json") //注意：这里必须使用 typ
		if tagVal != "" {
			fmt.Printf("Field %d: tag=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//调用结构体的第2个方法
	//这里是需要特别注意的地方：方法的排序默认是按照 函数名的ASCII码 来排序的
	val.Method(1).Call(nil)

	//调用结构体的第1个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)

	fmt.Println("res=", res[0].Int())


}

func main() {

	//使用反射遍历结构体的字段，调用结构体的方法，并获取结构体标签的值

	//创建一个Monster实例
	a := Monster{
		Name : "黄鼠狼经",
		Age : 400,
		Score : 40.2,
	}

	//TestStruct(a) //不传地址
	TestStruct(&a)
	fmt.Println(a)
}
