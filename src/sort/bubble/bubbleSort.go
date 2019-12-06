package bubble

import (
	"fmt"
)

//冒泡排序
func BubbleSort(arr *[6]int) {

	temp := 0
	for i := 0; i < len(*arr) - 1; i++ { //外层循环

		for j := 0; j < len(*arr) - 1 - i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func main() {
	//定义一个数组
	arr := [6]int{12,3,34,56,80,18}
	fmt.Println("排序前arr=", arr)

	//调用冒泡排序
	BubbleSort(&arr)
	fmt.Println("排序后arr=", arr)
}