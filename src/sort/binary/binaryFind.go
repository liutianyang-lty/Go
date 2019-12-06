package main

import (
	"fmt"
)

//二分查找思路: 我们要找的值为target
//1. 有一个有序数组arr
//2. 先找到中间下标 middle = (leftIndex + rigntIndex) / 2, 拿target跟arr[middle]中间值进行比较
//2-1. 如果arr[middle] > target, 说明要找的数在数组的左半区: leftIndex ---> middle-1
//2-2. 如果arr[middle] < target, 说明要找的数在数组的右半区: middle+1 ---> rightIndex
//2-3. 如果arr[middle] = target, 说明找到了
//3. 退出递归调用的条件: leftIndex > rightIndex

//二分查找
func BinaryFind(arr *[7]int, leftIndex int, rightIndex int, target *int) {

	//没有找到
	if leftIndex > rightIndex {
		fmt.Printf("没有找到%v", *target)
		return
	}

	//先找中间下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > *target {
		BinaryFind(arr, leftIndex, middle-1, target)
	} else if (*arr)[middle] < *target {
		BinaryFind(arr, middle+1, rightIndex, target)
	} else {
		fmt.Printf("找到了，下标为%v", middle)
	}
}

func main() {
	target := 0
	//让用户输入一个想要查找的数字
	fmt.Println("请输入一个想要查找的数字...")
	fmt.Scan(&target)
	//学习二分查找
	arr := [7]int{14,23,35,42,57,62,78}
	BinaryFind(&arr, 0, len(arr)-1, &target)
}