package main

import (
	"fmt"
	"time"
)

func main() {

	//使用select可以解决从管道读取数据的阻塞问题
	intChan := make(chan int, 10)
	for i:=0; i<10; i++ {
		intChan<- i
	}

	stringChan := make(chan string, 5)
	for i:=0; i<5; i++ {
		stringChan<- "hello" + " " + fmt.Sprintf("%d", i)
	}

	//传统方法在遍历管道时，如果不关闭管道会阻塞而导致 deadlock
	//在实际开发中，我们可能不好确定什么时候关闭管道
	//可以使用select方式解决
	for {
		select {
			//注意：这里，如果inChan一直没有关闭，不会一直阻塞而导致deadlock
			//会自动到下一个case匹配
			case v := <-intChan:
				fmt.Printf("从intChan读取的数据%d\n", v)
				time.Sleep(time.Second)
			case v := <-stringChan:
				fmt.Printf("从stringChan读取的数据%s\n", v)
				time.Sleep(time.Second)
			default:
				fmt.Printf("都取不到了，程序员可以加入自己的逻辑\n")
				time.Sleep(time.Second)
				break

		}
	}
}
