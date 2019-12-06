package main

import (
	"fmt"
	"testing" //引入go的 testing框架包
)

//编写一个测试用例，去测试函数addUpper是否正确
func TestAddUpper(t *testing.T) {

	//调用
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper错误 返回值=%v 期望值=%v", res, 55)

	}

	//正确则输出日志
	t.Logf("AddUpper(10) 执行正确...")
}

//只要以Test开头的函数就会被调用
func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用")
}