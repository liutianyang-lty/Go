package main

import(
	"fmt"
	"time"
)

func putNum(intChan chan int) {

	for i:=1; i<=40000; i++ {
		intChan<- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//var num int
	var flag bool
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <- intChan
		if !ok {
			break
		}
		flag = true
		//判断素数
		for i:=2; i<num; i++ {
			if num % i == 0 { //说明不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan<- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据，退出")

	//向exitChan 写入true
	exitChan<- true
}

func main() {

	//统计1-8000的数字中，哪些是素数
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()
	//开启协程向intChan写数据
	go putNum(intChan)

	//开启4个协程，从intChan取出数据，并判断是否为素数。如果是，就放入primeChan
	for i:=0; i<4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗费的时间=", end - start)
		//当从exitChan取出了4个结果后，就可以关闭primeChan管道
		close(primeChan)
	}()

	//遍历primeNum, 取出结果
	for {
		_, ok := <- primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("主线程退出")
}
