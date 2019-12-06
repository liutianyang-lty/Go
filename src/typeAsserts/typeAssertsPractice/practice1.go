package main

import (
	"fmt"
)

//定义一个Usb接口
type Usb interface {
	Start()
	Stop()
}

//定义结构体Phone
type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机结束工作...")
}

func (p Phone) Call() {
	fmt.Println("手机可以打电话...")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机结束工作...")
}

type Computer struct {

}

func (computer Computer) Working(usb Usb) {
	usb.Start()

	//使用类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {

	//类型断言最佳实践
	//定义一个usb接口数组，可以存放Phone和Camera的结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"华为"}
	usbArr[2] = Camera{"尼康"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
