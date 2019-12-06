package main

import (
	"fmt"
)

//结构体
type Rect struct {
	leftUp, rightDown Point
}

//结构体
type Rect2 struct {
	leftUp, rightDown *Point
}

//结构体
type Point struct {
	x int
	y int
}

func main() {

	//结构体使用注意事项和细节
	//1. 结构体的所有字段在内存中是连续的
	r1 := Rect{Point{1,2}, Point{3, 4}}
	//r1有四个整数，在内存中是连续分布的
	fmt.Printf("r1.lefUp.x 的地址=%p r1.leftUp.y 的地址=%p r1.rightDown.x 的地址=%p r1.rightDown.y 的地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	r2 := Rect2{&Point{10,20}, &Point{30,40}}
	fmt.Printf("r2.leftUp本身地址=%p r2.rightDown本身地址=%p\n", &r2.leftUp, &r2.rightDown)

	//他们指向的地址不一定是连续... 这个要看系统在运行时如何分配
	fmt.Printf("r2.leftUp指向的地址=%p r2.rightDown指向的地址=%p \n", r2.leftUp, r2.rightDown)

	//2. 结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字，个数和类型）
	//3. 结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强转
	//4. struct 的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
}
