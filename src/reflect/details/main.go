package main

func main() {

	//反射的注意事项和细节
	//1. reflect.Value.Kind 获取变量的类别，返回的是一个常量

	//2. Type是类型，Kind是类别，Type和Kind可能是相同的，也可能是不同的
	//比如：var num int = 10 num的Type是int, Kind也是int
	//比如：var stu Student stu的Type是 包名.Student, Kind是struct

	//3. 通过反射来修改变量，当使用SetXxx 方法来设置需要通过对应的指针
	//   类型来完成，这样才能改变传入的变量的值，同时需要使用到reflect.Value.Elem()方法

}

