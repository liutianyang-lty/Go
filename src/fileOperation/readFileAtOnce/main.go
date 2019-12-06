package main

import (
	"fmt"
	"io/ioutil"
	"mahonia"
)

func main() {

	var dec mahonia.Decoder
	dec = mahonia.NewDecoder("gbk")

	//使用ioutil.ReadFile 一次性读取文件
	file := "e:/test.txt"
	content, err := ioutil.ReadFile(file)  //只适合小文件读取
	if err != nil { //说明出错了
		fmt.Printf("read file err=%v", err)
	}

	//显示读取到的内容
	fmt.Printf("%v", dec.ConvertString(string(content)))
}
