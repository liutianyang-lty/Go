package main

import (
	"fmt"
)

//顺序查找
func OrderSort(arr *[4]string, heroName *string) {
	for i := 0; i < len(*arr); i++ {
		if *heroName == (*arr)[i] {
			fmt.Printf("找到%v, 下标为%v", *heroName, i)
			break
		} else if i == (len(*arr) - 1) {
			fmt.Printf("没有找到%v\n",*heroName)
		}
	}
}

func main() {
	//顺序查找
	//案例：有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王
	//猜数游戏：从键盘中任意输入一个名称，判断数列中是否包含此名称【顺序查找】

	//代码
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scan(&heroName)

	//调用顺序查找
	OrderSort(&names, &heroName);
}
