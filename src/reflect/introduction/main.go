package main

import (
	"fmt"
	"reflect"
)

//演示基础数据类型的反射
func reflectTest01(b interface{}) {

	//通过反射获取传入变量的type、kind、值
	//1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp =", rTyp) //rTyp = int

	//2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)  //n2= 102
	fmt.Printf("rVal =%v rVal type=%T\n", rVal, rVal) //rVal =100 rVal type=reflect.Value

	fmt.Printf("kind =%v kind=%v\n", rVal.Kind(), rVal.Kind())

	//将rVal转成interface{}
	iV := rVal.Interface()

	//将interface{} 通过断言转成需要的类型
	num := iV.(int)
	fmt.Println("num=", num)
}

type Student struct {
	Name string
	Age int
}

//演示对结构体的反射
func reflectTest02(b interface{}) {

	//通过反射获取传入变量的type、kind、值
	//1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp =", rTyp) //rTyp = main.Student

	//2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal =%v rVal type=%T\n", rVal, rVal) //rVal ={小明 28} rVal type=reflect.Value

	//将rVal转成interface{}
	iV := rVal.Interface()

	//将interface{} 通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

func main() {

	//对基本数据类型、interface{}、reflect.Value进行反射的基本操作

	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name : "小明",
		Age : 28,
	}
	reflectTest02(stu)
}
