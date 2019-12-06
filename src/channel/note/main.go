package main

import (
	"fmt"
)

func main() {

	//channel的注意事项和使用细节
	//1. 管道可以声明为只读或者只写（在默认情况下，管道是双向的）

	//(1)声明为只写
	var chan1 chan<- int
	chan1 = make(chan int, 3)
	chan1<- 100
	//num := <-chan1  //这里会报错

	//(2)声明为只读
	var chan2 <-chan int
	num := <-chan2
	//chan2<- 30 //这里会报错
	fmt.Println(num)

	//2. 使用select可以解决从管道取数据的阻塞问题

	//3. goroutine中使用recover，解决协程中出现panic的导致程序崩溃的问题

}
