package model

import "fmt"

// 创建一个Customer结构体，表示一个客户的信息
type Customer struct {
	Id     int
	Name   string
	Gender string // 性别
	Age    int
	Phone  string
	Email  string
}

// 使用工厂模式，返回一个Customer的实例，只有返回Customer实例后，调用NewCustomer才有效的获取到Customer（客户的信息）
func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}

}

// 使用工厂模式，返回一个Customer的实例，不带 Id的实例，因为Id是唯一的，不能让用户自己输入，让系统分配
func NewCustomer2(name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 定义一个用户信息的格式化输出
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", this.Id,
		this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
