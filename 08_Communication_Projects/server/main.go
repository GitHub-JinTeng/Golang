package main

import (
	"fmt"
	"net"
)

//创建一个 处理服务器和客户端的通讯
func process(conn net.Conn) {
	// 这里需要延时关闭 conn ，如果不关闭，则会发生很奇怪的现象
	defer conn.Close()
	// 一直循环读取客户端发送的信息
	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据...")
		n, err := conn.Read(buf[:4]) // 这个 buf[:4] 很重要，表示你发多少个数据，我就读多少个数据，后面的杂数据就不读了
		if n != 4 || err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Printf("读到的buf = %v\n", buf[0:4]) //打印输出时也必须要注意，我们只需要打印客户端发送消息的长度内容即可
	}
}
func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println(" net.Listen err ", err)
		return
	}
	// 延时关闭 listen
	defer listen.Close()
	// 一旦监听成功，就等待客户端来连接服务器
	// 循环的等待客户端来连接
	for {
		fmt.Println("等待客户端来连接服务器....\n")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err", err)
			// 这里不需要 return，因为一个客户端连接不成功，还有其他客户端连接，因此不能关闭
		}

		// 一旦连接成功，则开启一个协程与客户端保持通讯。。
		go process(conn)
	}
}
