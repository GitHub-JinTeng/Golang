package model

import "fmt"

// 创建一个员工结构体
type person struct { //这里定义时首字母先小写，允许外部的包调用
	Name string
	age  int     // 结构体字段变小写，其他包不能直接访问
	sal  float64 // 工资
}

// 写一个工厂模式的函数（相当于别语言的构造函数） 使外部可调用该结构体
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了能够访问 age 和 sal 我们编写一个SetXxx()的方法 和 GetXxx()方法
// 创建一个SetXxx()方法 “ 用于对属性判断并赋值 ”
func (p *person) SetAge(age int) { // 这里不需要写 Name，因为它是大写的，可以直接调用
	// 添加对属性判断并赋值
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("您输入的年龄有误！")
	}
}

//创建一个GetXxx()方法 “ 用于获取属性的值 ”
func (p *person) GetAge() int {
	return p.age // 将年龄返回给调用者
}

// 为了能够访问 sal 和 sal 我们编写一个SetXxx()的方法 和 GetXxx()方法
// 创建一个SetXxx()方法 “ 用于对属性判断并赋值 ”
func (p *person) SetSal(sal float64) {
	// 添加判断业务逻辑
	if sal < 3000 || sal > 50000 {
		fmt.Println("您输入的薪水有误！")
	} else {
		p.sal = sal
	}
}

//创建一个GetXxx()方法 “ 用于获取属性的值 ”
func (p *person) Getsal() float64 {
	return p.sal
}
