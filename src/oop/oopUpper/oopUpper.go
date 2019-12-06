package main

import (
	"fmt"
)

//快速入门struct
//声明结构体
type Cat struct {
	Name string
	Age int
	Color string
	hobby string
}

func main() {

	//学习Go中的面向“对象”
	//其实Go中是没有类这个概念的，而是以结构体来作为类的替代
	//Golang语言面向对象编程说明：
	//(1) Golang也支持面向的对象编程，但是和传统的面向对象编程有区别，并不是纯粹的面向对象语言
	//(2) Golang没有类（class）,Go语言的结构体（struct）和其他编程语言的类（class）有相同的地位，可以理解为Golang是基于struct来实现OOP特性的
	//(3) Golang面向对象编程非常简洁，去掉了传统的OOP语言的继承、方法重载、构造函数和析构函数、隐藏的this指针等等
	//(4) Golang仍然有面向对象编程的继承、封装和多态的特性，只是实现方式和其他语言不一样
	//(5) Golang面向对象很优雅，OOP本省就是语言类型系统的一部分，通过接口(interface)关联，耦合性低。Golang面向接口编程是非常重要的特性

	//创建一个结构体变量（结构体是值类型）
	var cat1 Cat
	cat1.Name = "小花"
	cat1.Age = 2
	cat1.Color = "花色"
	cat1.hobby = "吃鱼"
	fmt.Println("cat1=\n", cat1)
	fmt.Printf("结构体cat1的地址为%p\n", &cat1)
	fmt.Printf("cat1.Name的地址为%p\n", &cat1.Name)
	fmt.Printf("cat1.Age的地址为%p\n", &cat1.Age)
	fmt.Printf("cat1.Color的地址为%p\n", &cat1.Color)
	fmt.Printf("cat1.hobby的地址为%p\n", &cat1.hobby)

	//匿名结构体
	person := struct {
		name string
		age int
	}{name:"帐篷", age:18}
	fmt.Println("person=", person)
}

