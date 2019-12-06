package main

import(
	"fmt"
	"flag"
)

func main() {

	//使用flag包解析命令行参数
	//定义几个变量，用于接收输入的参数
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用命令行中输入的 -u 后面的参数值
	//"u", 就是 -u 指定参数
	//"", 默认值
	//最后一个参数为一个说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认3306")

	flag.Parse()

	//输出结果
	fmt.Printf("user：%v pwd：%v host：%v port：%v", user, pwd, host, port)
}
