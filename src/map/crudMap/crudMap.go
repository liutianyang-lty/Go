package main

import (
	"fmt"
)

func main() {

	//学习map的增删改查操作
	//map的新增和更新：map[key] = value 如果key还没有，就是增加；如果key存在，就是更新
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"
	fmt.Println(cities)

	//更新
	cities["no3"] = "杭州"
	fmt.Println(cities)

	//删除：delete()
	delete(cities, "no1")
	fmt.Println(cities)

	//细节说明：
	//(1)如果我们要删除map的所有key, 没有办法一次性删除，只能进行遍历删除
	//(2)直接make一个新的空间（如果想一次性删除map中的所有key，建议使用这种操作）

	//查询
	val,ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}
}
