package main

import "fmt"

func main() {
	// 功能 1：
	// 声明一个变量接收用户的选项
	var key string
	// 声明一个变量来控制流程的退出,我们先让它一直执行，当满足某个条件时就变成 false，退出软件
	loop := true
	// 功能 2：
	// 定义表头： 收支的详情使用字符串来记录 （ 当有收支时，只需对details进行拼接即可 ）
	details := " 收支\t账户金额\t收支金额\t说 \t\t明"

	// 声明一个变量，当没有任何的收支记录时，就提示这句话 “还没有收支记录，来一笔吧！”
	flag := false // 先定义为false，如果没有收支就直接使用该结果来判断执行输出提示

	//定义账户的余额：
	balance := 10000.0

	// 每次收支的金额
	money := 0.0
	// 每次收支的说明：
	note := ""
	for {
		// 显示主菜单
		fmt.Println("\n-----------------家庭收支记账软件---------------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收支")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退    出")
		fmt.Println()
		fmt.Println("                请选择( 1-4 )")

		fmt.Scanln(&key)
		// 一般在选项应用场景大多都使用 switch
		switch key {

		// 显示收支明细
		case "1":
			fmt.Println("-----------------当前收支明细记录---------------------")

			if flag { // 判断 flag 的值
				fmt.Println(details) // 有收支明细，直接输出" 收支  账户金额   收支金额    说  明 "
			} else {
				fmt.Println("还没有收支记录，来一笔吧！")
			}

		// 登记收入
		case "2":

			fmt.Println("本次收入金额：")
			fmt.Scanln(&money) // 控制台输入收支金额

			balance += money // 每次有收入之后，余额都会发生变化

			fmt.Println("本次收支说明：")
			fmt.Scanln(&note) // 控制台输入本次收支说明

			// 然后将这个收入情况，拼接到 details变量
			details += fmt.Sprintf("\n 收入\t%v\t\t%v\t\t\t%v", balance, money, note)

			//收入成功，有收支信息，把 flag 变成true
			flag = true
		// 登记支出
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			// 这里需要做必要的判断
			if money <= balance {
				//如果支出的金额小于或等于余额时，进行支出处理，处理后显示余额
				balance -= money
			} else {
				// 如果输入的支出金额大于余额，则提示余额不足！
				fmt.Println("您的余额不足！")
				return
			}

			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n 支出\t%v\t\t%v\t\t\t%v", balance, money, note)

			//收入成功，有收支信息，把 flag 变成true
			flag = true

		// 引入退出软件
		case "4":
			fmt.Println("您确定要退出软件吗？ y / n")
			//退出功能也应该做个判断，或许是用户点错了
			choice := ""
			// 如果用户输入的不是 y 或 n 就一直循环让用户输入
			for {
				fmt.Scanln(&choice)
				// 判断用户输入的是不是 y 或 n
				if choice == "y" || choice == "n" {
					break // 输入 y 或 n 则退出该for循环，不再让用户继续输入
				} else {
					fmt.Println("您的输入有误，请重新输入 y / n")
				}
			}
			// 如果用户输入的是 y则退出软件 ，是 n 就继续执行主菜单
			if choice == "y" { // y成立就退出，如果输入的是n,
				loop = false // 这里设定为 false ，外层取反就正好是true，执行退出
			} else if choice == "n" {
				loop = true // 这里设定为 true ，外层取反就正好是false，不执行退出，继续在主菜单中
			}
			// 细节：在这里return只能结束switch判断而已，对for循环不起作用，所以要在switch外return

		// 错误选项
		default:
			fmt.Println("您输入的选项有误！")
		}
		// switch 里面的loop已经为false了，这里的 ！loop 取反是为了让loop变成true让它执行break，结束for循环
		if !loop {
			break
		}
	}
	fmt.Println("您已退出家庭收支记账软件的使用")
}
