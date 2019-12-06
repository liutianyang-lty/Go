package main

import (
	"fmt"
)

type A struct {
	Num int
}

func (a A) test() {
	fmt.Println(a.Num)
}

func main() {

	//学习Golang中的方法
	//Golang中的方法是作用在指定的数据类型上的（即:和指定的数据类型绑定），因此自定义类型都可以有方法，而不仅仅是struct

	//1. 方法的声明和调用
	//说明： func (a A) test() {} 表示A结构体有一个方法，方法名叫test
	//(a A) 体现test方法是和 A 类型绑定的
}
