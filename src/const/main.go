package main

import "fmt"

func main() {

	//常量的学习
	//1.常量使用const声明
	//2.常量在定义时必须初始化
	//3.常量不能修改
	//4.常量只能修饰bool、数值类型(int、float类型)、string类型
	//5.Golang中没有规定说常量必须大写
	//6.仍然通过首字母的大小写来控制常量的访问范围

	const tax int = 100
	fmt.Println(tax)

	//const声明方式：
	//一：
	const (
		a = 1
		b = 2
	)

	//二：
	//const (
	//	c = iota
	//	d
	//	e,f = iota, iota  //输出结果：0 1 2 2
	//)

	//const (
	//	c = iota
	//	d
	//	e
	//	f  //输出结果为0 1 2 3
	//)

	//fmt.Println(c,d,e,f)
}
