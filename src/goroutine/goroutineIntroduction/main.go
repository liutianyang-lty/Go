package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写一个函数，每隔1秒输出 "hello,world"
func test() {
	for i:=1; i<=10; i++ {
		fmt.Println("test () hello,world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() { //主线程

	//goroutine-快速入门
	//案例：
	//在主线程中，开启一个goroutine，该协程每隔1秒输出"hello,world"
	//在主线程中也每隔一秒输出"hello,golang"，输出10次后，退出程序
	//要求主线程和goroutine同时执行

	go test() //开启了一个协程

	for i:=1; i<=10; i++ {
		fmt.Println("main () hello,golang"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}


	//快速入门小结：
	//1. 主线程是一个物理线程，直接作用在cpu上。是重量级的，非常耗费CPU资源
	//2. 协程从主线程开启，是轻量级的线程，是逻辑态。对资源消耗相对小
	//3. Golang的协程机制是重要的特点，可以轻松开启上万个协程。其他编程语言
	//   的并发机制是一般基于线程的，开启过多的线程，资源耗费大
}
