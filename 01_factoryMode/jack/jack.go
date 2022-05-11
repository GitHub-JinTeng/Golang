//package jack
//
//import "fmt"
//
//// 1、创建 account 结构体
//type account struct {
//	accountNo string
//	balance   float64
//	pwd       string
//}
//
//// 2、创建 account工厂模式 让外部可调用
//func NewAccount(AccountNo string, Balance float64, Pwd string) *account {
//	return &account{
//		accountNo: AccountNo,
//		balance:   Balance,
//		pwd:       Pwd,
//	}
//}
//
//// 3、 分别对 account 字段进行 Set() 与 Get() 让外部可以访问account字段
//// 3.1 账号  Set()
//func (a *account) SetAccountNo(AccountNo string) {
//	// 添加账号判断
//	if len(AccountNo) > 6 && len(AccountNo) < 10 {
//		a.accountNo = AccountNo
//	} else {
//		fmt.Println("您输入的账号有误！请输入长度为6~10之间")
//		return
//	}
//}
//
//// 3.2账号 Get()
//func (a *account) GetAccountNo() string {
//	return a.accountNo
//}
//
//// 3.3 余额 Set()
//func (a *account) SetBalance(Balance float64) {
//	// 判断余额
//	if Balance > 20 {
//		a.balance = Balance
//	} else {
//		fmt.Println("您输入的余额有误！存款余额必须大于20")
//		return
//	}
//}
//
//// 3.4余额 Get()
//func (a *account) GetBalance() float64 {
//	return a.balance
//}
//
//// 3.5 密码 Set()
//func (a *account) SetPwd(Pwd string) {
//	//密码判断
//	if len(Pwd) == 6 {
//		a.pwd = Pwd
//	}
//}
//
//// 3.6 密码 Get()
//func (a *account) GetPwd() string {
//	return a.pwd
//}
package jack

import "fmt"

// 定义person结构体
type person struct {
	name string
	age  int
	sal  float64
}

// 创建工厂模式函数
func NewPerson(Name string, Age int, Sal float64) *person {
	return &person{
		name: Name,
		age:  Age,
		sal:  Sal,
	}
}

//为了外部可访问，创建 Set() 与 Get() 函数
func (p *person) SetName(Name string) {
	// 添加判断
	if len(Name) >= 2 && len(Name) <= 4 { //len是按字节来算的，所以输入2个中文是会报错的，一个中文==3个字节
		p.name = Name
	} else {
		fmt.Println("您输入的名字有误！")
		return
	}
}

// 返回 name
func (p *person) GetName() string {
	return p.name
}

func (p *person) SetAge(Age int) {
	// 添加判断
	if Age <= 0 || Age > 150 {
		fmt.Println("您输入的年龄有误！")
		return
	} else {
		p.age = Age
	}
}

// 返回年龄
func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(Sal float64) {
	//添加判断.
	if Sal < 3000 || Sal > 50000 {
		fmt.Println("您输入的薪水有误！")
	} else {
		p.sal = Sal
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
