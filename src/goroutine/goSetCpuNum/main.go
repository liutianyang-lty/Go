package main

import (
	"fmt"
	"runtime"
)

func main() {

	//获取当前系统的CPU个数
	num := runtime.NumCPU()

	//自己设置使用的CPU个数
	runtime.GOMAXPROCS(num - 1)
	fmt.Println("num=", num)
}