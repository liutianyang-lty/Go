package main

import(
	"fmt"
)

type Student struct {
	Name string
	Age int
}

func (stu *Student)  String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {

	//方法使用的深度剖析
	//func (receiver type) methodName (参数列表) (返回值列表) {
	// 		方法体
	//		return 返回值
	// }

	//方法注意事项和细节讨论
	//1. 结构体类型是值类型，在方法调用中，准守值类型的传递机制，是值拷贝传递方式
	//2. 如果程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
	//3. Golang中的方法作用在指定的数据类型上的，因此自定义类型，都可以有方法，而不仅仅是struct
	//4. 方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包中访问，方法首字母大写，可以在本包和其它包访问
	//5. 如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出

	//定义一个student变量
	stu := Student{
		Name : "tom",
		Age : 20,
	}
	fmt.Println(&stu) //结果为：Name=[tom] Age=[20]
	fmt.Println(stu) //结果为：{tom 20}
}
