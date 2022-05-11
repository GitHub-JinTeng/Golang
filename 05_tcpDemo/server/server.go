package main

import (
	"fmt"
	"io"
	"net"
)

// 创建一个函数来处理客户端的数据
func process(con net.Conn) {
	// 首先先延迟关闭 con
	defer con.Close() //关闭 con
	//循环接收客户端发送的数据
	for {
		// 每次接收客户端发送数据的 10个字节,存放到 buf 变量中
		buf := make([]byte, 10)
		//con.Read(buf)
		// 1.等待客户通过con 发送信息
		// 2.如果客户端没有Write[发送数据]，那么协程就阻塞在这里
		// 3.通过 con.RemoteAddr().String()获取远程客户端的地址
		//fmt.Printf("服务器在等待客户端%s,发送信息\n", con.RemoteAddr().String()) //看起来太乱了

		// 读取客户端发送的数据
		n, err := con.Read(buf) // 从 con 读取数据到 buf中
		//判断读取是否成功
		if err != nil || err == io.EOF {
			fmt.Println("客户端已退出...", err)
			// 如果 服务器读取客户端的数据失败时必须退出，否则会发生死循环
			break
		}
		// 显示客户端发送的内容到服务器的终端
		// 细节： 1、因为是一行一行的读取，每当读取到‘\n’就结束，进行下一轮的读取
		//	2、因为每次buf的缓冲区是10，当用户输入7个字节的数据时，我们只需要输出7个字节的数据就行了，后面的就不再输出，因此使用[:n],读取到客户端发送的最后一个数据
		fmt.Printf("客户端发送的数据为：%v", string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	// Listen ：监听
	// Tcp ：表示网络协议是 tcp
	// 0.0.0.0:8888 表示在本地监听 8888 端口号
	listen, err := net.Listen("tcp", "0.0.0.0:8888") // listen 是接口类型
	if err != nil {
		fmt.Println("listen err", err)
		// 监听服务器错误则退出
		return
	}
	//循环等待客户端来连接
	for {
		fmt.Println("等待客户端来连接...")
		con, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err ", err)
			// 当客户端连接服务器发生错误时，并不需要关闭服务器，否则导致可以连接服务器的客户端也被迫关闭，因此不需要关闭服务器
		} else {
			// 客户端成功连接到服务器端时，打印输出客户端输入的数据，并且使用con.RemoteAddr()打印输出客户端的远程地址
			fmt.Printf("客户端:%v 发送的数据为：%v\n", con.RemoteAddr().String(), con)
		}
		// 开启一个协程，让协程为客户端服务
		go process(con)
	}

}
