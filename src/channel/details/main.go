package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {

	//channel使用细节：
	//1.channel 中只能存放指定的数据类型
	//2.channel 的数据放满后，不能再放入
	//3.如果从channel中取出数据后，还可以继续放入
	//4. 在没有使用协程的情况下，如果channel数据取完了，再取，就会报dead lock错误

	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	//向管道中添加数据
	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan<- cat

	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat = %T, newCat = %v\n", newCat, newCat)

	fmt.Printf("newCat.Name=%v",newCat.(Cat).Name)  //这里要使用类型断言取得猫的名字
}
