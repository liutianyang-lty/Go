package main

import (
	"fmt"
	"oop/factory/model"
)

func main() {

	//创建一个Student实例
	//var stu = model.student{
	//	Name : "tom",
	//	Score : 78.4,
	//}

	//当student结构体是首字母小写，使用工厂模s式
	var stu = model.NewStudent("tome1", 98.7)
	fmt.Println(stu)
	fmt.Println("name=",stu.Name, " SOCORE=", stu.GetScore()  )
}