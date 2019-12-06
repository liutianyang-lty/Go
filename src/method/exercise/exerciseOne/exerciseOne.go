package main

import (
	"fmt"
)

type MethodUtils struct {

}

//给MethodUtils添加方法
func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {

	//练习一
	var mu MethodUtils
	mu.Print()
}

