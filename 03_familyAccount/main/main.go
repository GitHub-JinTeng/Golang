package main

import (
	"fmt"
	a "test/FamilyAccount/utils"
)

func main() {
	fmt.Println("这就完成了面向对象的使用")
	// 创建了 FamilyAccount 的指针实例，名称为 NewFamilyAccount;然后通过这个指针实例调用 MainMenu()方法，
	//在 NewFamilyAccount()方法调用时，为什么只看到 Mainmenu()方法的原因是，在 utils包中的方法，只有MainMenu包是首字母大写的，
	// 其他包都是首字母小写的，所以在 NewFamilyAccount()方法调用时看到的可用方法只有 Mainmenu()方法
	a.NewFamilyAccount().MainMenu()
	// MainMenu()方法里面又调用了以下的方法：
	//showDetails() 收支明细方法
	//income()		收入方法
	//pay()			支出方法
	//exit()		退出方法

}
