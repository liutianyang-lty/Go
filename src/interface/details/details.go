package main

import (
	"fmt"
)

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Stu struct {

}

func (stu Stu) test01() {
	fmt.Println("BInterface....")
}

func (stu Stu) test02() {
	fmt.Println("CInterface....")
}

func (stu Stu) test03() {

}

func main() {

	//一个接口(A接口)可以继承多个别的接口(比如B，C接口)，这时如果要是实现A接口，也必须将B,C接口的方法
	//全部实现

	var stu Stu
	var a AInterface = stu
	a.test01()
}
