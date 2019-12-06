package main

import "fmt"

func main() {

	//channel的遍历和关闭

	//1. channel的关闭：
	//使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写
	//数据了，但是仍然可以从该channel读取数据
	intChan := make(chan int, 3)
	intChan<- 100
	intChan<- 200
	close(intChan)
	//intChan<-  300 //此时再写入的话会报错：panic: send on closed channel

	n1 := <-intChan
	fmt.Println(n1)



	//2. channel的遍历: 只能使用for-range方式
	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完成后，就会退出遍历

	intChan2 := make(chan int, 100)
	for i:=0; i<100; i++ {
		intChan2<- i * 2
	}

	//遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
