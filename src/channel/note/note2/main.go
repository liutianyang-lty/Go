package main

import (
	"fmt"
	"time"
)

//函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello, world")
	}
}

func test() {

	defer func() {
		//捕获抛出 的 panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" //错误用法,此时报错：panic: assignment to entry in nil map
}

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10 ; i++  {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}

	//开发 中，通常如果一个协程出现问题，是不希望影响其他协程的运行的
	//解决方法：如上
}