package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//封装成一个函数，接收两个文件路径
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) { //返回的第一个参数是复制的字节数

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}

	defer srcFile.Close()

	//通过srcFile文件句柄 获取reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	writer := bufio.NewWriter(dstFile)

	defer dstFile.Close()
	return io.Copy(writer, reader)

}

func main() {

	//copy文件
	//将d:/flower.jpeg文件copy到 e:/cat.jpeg

	srcFile := "d:/flower.jpeg"
	dstFile := "e:/cat.jpeg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("copy完成！")
	} else {
		fmt.Println("copy file err=%v", err)
	}

}
