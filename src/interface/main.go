package main

import (
	"fmt"
)

//声明一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

//手机
type Phone struct {

}

//相机
type Camera struct {

}

//计算机
type Computer struct {

}

//让结构体Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机结束工作...")
}

//Camera实现接口
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机结束工作...")
}

//编写一个方法Working方法，接受一个Usb接口变量
//只要实现了Usb接口
func (c Computer) Working(usb Usb) {
	//通过usb接口变量调用方法
	usb.Start()
	usb.Stop()
}

func main() {

	//测试
	//创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
