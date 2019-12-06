package main

import (
	"customerManage/model"
	"customerManage/service"
	"fmt"
)

type customerView struct {

	//定义字段
	key string //接受用户输入...
	loop bool //表示是否循环显示主菜单

	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {

	customers := this.customerService.List()

	//显示
	fmt.Println("---------------------客户列表---------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i< len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------------------客户列表完成-------------------\n")
}

//添加客户信息
func (this *customerView) add() {

	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if this.customerService.Add(customer) {

		fmt.Println("---------------------添加完成---------------------\n")

	} else {

		fmt.Println("---------------------添加失败---------------------")
	}

}

//删除客户
func (this *customerView) delete() {

	fmt.Println("-------------------删除客户-------------------")
	fmt.Println("请选择待删除的客户编号(-1退出)：")

	//获取客户ID
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	//提示
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)

	if choice == "y" || choice == "Y" {

		if this.customerService.Delete(id) {
			fmt.Println("-------------------删除完成-------------------")
		} else {
			fmt.Println("-------------------删除失败，输入的id号不存在-------------------")
		}
	}

}

//退出
func (this *customerView) exit() {

	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N)")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}


}

//编辑客户信息
func (this *customerView) update() {

	fmt.Println("---------------------编辑客户---------------------")

	fmt.Println("请选择需要修改的客户编号：")
	id := 0
	fmt.Scanln(&id)

	//先查找是否有该客户
	if this.customerService.FindById(id) == -1 {
		fmt.Println("你选择的客户编号不存在")
		return
	}

	//通过ID获取客户信息
	customer := this.customerService.GetCustomerInfo(id)

	fmt.Printf("姓名(%v):", customer.Name)
	name := ""
	fmt.Scanln(&name)

	fmt.Printf("性别(%v):", customer.Gender)
	gender := ""
	fmt.Scanln(&gender)

	fmt.Printf("年龄(%v):", customer.Age)
	age := 0
	fmt.Scanln(&age)

	fmt.Printf("电话(%v):", customer.Phone)
	phone := ""
	fmt.Scanln(&phone)

	fmt.Printf("电邮(%v):", customer.Email)
	email := ""
	fmt.Scanln(&email)

	//调用Update方法
	if this.customerService.Update(id, name, gender, age, phone, email) {

		fmt.Println("---------------------编辑完成---------------------\n")

	} else {

		fmt.Println("---------------------编辑失败---------------------")
	}
}

//显示主菜单
func (this *customerView) mainMenu() {

	for {
		fmt.Println("------------------客户信息管理软件------------------")
		fmt.Println("					1 添 加 客 户")
		fmt.Println("					2 修 改 客 户")
		fmt.Println("					3 删 除 客 户")
		fmt.Println("					4 客 户 列 表")
		fmt.Println("					5 退		出")
		fmt.Println("请选择(1-5)")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !this.loop {
			break
		}
	}
}

func main() {

	//在主函数中创建一个customerView实例
	customerView := customerView{
		key : "",
		loop : true,
	}

	//这里完成对customerVieW结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}
