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

func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//判断奇偶数
func (mu *MethodUtils) JudgeNum(num int) {
	if num % 2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Print("是奇数")
	}
}

//打印
func (mu *MethodUtils) Print3(n int, m int, key string) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func main() {

	//练习一
	var mu MethodUtils
	mu.Print()
	fmt.Println()
	mu.Print2(2, 6)

	mu.JudgeNum(11)
	fmt.Println()
	mu.Print3(3, 5, "#")
}
