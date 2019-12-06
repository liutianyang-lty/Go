package main

import (
	"fmt"
	"os"
)

func main() {

	//打开一个文件: file叫文件指针或文件句柄
	//注意：文件是一个指针类型
	file, err := os.Open("e:/test100.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//输出文件
	fmt.Printf("file=%v", file) //这里得到的是一个指针

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
