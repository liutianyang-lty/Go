package main

import (
	"fmt"
)

type Student struct {
	Name string
	Gender string
	Age int
	Id int
	Score float64
}

type Box struct {
	len float64
	width float64
	height float64
}

type Visitor struct {
	Name string
	Age int
}

func (visitor * Visitor) showPrince() {


	if visitor.Age > 18 {
		fmt.Printf("游客的名字为%v 年龄为%v 收费20元\n", visitor.Name, visitor.Age)

	} else {
		fmt.Printf("游客的名字为%v 年龄为%v 免费\n", visitor.Name, visitor.Age)
	}
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

func (s *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		s.Name,s.Gender,s.Age,s.Id,s.Score)
	return infoStr
}

func main() {

	//面向对象编程实例
	var s = Student{
		Name : "tom",
		Gender : "male",
		Age : 4,
		Id : 1000,
		Score : 66.1,
	}
	info := s.say()
	fmt.Println(info)

	var box Box
	box.len = 1.1
	box. width = 2.0
	box.height = 3.0
	volumn := box.getVolumn()
	fmt.Println("体积为：", volumn)

	var visitor Visitor
	for {
		fmt.Println("请输入您的名字")
		fmt.Scanln(&visitor.Name)
		if visitor.Name == "n" {
			fmt.Println("退出程序")
			break
		}
		fmt.Println("请输入您的年龄")
		fmt.Scanln(&visitor.Age)
		visitor.showPrince()
	}
}
