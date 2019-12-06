package main

import (
	"fmt"
)

func main() {

	//map的遍历
	//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	for key, val := range cities {
		fmt.Printf("key=%v val=%v\n", key, val)
	}

	//遍历二维map
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string)
	studentMap["stu01"]["name"] = "bob"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "杭州市拱墅区"

	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "tony"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "杭州市滨江区"

	for k1, v1 := range studentMap {
		fmt.Printf("k1=%s \n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%s v2=%s \n", k2, v2)
		}
		fmt.Println()
	}

	//map的长度
	fmt.Println(len(studentMap))
}
