package main

import (
	"familyAccount/oopAccount/utils"
	"fmt"
)

func main() {

	fmt.Println("这个是面向对象的完成~~")
	utilsObj := utils.NewFamilyAccount()
	utilsObj.MainMenu()
}