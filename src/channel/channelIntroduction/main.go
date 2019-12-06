package main

import (
	"fmt"
)

func main() {

	//定义/声明
	//var 变量名 chan 数据类型
	//例子：var intChan chan int

	//说明：
	//channel 是引用类型
	//channel 必须初始化后才能写入数据，即make后才能使用
	//管道是有类型的，intChan 只能写入整数int
	//管道是不能动态增长的

	//创建管道
	var intChan chan int
	intChan = make(chan int, 3)


	fmt.Println(intChan) //此时返回的是一个地址

	//向管道写入数据
	intChan<- 10
	num := 233
	intChan<- num
	intChan<- 334

	// 看看管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	// 获取管道内的元素
	var num2 int
	num2 = <-intChan
	num3 := <- intChan
	num4 := <- intChan
	//num5 := <- intChan //这里会报错：fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	//fmt.Println(num5)
}
