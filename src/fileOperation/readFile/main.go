package main

import (
	"bufio"
	"fmt"
	"io"
	"mahonia"  //编码转换
	"os"
)

func main() {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")

	//读取文件的案例
	//读取文件的内容并显示在终端，使用os.Open, file.Close, bufio.NewReader(), reader.ReadString
	file, err := os.Open("e:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时关闭file
	defer file.Close() //防止内存泄露

	//创建一个 *Reader , 是带缓冲的, 默认缓冲区为4096个字节
	reader := bufio.NewReader(file)

	//循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Println("UTF-8 to GBK:", enc.ConvertString(str))
	}

	fmt.Println("文件读取结束")
}
