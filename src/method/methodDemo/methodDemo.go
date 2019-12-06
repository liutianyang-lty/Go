package main

import (
	"fmt"
)

//定义一个结构体
type Person struct {
	Name string
}

//给Person类型绑定一个方法
func (p Person) test() {
	p.Name = "jack"
	fmt.Println("test() name=", p.Name)
}

//给Person类型添加speak方法
func (p Person) speak() {
	fmt.Println(p.Name, "是一个goodman~~")
}

//给Person类型添加 jisuan 方法
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果为", res)
}

//给Person类型添加 jisuan2 方法
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果为", res)
}

//给Person类型添加 getSum 方法
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {

	var p Person
	p.Name = "tom"
	p.test() //调用方法

	//对上面的总结：
	//1. test方法和Person类型绑定
	//2. test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其他类型的变量来调用
	//3. func (p Person) test() {}  p表示哪个Person变量调用，这个p就是它的副本，这点和函数传参非常相似
	//4. p 这个名字，由程序员指定，不是固定的，一般取有意义的名字

	//方法的快速入门
	//(1) 给Person结构体添加speak方法，输出 xxx 是一个好人
	p.speak()

	//(2) 给Person结构体添加 jisuan 方法， 可以计算从1+..+ 1000的结果，说明方法体内可以和函数一样进行各种运算
	p.jisuan()

	//(3) 给Person结构体添加 jusuan2 方法，该方法可以接受一个数n, 计算从1到n的和
	p.jisuan2(10)

	//(4) 给Person结构体添加getSum方法，可以计算两个数的和，并返回结果
	res := p.getSum(10, 20)
	fmt.Println(res)


	//方法的调用和传参机制原理：
	//说明：方法的调用、传参机制和函数的基本一样，不一样的地方是方法调用时，会将调用方法的变量当做实参也传递给方法
}
