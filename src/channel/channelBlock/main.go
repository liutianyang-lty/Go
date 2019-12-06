package main

import(
	"fmt"
	"time"
)

//writeData()
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan<- i
		fmt.Println("writeData", i)
	}
	close(intChan)
}

//readData()
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second) //写的快，读的慢
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	//读取完数据后, 向exitChan写数据
	exitChan<- true
	close(exitChan)
}

func main() {

	//goroutine和channel协同工作的案例：
	//1) 开启一个writeData协程，向管道intChan中写入50个整数
	//2) 开启一个readData协程，从管道intChan中读取数据
	//3) writeData和readData操作的是同一个管道
	//4) 主线程需要等待writeData和readData协程都完成工作才能退出

	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <- exitChan
		if !ok {
			break
		}
	}
}