package main

import (
	"fmt"
	"reflect"
)

//Cal结构体
type Cal struct {
	Num1 int
	Num2 int
}

//方法
func (this Cal) GetSub(name string) {
	fmt.Printf("%s 完成了减法运行，%d - %d = %d\n", name, this.Num1,
		this.Num2, this.Num1 - this.Num2)
}

func main() {

	//作业：
	//1. 编写一个Cal结构体，有两个字段Num1，Num2
	//2. 方法 GetSub(name string)
	//3. 使用反射遍历Cal结构体所有的字段
	//4. 使用反射机制完成对GetSub的调用，输出形式为 "tom 完成了减法运行，8 - 3 = 5"


	//(1) 声明一个结构体变量
	cal := Cal{}

	//(2) 获取reflect.Value
	val := reflect.ValueOf(&cal)

	//(3) 获取结构体含有的字段数
	num := val.Elem().NumField()
	fmt.Printf("cal结构体有%d个字段\n", num) //2

	//(4) 给机构体设置字段的值
	val.Elem().Field(0).SetInt(8)
	val.Elem().Field(1).SetInt(3)

	//(5) 或得结构体的方法的个数
	numMethod := val.Elem().NumMethod()
	fmt.Printf("cal结构体有%d个方法\n", numMethod)

	//(6) 调用结构体的方法
	var calReflectValue []reflect.Value
	calReflectValue = append(calReflectValue, reflect.ValueOf("tom"))
	val.Elem().Method(0).Call(calReflectValue)
}
