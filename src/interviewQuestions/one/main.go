package main

type student struct {
	Name string
	Age int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name:"zhou", Age:31},
		{Name:"li",Age:23},
		{Name:"wang", Age:22},
	}

	//for _, stu := range stus {
	//	m[stu.Name] = &stu //这样的循环是不对的
	//}

	//正确的循环赋值方法：
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

func main() {

	//调用pase_student方法
	//pase_student() //返回结果为map[li:0xc0000044a0 wang:0xc0000044a0 zhou:0xc0000044a0]

	pase_student()
}
