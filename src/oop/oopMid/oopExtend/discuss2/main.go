package main

import (
	"fmt"
)

type A struct {
	Name string
	age int
}

type B struct {
	Name string
	Score float64
}

type C struct {
	A
	B
}

type D struct {
	a A //嵌套了一个有名结构体
}


type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {

	//继承的深入讨论2
	var c C
	c.A.Name = "tom"
	fmt.Println("A 的name", c.A.Name)

	var d D
	d.a.Name = "tony" //这里必须带上a
	fmt.Println("有名结构体的name", d.a.Name)

	tv := TV{
		Goods{"电视机001", 5000.99},
		Brand{"海尔","山东"},
	}

	fmt.Println("tv", tv)
}
