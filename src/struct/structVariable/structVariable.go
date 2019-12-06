package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {

	//创建结构体变量和访问结构体字段
	//方式一：直接声明  var person Person
	//方式二：var person Person = Person{}
	p2 := Person{}
	fmt.Println(p2)

	//方式三：var person *Person = new(Person)
	p3 := new(Person)
	//(*p3).Name = "hello" 也可以这样写 p3.Name = "hello"
	//原因：go的设计者 为了程序员使用方便，底层对p3.Name = "hello" 进行了处理
	//会给 p3 加上 取值运算 (*p3).Name = "hello"
	(*p3).Name = "hello"
	(*p3).Age = 18
	p3.Age = 100
	fmt.Println(*p3)

	//方式四：var person *Person = &Person{}
	var p4 *Person = &Person{}
	(*p4).Name = "scott"
	p4.Name = "bob"
	(*p4).Age = 32
	fmt.Println(*p4)

	//第3种和第4种方式返回的是结构体指针
	//结构体指针访问字段的标准方式是：(*结构体指针).字段名
}
