package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//打开一个已存在的文件，在原来的内容后追加内容 "ABC!ENGLISH!"
	filePath := "e:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	//写入文件
	str := "ABC!ENGLISH! \r\n"
	//写入时，使用带缓冲的*Writer
	writer := bufio.NewWriter(file)
	for i:=0; i<10; i++ {
		writer.WriteString(str)
	}

	//此时的内容还在缓冲区中，要把缓冲区的内容给冲洗到磁盘
	writer.Flush()
}