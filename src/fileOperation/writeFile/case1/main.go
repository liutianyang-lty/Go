package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//创建一个新文件，写入5句"hello，Gardon"
	filePath := "e:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	//写入文件
	str := "hello,Gardon \r\n"
	//写入时，使用带缓冲的*Writer
	writer := bufio.NewWriter(file)
	for i:=0; i<5; i++ {
		writer.WriteString(str)
	}

	//此时的内容还在缓冲区中，要把缓冲区的内容给冲洗到磁盘
	writer.Flush()
}
