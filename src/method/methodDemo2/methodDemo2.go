package main

import (
	"fmt"
)

//声明结构体Circle
type Circle struct {
	Radius float64
}

//给结构体Circle添加方法area
func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {

	//案例二：
	//(1) 声明一个结构体Circle，字段为radius
	//(2) 声明一个方法area和Circle绑定，可以返回面积
	var c Circle
	c.Radius = 4.0
	res := c.area()
	fmt.Println("面积是：", res)

}
