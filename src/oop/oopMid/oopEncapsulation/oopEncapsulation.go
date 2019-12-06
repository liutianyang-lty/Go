package main

import (
	"fmt"
)

func main() {

	//面向对象编程--封装
	//封装就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在内部，
	//程序的其他包只能通过被授权的操作(方法),才能对字段进行操作

	//封装的理解和好处：
	//1.隐藏实现细节
	//2.可以对数据进行验证，保证安全合理

	//如何体现封装
	//1.对结构体中的属性进行封装
	//2.通过 方法、包 实现封装

	//封装实现步骤：
	//1. 将结构体、字段的首字母小写，其他包不能使用
	//2. 给结构体所在的包提供一个工厂模式的函数，首字母大写
	//3. 提供一个首字母大写的Set方法，用于对属性判断并赋值
	//4. 提供一个首字母大写的Get方法，用于获取属性的值

	//案例：不能随便查看人的年龄、工资等隐私，并对输入的年龄进行合理的验证。
}