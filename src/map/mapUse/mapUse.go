package main

import (
	"fmt"
)

func main() {
	//map的引用案例
	//我们要存放3个学生的信息，每个学生有name和sex信息
	//map[string]map[string]string

	studentMap := make(map[string]map[string]string, 10)

	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"

	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"

	fmt.Println(studentMap["stu01"])
}
