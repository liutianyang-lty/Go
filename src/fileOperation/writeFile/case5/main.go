package main

import (
	_ "bufio"
	"fmt"
	_ "io"
	"io/ioutil"
	_ "os"
)

func main() {

	//将e:/abc.txt 文件内容导入到 e:/kkk.txt

	file1Path := "e:/abc.txt"
	file2Path := "e:/kkk.txt"

	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("read file err=%v", err)
		return
	}

	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Println("write file error=%v", err)
	}
}
