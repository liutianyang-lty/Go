package main

import (
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {

	var a A
	var b B
	fmt.Println(a,b)
	//注意点2：结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字，个数和类型）
	//数据类型转换
	a = A(b)
	fmt.Println(a,b)
}

