package main


func main() {

	//单元测试细节
	//1. 测试用例文件名必须以_test.go结尾。比如cal_test.go
	//2. 测试用例函数必须以Test开头，一般来说就是Test+被测试的函数名，比如TestAddUpper(大驼峰命名)
	//3. TestAddUpper(t *testing.T) 的形参必须是*testing.T
	//4. 一个测试用例文件中，可以有多个测试用例函数。比如TestAddUpper, TestSub
	//5. 运行测试用例指令：go test [如果运行正确，无日志，错误时，会输出日志]
	//   go test -v [运行正确或错误，都输出日志]
	//6. 当出现错误时，可以使用t.Fataf来格式化输出错误信息，并退出程序
	//7. t.Logf方法可以输出相应的日志
	//8. PASS表示测试用例运行成功，FAIL表示运行失败

	//9. 测试单个文件，一定要带上被测试的原文件：go test -v cal_test.go cal.go
	//10. 测试单个函数：go test -v -test.run TestAddUpper
}
