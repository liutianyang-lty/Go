package main

import (
	"testing" //引入go的 testing框架包
)

//编写一个测试用例，去测试函数addUpper是否正确
func TestGetSub(t *testing.T) {

	//调用
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("getSub错误 返回值=%v 期望值=%v", res, 7)

	}

	//正确则输出日志
	t.Logf("getSub 执行正确...")
}
