package utils

import "fmt"

// 创建 FamilyAccount 字段 ,对比面向过程和面向对象的区别
type FamilyAccount struct {

	// 功能 1：显示主菜单 和退出
	// 声明一个字段 key 来接收用户的选项
	key string

	// 声明一个字段  loop 来控制是否退出 for,我们先让它一直执行，当满足某个条件时就变成 false，退出软件
	loop bool

	// 功能 2： 收入明细、支出明细、每次收支说明
	// 声明一个字段 details 来提示" 收支\t账户金额\t收支金额\t说 \t\t明" 收支的详情使用字符串来记录 （ 当有收支时，只需对details进行拼接即可 ）
	details string

	// 声明一个 flag 字段，当没有任何的收支记录时，就提示这句话 “还没有收支记录，来一笔吧！”
	// 记录收支行为；flag := false 先定义为false，如果没有收支就直接使用该结果来判断执行输出提示
	flag bool

	// 定义一个 balance字段来记录账户的余额：
	balance float64

	// 定义一个 money 字段来记录每次收支的金额：
	money float64

	// 定义一个 note 字段来记录每次收支的说明：
	note string
}

// 编写要给工厂模式的构造方法，返回一个FamilyAccount 实例，是为了预防在声明结构体时使用的是首字母小写
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true, // 这里定义时赋true是因为主菜单有个取反操作，所以true取反后变成false，不退出主菜单，只有用户点击退出后才把loop的值改为false
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说 \t\t明",
	}
}

//功能区
// 1)  将收支明细写成一个方法 showDetails() ，和*FamilyAccount 绑定
func (this *FamilyAccount) showDetails() {
	fmt.Println("-----------------当前收支明细记录---------------------")
	if this.flag { // 判断 flag 的值
		fmt.Println(this.details) // 有收支明细，直接输出" 收支  账户金额   收支金额    说  明 "
	} else {
		fmt.Println("当前没有收支明细....来一笔吧！")
	}
}

// 2) 将登记收入写成一个方法 income() ，和*FamilyAccount 绑定
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	// 收入后，余额会产生变化
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	// 将本次收入明细 拼接到 收支明细: details 变量中
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t\t%v", this.balance, this.money, this.note)
	// 收入成后，收支明细有记录，则将 flag 的值改为 true
	this.flag = true
}

// 3) 将登记支出写成一个方法 pay(),和 *FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额为：")
	fmt.Scanln(&this.money)
	// 对支出对一个判断
	if this.money <= this.balance {
		// 支出的金额小于或等于余额，支出成功！
		this.balance -= this.money
	} else {
		fmt.Println("您的余额不足！")
	}
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	// 将本次支出明细 拼接到 收支明细: details 变量中
	this.details += fmt.Sprintf("\n 支出\t%v\t\t%v\t\t\t%v", this.balance, this.money, this.note)
	// 支出成后，收支明细有记录，则将 flag 的值改为 true
	this.flag = true
}

// 4) 将推出功能写成一个方法 exit() ，和 *FamilyAccount 绑定
func (this *FamilyAccount) exit() {
	fmt.Println("您要退出吗？ y / n")
	// 添加退出提示判断，验证用户不是误点了退出键
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		} else {
			fmt.Println("您的输入有误，请重新输入！ y / n ")
		}
	}

	if choice == "y" {
		// 证明用户要退出软件，这里写成false，到最外层的if判断时就取反正好是true，达到执行 break最外层主菜单 for循环
		this.loop = false
	} else if choice == "n" {
		// 证明用户误点了退出键
		this.loop = true //这里写成true，到最外层的if判断时就取反正好是false，就达到不退出最外层主菜单 for循环
	}
}

// 给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		// 显示主菜单
		fmt.Println("\n-----------------家庭收支记账软件---------------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退    出")
		fmt.Println()
		fmt.Println("                请选择( 1-4 )")

		fmt.Scanln(&this.key)
		switch this.key {
		// 收支明细：
		case "1":
			this.showDetails()
		// 登记收入
		case "2":
			this.income()
		// 登记支出
		case "3":
			this.pay()
		// 退出
		case "4":
			this.exit()
		// 当用户选择的选项有误时，提示！
		default:
			fmt.Println("请输入正确选项！")
		}
		// 根据退出功能的用户输入做退出判断，下面将会取 loop的反值
		if !this.loop {
			break // 退出主菜单for循环
		}
	}
}
