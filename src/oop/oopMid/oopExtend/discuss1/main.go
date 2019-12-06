package main

import (
	"fmt"
)

type A struct {
	Name string
	age int
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
		fmt.Println("B SayOk", b.Name)
}

func (a *A) SayOk() {
	fmt.Println("A SayOk",a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello",a.Name)
}

func main() {

	//1. 结构体可以使用嵌套匿名结构体所有的字段和方法

	//2. 匿名结构体字段访问可以简化：b.A.age ==> b.age

	//3. 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问。如果
	//   希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分

	//4. 结构体嵌入两个(或多个)匿名结构体，如两个匿名结构体有相同的字段和方法(同时结构体本身
	//   没有相同的字段或方法)，在访问时，就必须明确指定匿名结构体名字，否则会编译报错

	//5. 如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在
	//   访问组合的结构体的字段或方法时，必须带上结构体的名字

	//6. 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值

	//var b B
	//b.A.Name = "tom"
	//b.A.age = 19
	//b.A.SayOk()
	//b.A.hello()
	//
	////将上面的简化
	//b.Name = "smith"
	//b.age = 24
	//b.SayOk()
	//b.hello()

	var b B
	b.Name = "jack"
	b.age = 100
	b.SayOk()
	b.hello()

	var a A
	fmt.Println(a.Name)
	a.SayOk()
	a.hello()
}
