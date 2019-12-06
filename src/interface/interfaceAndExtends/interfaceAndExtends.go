package main

import (
	"fmt"
)

//Monkey结构体
type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, " 生来会爬树...")
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey
}

//让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, " 通过学习，会飞翔...")
}

//让LittleMonkey实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, " 通过学习，会游泳...")
}

func main() {

	//实现接口 和 继承

	//创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Swimming()
	monkey.Flying()


	//说明：
	//1. 继承的价值主要在于：解决代码的复用性和可复用性
	//2. 接口的价值主要在于：设计，设计好各种规范(方法)
	//3. 接口比继承更加灵活
	//4. 接口在一定程度上实现代码解耦
}
