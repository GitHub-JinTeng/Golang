package service

import (
	"fmt"
	model "test/customerManage/Model"
)

// 该 CustomerService，完成对 Model包中的Customer结构体的操作，包括 “ 增删改查  ”
type CustomerService struct {

	//创建一个customers字段的切片来存放用户信息
	customers []model.Customer

	// 创建用来存放客户的Id
	customerNum int
}

// 编写一个方法，可以返回 *CustomerService 的指针类型，利于对增删改查操作
func NewCustomerService() *CustomerService {

	// 创建 CustomerService结构体的实例，然后添加一个客户信息
	customerService := &CustomerService{}
	customerService.customerNum = 1 // 一个客户信息
	// 创建一个客户信息
	customer := model.NewCustomer(1, "张三", "男", 30, "010-5678", "zhangsan@abc.com")
	// 将客户信息追加到客户的切片 customers 中
	customerService.customers = append(customerService.customers, customer)

	return customerService
}

// 声明一个List()方法，和 CustomerService 绑定，来返回客户信息的切片
func (this *CustomerService) List() []model.Customer {

	return this.customers
}

// 添加客户到 customer切片
// Add的参数 customer 参数类型为：切片 Customer
func (this *CustomerService) Add(customer model.Customer) bool {
	// 我们确定一个分配 id的规则，就是添加的顺序
	//因为 Id是唯一的，不需要用户自己输入，让系统自己分配，那么系统怎么分配？使用customerNum++当做系统分配的Id
	this.customerNum++ // customerNum 从定义的 1开始自增
	// 把每次的customerNum赋给Id
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 删除客户信息
func (this *CustomerService) Delete(id int) bool {
	// 从下面的 FindById()方法查找Id的结果 赋给 index ，如果在FindById()方法找到Id，
	//则 index ！= -1 ，如果没有找到，则 index = -1
	index := this.FindById(id)
	// 对 index 进行判断
	if index == -1 {
		return false
	}
	// 如何从切片删除一个元素
	// 假如：index = 4 ,
	// [:index]== [0:3] ,[index+1:]==[4:] ？？ 这不就 == [:]
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

// 根据id查找客户在切片中的对应下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历用户信息 this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		// 如果 customers切片的当前客户 Id == 用户输入的删除id,则表示找到要删除的客户
		if this.customers[i].Id == id {
			// 找到Id后就把该用户的Id赋给index
			index = i
		} else {
			// 此时 index = -1
			index = -1
		}
	}
	return index
}

func (this *CustomerService) Uppdate(id int, name string) {
	fmt.Printf("请输入你要修改的Id客户：")
	idx := this.FindById(id)
	if idx < 0 {
		// todo 打印没有该用户
		return
	}
	old := this.customers[idx].Name
	this.customers[idx].Name = name
	fmt.Printf("命名修改：%s -> %s \n", old, name) // 旧值->新值
	//index := -1
	////遍历用户信息 this.customers 切片
	//for i := 0; i < len(this.customers); i++ {
	//	// 如果 customers切片的当前客户 Id == 用户输入的删除id,则表示找到要删除的客户
	//	if this.customers[i].Id == id {
	//		// 找到Id后就把该用户的Id赋给index
	//		index = i
	//	} else {
	//		// 此时 index = -1
	//		index = -1
	//	}
	//	if index == -1{
	//		fmt.Println("您输入的客户Id不存在，请重新输入！")
	//	}else{
	//		this.customers[i] = model.NewCustomer2(name,gender,age,phone,email)
	//	}
	//}

}
