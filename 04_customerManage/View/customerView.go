package main

import (
	"fmt"
	model "test/customerManage/Model"
	service "test/customerManage/Service"
)

// 因为customerView 是在本包使用，所以首字母可以不大写
type customerView struct {
	// 定义必要的字段
	key  string // 接收用户的输入...
	loop bool   //表示是否退出主菜单
	// 为了能调用 Service包中的 customerService实例里面的方法，我们添加一个 customerService的字段 ,类型为：*CustomerService 指针
	customerService *service.CustomerService
}

//获取用户输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("-----------------------添加客户-------------------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("姓别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮件：")
	email := ""
	fmt.Scanln(&email)
	// 构建一个新的 Customer实例
	// 注意：Id号，没有让用户自己输入，Id是唯一的，需要让系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Print("-----------------------添加成功-------------------------")
	} else {
		fmt.Print("-----------------------添加失败-------------------------")
	}
}

//获取用户输入的Id，删除该Id对应的客户
func (this *customerView) delete() {
	fmt.Println("---------------------删除客户---------------")
	fmt.Println("请选择删除的客户编号（-1退出）：")
	//  先默认将id 的值赋为 -1 ，如果此时客户直接按回车，不做任何操作时，id=-1，表示不退出，如果客户
	//要退出，则输入 -1 ，
	id := -1
	fmt.Scanln(&id)
	// 判断用户是否放弃删除
	if id == -1 {
		return //放弃删除客户
	}
	// 输入的是某个id，则提示用户确认是否删除
	fmt.Println("确认是否删除？ y / n")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		// 调用 customerService 的 Delete()方法删除用户输入的 id
		if this.customerService.Delete(id) {
			fmt.Println("---------------------删除成功---------------\n")
		} else {
			fmt.Println("--------------删除失败，输入的Id不存在---------------\n")
		}
	} else {
		return
	}
}

// 因为所有用户的信息都在Service包中的 List()方法中，
//所以需要创建一个方法 list()方法来调用 Service包里面的 List()方法来获取所有客户的信息
func (this *customerView) list() {
	//首先，获取 Service包的 List()方法中存在的客户信息
	customers := this.customerService.List()
	// 显示
	fmt.Println("----------------------客户信息列表----------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	// 遍历获取到的用户信息
	for i := 0; i < len(customers); i++ {
		// 传统方式输出：fmt.Printf("Id%v\tName%v\t....") 这样太不美观了
		// 使用调用封装的方式来输出,customers[i]表示切片中的每个用户的信息，GetInfo()表示封装好了的用户信息的格式化输出
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------------------------------------------------------------\n\n")
}

func (this *customerView) exit() {
	fmt.Println("确认是否退出？ y / n")
	for {
		choice := ""
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			this.loop = false
			fmt.Println("您已退出系统！")
			break
		} else if choice == "n" || choice == "N" {
			this.loop = true
			fmt.Println("您已取消退出系统！")
			break
		} else {
			fmt.Println("您的输入有误，请重新输入！ y / n ")
		}

	}

}

// 显示主菜单
func (this *customerView) mainMenu() {
	//显示主菜单使用for循环
	for {
		fmt.Println("----------------------客户信息管理软件----------------------------")
		fmt.Println("                        1.添加客户")
		fmt.Println("                        2.修改客户")
		fmt.Println("                        3.删除客户")
		fmt.Println("                        4.显示客户")
		fmt.Println("                        5.退出系统")
		fmt.Println("                       请选择（1-5）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			// 添加客户
			this.add()
		case "2":
			fmt.Println("2.修改客户")
		case "3":
			// 删除客户
			this.delete()
		case "4":
			// 显示所有用户信息
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("您的选项有误，请重新输入！")

		}
		// 将 loop 取反来控制是否退出
		if !this.loop {
			break
		}
	}

}
func main() {
	goodsSvc :=

	// 在main函数中创建一个customerView实例，给结构体的字段赋值，并且运行显示主菜单
	customerView := customerView{
		key:  "",
		loop: true, // 赋值为true 是先让它一直循环主菜单
	}
	// 上面在customerView结构体中声明了 customerService 字段，那么在这里完成 customerService 字段的初始化
	// customerService 调用 NewCustomerService()业务处理的工厂模式
	customerView.customerService = service.NewCustomerService()
	service.NewOrderService(customerView.customerService)
	// 调用显示主菜单的方法
	customerView.mainMenu()
}
