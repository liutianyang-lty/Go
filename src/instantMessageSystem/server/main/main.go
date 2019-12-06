package main

import (
	"fmt"
	"instantMessageSystem/server/model"
	"net"
	"time"
)

//处理和客户端的通讯
func process(coon net.Conn) {

	defer coon.Close()

	//调用总控，创建一个总控
	processor := &Processor{
		Coon:coon,
	}

	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误err=", err)
	}
}

//初始化
func init() {
	//当服务器启动时，就去初始化redis连接池
	initPool("localhost:6379", 16, 0, 300 *time.Second)
	initUserDao()
}

//完成对UserDao的初始化
func initUserDao() {
	//这里的pool 本身就是一个全局变量(在redis.go中已经定义)
	//这里需要注意一个初始化顺序问题
	// initPool ---> initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	//提示信息
	fmt.Println("服务器[新的结构]在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	//监听成功，等待客户端连接
	for {

		fmt.Println("等待客户端连接...")
		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//起一个协程，和客户端保持通讯...
		go process(coon)
	}
}
