package main

import (
	"fmt"
)

func main() {

	//遍历二维数组
	var arr = [2][3]int{{1,2,3}, {4,5,6}}

	//方式一：
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t",arr[i][j])
		}
		fmt.Println()
	}

	//方式二：for-range
	for i, v := range arr {
		for j, v1 := range v {
			fmt.Printf("arr[%v][%v]=%v", i, j, v1)
		}
		fmt.Println()
	}
}

