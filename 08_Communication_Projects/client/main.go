package main

import (
	"fmt"
	"learn/08_Communication_Projects/client/login"
	"os"
)

// 1、定义两个全局变量，一个表示用户的Id 另一个表示用户的密码 Pwd
var userId int
var userPwd string

func main() {
	/*
		-----------------欢迎登陆多人聊天系统:-----------------
		     				1登录聊天系统
							2注册用户
		     				3退出系统
		     请选择(1-3):
		----------------------------------------------------
		     1
		     登录..
		     请输入用户id:
				100
		     请输入用户密码:
				200

		     你输入的userid=100 pwd=200
	*/

	// 定义一个变量 key 来接收用户的选项
	var key int
	// 定义一个变量来 判断是否继续显示菜单
	var loop = true

	// 循环显示主菜单
	for loop {
		fmt.Println("-----------欢迎登陆多人聊天系统:------------")
		fmt.Println("\t\t1 登录聊天系统")
		fmt.Println("\t\t2 用户注册")
		fmt.Println("\t\t3 退出系统")
		fmt.Println("\t请选择( 1--3)")

		// 从控制台接收用户的选项 key
		fmt.Scanf("%d\n", &key)
		// 使用 switch 对用户输入的 key 进行判断
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false // 当进入登录聊天室之后就不再显示主菜单了，需要跳转到登录界面进行登录，因此这里需要关闭循环显示主菜单
		case 2:
			fmt.Println("用户注册")
			loop = false // 当进入用户注册之后就不再显示主菜单了，需要跳转到用户注册界面进行注册，因此这里需要关闭循环显示主菜单
		case 3:
			fmt.Println("退出系统")
			loop = false // 退出系统
			os.Exit(0)   // 也可以通过系统来进行关闭
		//Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer的函数不会被执行。
		default:
			fmt.Println("您的输入有误，请重新输入！")
		}
	}
	// 用户的输入，显示新的提示信息
	if key == 1 {
		// 说明用户要登录
		fmt.Println("请输入用户的Id：")
		fmt.Scanf("%d\n", &userId) // 特别注意：这里有个坑，当你使用fmt.Scanf时，在%d后面必须加上换行符 "\n" 否则你一回车
		// 它就会认为你的回车是你输入的密码，因此它将会直接跳过输入密码的环节，但你使用fmt.Scanln 就可以不用再添加换行符 “\n”
		fmt.Println("请输入你的用户密码：")
		fmt.Scanf("%s\n", &userPwd)
		//先把登录函数写到另一个文件，比如 login.go
		err := login.Login(userId, userPwd) //上次在这无法调用 login包，是因为login包的login函数没有首字母大写
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("用户注册的逻辑代码")
	}

}
