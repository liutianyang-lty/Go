package main

import "fmt"

//被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}

	return res
}

func main() {

	//学习单元测试

	//传统方式测试
	res := addUpper(10)
	if res != 15 {
		fmt.Printf("addUpper错误 返回值=%v 期望值=%v", res, 55)
	} else {
		fmt.Println("addUpper执行正确")
	}


	//单元测试--基本介绍
	//Go语言中自带一个轻量级的测试框架testing和自带的go test 命令来实现
	//单元测试和性能测试，testing框架和其他语言中的测试框架类似，可以基于
	//这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例
}