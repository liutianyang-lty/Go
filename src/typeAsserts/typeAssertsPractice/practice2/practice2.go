package main

import (
	"fmt"
)

//编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items... interface{}) {
	
	for index, x := range items {
		switch x.(type) {
		case bool :
			fmt.Printf("第%v个参数是bool类型， 值是%v", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型， 值是%v", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型， 值是%v", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型， 值是%v", index, x)
		default:
			fmt.Println("参数类型不确定")
		}
	}
}

func main() {

	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var name string = "tom"
	TypeJudge(n1, n2, name)
}