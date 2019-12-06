package main

import (
	"bufio"
	"fmt"
	"io"
	_"mahonia"
	"os"
)

func main() {

	//var dec mahonia.Decoder
	//dec = mahonia.NewDecoder("gbk")

	//打开一个存在的文件，将原来的内容独处显示在终端，并且追加5句"hello，北京！"
	//读写+追加
	filePath := "e:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	//先读取原来文件的内容，并显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Print(str)
	}

	//写入文件
	str := "hello，北京！\r\n"
	//写入时，使用带缓冲的*Writer
	writer := bufio.NewWriter(file)
	for i:=0; i<5; i++ {
		writer.WriteString(str)
	}

	//此时的内容还在缓冲区中，要把缓冲区的内容给冲洗到磁盘
	writer.Flush()
}