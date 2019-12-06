package main

import (
	"fmt"
	"oop/oopMid/exercise/model"
)

func main() {
	account := model.NewAccount("jzh11111", "999999", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}

}