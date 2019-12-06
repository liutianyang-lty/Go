package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}

//小学生
type Pupil struct {
	Student //嵌入了Student的匿名结构体
}

//大学生
type Graduate struct {
	Student
}

//将Pupil和Graduate 共有的方法绑定到*Student
//公用方法
func (stu *Student) ShowInfo() {
	fmt.Printf("姓名=%v 年龄=%v 成绩=%v \n", stu.Name, stu.Age, stu.Score)
}

//公用方法
func (stu *Student) SetScore(score int) {
	//业务判断
	stu.Score = score
}

//小学生自己的方法
func (p *Pupil) testing() {
	fmt.Println("小学生正在进行考试......\n")
}

//大学生自己的方法
func (g *Graduate) testing() {
	fmt.Println("大学生正在进行考试......\n")
}

func main() {

	//面向对象编程--继承(解决代码复用的问题)
	//在Golang中，如果一个struct嵌套了另一个匿名结构体，name这个结构体
	//可以直接访问匿名结构体的字段和方法，从而实现了继承特性

	//继承案例：
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(89)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "mary"
	graduate.Student.Age = 23
	graduate.testing()
	graduate.Student.SetScore(100)
	graduate.Student.ShowInfo()
}

