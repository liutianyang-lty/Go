package service

import (
	"customerManage/model"
)

//该结构体完成对customer的操作，包含增删改查
type CustomerService struct {

	//定义一个切片保存客户信息
	customers []model.Customer

	//声明一个字段，表示当前切片含有多少个客户
	//后期，该字段可以作为新客户的ID
	customerNum int
}

//编写一个方法，返回*CustomerService
func NewCustomerService() *CustomerService {

	customerService := &CustomerService{

	}

	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@souhu.com")
	customerService.customers = append(customerService.customers, customer)

	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {

	return this.customers
}

//添加客户
func (this *CustomerService) Add(customer model.Customer) bool {

	//确定一个分配ID 的规则
	this.customerNum++

	customer.Id = this.customerNum

	this.customers = append(this.customers, customer)
	return true
}

//删除客户
func (this *CustomerService) Delete(id int) bool {

	index := this.FindById(id)
	if index == -1 {
		return false
	}

	//从切片中删除元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true

}

//根据ID获取客户在切片中对应的下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {

	//遍历
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//找到
			index = i
		}
	}
	return index
}

//根据id获取客户的信息
func (this *CustomerService) GetCustomerInfo(id int) model.Customer {

	//遍历
	var customer model.Customer

	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			customer = this.customers[i]
		}
	}
	return customer
}

func (this *CustomerService) Update(id int, name string, gender string, age int, phone string, email string) bool {

	//根据ID进行更新数据
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			if name != "" {
				this.customers[i].Name = name
			}
			if gender != "" {
				this.customers[i].Gender = gender
			}
			if age != 0 {
				this.customers[i].Age = age
			}
			if phone != "" {
				this.customers[i].Phone = phone
			}
			if email != "" {
				this.customers[i].Email = email
			}

		}
	}

	return true
}