package main

import "fmt"

func main() {

	//学习二位数组
	//二维数组的演示案例
	/*
	0 0 0 0 0 0
	0 0 1 0 0 0
	0 2 0 3 0 0
	0 0 0 0 0 0
	 */

	//声明一个二维数组
	var arr [4][6]int

	//赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ");
		}
		fmt.Println()
	}

	fmt.Println()

	//二维数组在内存总的布局
	var arr1 [2][3]int
	//初始化
	arr1[1][1] = 10

	fmt.Println(arr1)

	fmt.Printf("arr1[0]的地址为%p\n", &arr1[0])
	fmt.Printf("arr1[1]的地址为%p\n", &arr1[1])

	fmt.Printf("arr1[0][0]的地址为%p\n", &arr1[0][0])
	fmt.Printf("arr1[1][0]的地址为%p\n", &arr1[1][0])

	//二维数组的另一种声明、赋值方式
	var arr2 [2][3]int = [2][3]int{{1,2,3}, {4,5,6}}
	fmt.Println(arr2)


}

