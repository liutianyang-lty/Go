package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改
// num int 的值
// student结构的值

func reflect01(b interface{}) {

	//获取reflect.Value
	rVal := reflect.ValueOf(b)

	//修改
	rVal.Elem().SetInt(20)
}

func reflect02(b interface{}) {

	//获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal Type=%v Kind=%v value=%v\n", rVal.Type(), rVal.Kind(), rVal)

	//转换成interface
	inter := rVal.Interface()

	//转换成float64
	floa := inter.(float64)
	fmt.Println("rVal=", floa)
}

func main() {

	var num int = 10
	reflect01(&num)
	fmt.Println("NUM=", num)

	var v float64 = 1.2
	reflect02(v)

	var str string = "tome"
	fs := reflect.ValueOf(&str) //注意：这里一定要给传入地址
	fs.Elem().SetString("jack")
	fmt.Printf("%v\n",str)
}
