package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {

	//学习类型断言
	var a interface{}
	var point Point = Point{1, 2}
	a = point

	var b Point
	//b = a  //这里是不可以直接赋值的
	//可以这样写
	b = a.(Point)
	//上面一行代码就是类型断言，表示判断a是否指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则报错

	fmt.Println(b)

	//类型断言的其它案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口可以接收任意类型
	// x=>float32
	y := x.(float32)
	fmt.Printf("y的类型是%T", y)



	//带检测的类型断言
	z, ok := x.(float64)

	if ok {
		fmt.Printf("%T\n", z)
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println("继续执行...")

}
