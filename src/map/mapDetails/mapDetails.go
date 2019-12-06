package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[1] = 99
}

//定义结构体
type Stu struct {
	Name string
	Age int
	Address string
}

func main() {

	//1. map是引用类型，准守引用类型传递的机制，在一个函授站接受map，修改改后，会直接修改原来的map
	map1 := make(map[int]int)
	map1[1] = 90
	map1[3] = 89
	map1[20] = 12
	modify(map1)
	fmt.Println(map1)

	//2. map的容量达到后，再想map增加元素，并不会发生panic，也就是说map能动态增长

	//3. map的value值经常使用struct类型，更适合管理复杂的数据（比用map更好）

	//map的value为struct类型 案例：
	students := make(map[string]Stu, 10)
	//创建两个学生
	stu1 := Stu{
		Name:"tome",
		Age:18,
		Address:"北京",
	}
	stu2 := Stu{
		Name:"mary",
		Age:20,
		Address:"上海",
	}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	//遍历各个学生的信息
	for k, v := range students {
		fmt.Printf("学号：%v, 名字是：%v, 年龄是：%v, 地址是：%v\n", k, v.Name, v.Age, v.Address)
	}
}

