package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//	net.Dial("tcp", "192.168.10.4:8888")
	// Dial ： 向服务器拨号
	// "192.168.10.4:8888"	表示在本地电脑地址的8888端口号发起请求
	con, err := net.Dial("tcp", "192.168.10.4:8888")
	if err != nil {
		fmt.Println("Dial err", err)
		// 向服务端发送请求错误，则退出
		return
	}
	//fmt.Println("con 成功", con)
	// 功能1、客户端可以发送单行数据，然后就退出

	// 1.通过 os.Stdin 表示标准输入[终端] ，从终端读取用户输入的数据到 reader变量中；
	reader := bufio.NewReader(os.Stdin)

	for { // 循环读取用户输入的数据

		// 2.从 reader变量中读取数据，读取数据时，遇到换行符['\n']结束
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readeString err", err) // 如果从终端读取用户输入的数据发生错误时，不需要退出
		}
		// 添加退出功能：如果用户输入的数据时 exit 则直接退出
		if line == "exit" {
			fmt.Println("客户退出了...")
			break // 退出 for循环 读取终端数据
		}
		// 在输出的数据时，line的数据自带了 换行符’\n’（有些换行符为：‘\r’），因此我们需要去掉换行字符
		line = strings.Trim(line, " \r\n") //去掉 空格 " " \r  \n
		// 读取数据成功； 将 line 读取到的数据发送到服务器端
		_, err = con.Write([]byte(line + "\n")) // 因为去掉了自带的换行，因此在这我们需要在发送数据到服务器端时加上换行符'\n'
		if err != nil {
			fmt.Println("con.Write err", err) // 发送数据失败也不需要退出
		}
	}
}
