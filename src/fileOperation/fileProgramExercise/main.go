package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount int // 记录英文个数
	NumCount int // 记录数字个数
	SpaceCount int // 记录空格个数
	OtherCount int // 记录其他字符个数
}

func main() {

	//文件编程案例：统计一个文件中含有的英文、数字、空格及其他字符数量
	//思路：
	//先打开一个文件，创建一个Reader
	//以行为单位，进行逐行统计
	//然后将结果保存到一个结构体

	fileName := "e:/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=%v", err)
		return
	}

	defer file.Close()

	//定义一个CharCount 实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		//遍历str
		for _, v := range str {

			switch {
				case v >= 'a' && v <= 'z':
					fallthrough
				case v >= 'A' && v <= 'Z' :
					count.ChCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				default:
					count.OtherCount++
			}
		}
	}

	//输出统计结果
	fmt.Printf("字符个数=%v 数字个数为%v 空格的个数=%v 其他字符个数%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}