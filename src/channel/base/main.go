package main

import (
	"fmt"
	"sync"
	"time"
)

//思路
//1. 编写一个函数，来计算各个数的阶乘，并放入到一个map中
//2. 启动多个协程，统一的将结果放入到map中
//3. map应该是全局的

var (
	myMap = make(map[int]int, 10)

	//声明一个全局互斥锁
	//lock就是一个全局互斥锁 (写锁)
	//Mutex: 互斥
	lock sync.Mutex
)

//计算n的阶乘，然后将结果放入map中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//将结果放入map
	//加锁
	lock.Lock()
	myMap[n] = res  //会出现资源竞争问题
	//解锁
	lock.Unlock()
}

func main() {

	//channel学习
	//需求：现在要计算1-200 的各个数的阶乘，并把各个数的阶乘放入到map中
	//最后显示出来，要求使用goroutine完成

	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠
	time.Sleep(time.Second * 4)

	//为什么读的时候也需要加锁：
	//
	lock.Lock()
	//输出结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()





	//协程之间的资源竞争问题：
	//1. 全局变量加锁同步
	//2. channel

	//为什么需要channel
	//(1) 主线程在等待所有的goroutine全部完成的时间很难确定
	//(2) 如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能
	//    还有goroutine处于工作状态，这时也会随主线程的退出而销毁
	//(3) 通过全局变量加锁同步来实现通讯，也不利于用多个协程对全局变量的读写操作
	//(4) 于是引出新的通讯机制--channel


	//channel基本介绍
	//1. channel本质就是一个数据结构--队列
	//2. 数据是先进先出FIFO
	//3. 线程安全，多个goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
	//4. channel是有类型的。如果channel是string类型，则里面只能存放string类型
}
