package main

import(
	"fmt"
	"time"
)

func main() {

	start := time.Now().Unix()
	for num := 1; num <= 80000; num++ {
		flag := true
		//判断素数
		for i:=2; i<num; i++ {
			if num % i == 0 { //说明不是素数
				flag = false
				break
			}
		}
		if flag {

		}
	}
	end := time.Now().Unix()
	fmt.Println("传统方法耗时=", end - start)
}
