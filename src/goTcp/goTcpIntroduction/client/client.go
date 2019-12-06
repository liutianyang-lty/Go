package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {

	//client
	coon, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//创建一个reader
	reader := bufio.NewReader(os.Stdin)

	//循环进行用户输入
	for {

		//从终端读取一行
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		//判断是否退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Printf("client %s exit", coon.RemoteAddr().String())
			break
		}

		//将数据发送给服务端
		_, err = coon.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("coon.Write err=", err)
		}
	}
}
