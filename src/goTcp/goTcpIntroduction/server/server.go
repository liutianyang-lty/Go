package main

import (
	"fmt"
	"net"
)

func process(coon net.Conn) {

	//创建一个切片
	buf := make([]byte, 1024)

	defer coon.Close()

	//循环接收客户端发送的信息
	for {
		n, err := coon.Read(buf)
		if err != nil {
			fmt.Println("coon Read err=", err)
			return
		}

		//显示信息
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	//Go--tcp编程
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		//监听失败
		fmt.Println("listen err=", err)
		return
	}

	defer listen.Close()

	//循环等待客户端的连接
	for {
		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept success, clientIp=%v\n", coon.RemoteAddr().String())
		}

		//起一个协程服务于客户端
		go process(coon)
	}

}