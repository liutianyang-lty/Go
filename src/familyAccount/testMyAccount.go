package main

import (
	"fmt"
)

func main() {

	//声明一个变量，保存接收用户输入的选项
	key := ""

	//声明一个变量
	loop := true

	//定义账户的余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定义个变量，标识是否产生了收支
	flag := false


	//收支详情使用字符串记录
	//当有收支时，只需对details进行拼接

	details := "收支\t账户金额\t收支金额\t说	明"


	//显示主菜单
	for {

		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")

		fmt.Scanln(&key)

		switch key {
		case "1" :
			fmt.Println("\n-----------------当前收支明细记录-----------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细... 请来一笔吧")
			}

		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//将这个收入情况，拼接到details变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance,money,note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)

			//做判断
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			//将这个收入情况，拼接到details变量
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance,money,note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗? y/n")
			choice := ""
			for {

				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入 y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")

		}

		if !loop {
			break
		}
	}

	fmt.Println("你退出了家庭记账软件 的使用")
}
