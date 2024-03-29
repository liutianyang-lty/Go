package model

import (
	"fmt"
)

//声明一个customer结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写一个构造方法
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

//不带ID的创建实例的方法
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

//返回用户的信息,格式化
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",this.Id, this.Name, this.Gender,
		this.Age, this.Phone, this.Email)
	return info
}