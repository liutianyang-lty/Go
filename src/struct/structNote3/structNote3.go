package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
}

type Stu Student


func main() {

	//注意点3：结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强转
	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1)
	fmt.Println(stu1,stu2)
}
