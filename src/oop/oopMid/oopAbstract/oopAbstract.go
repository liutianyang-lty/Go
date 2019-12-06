package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

//方法：存款
func (account *Account) Deposite(money float64, pwd string) {

	//对密码进行验证
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功~~")
}

//方法：取款
func (account *Account) WithDraw(money float64, pwd string) {

	//对密码进行验证
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你取出的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功~~")
}

//方法：查询余额
func (account *Account) Query(pwd string) {

	//对密码进行验证
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为=%v 余额=%v", account.AccountNo, account.Balance)

}

func main() {

	//面向对象编程思想--抽象
	//如何理解抽象
	//把一类事物的共有的属性(字段)和行为(方法)提取出来，形成一个物理模型(模板)。这种研究问题的方法称为抽象
	//案例：创建一个银行账号结构体：账号、密码、余额；可以存款、取款、查询余额

	//测试
	account := Account{
		AccountNo: "gs1111111",
		Pwd: "666666",
		Balance: 100.0,
	}
	account.Query("666666")
	account.Deposite(200.00, "666666")
	account.Query("666666")

	account.WithDraw(400.00, "666666")
	account.Query("666666")
}
