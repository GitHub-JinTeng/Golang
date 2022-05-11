package login

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"learn/08_Communication_Projects/Commom/message"
	"net"
)

// 写一个函数，完成登录
func Login(userId int, userPwd string) (err error) {
	/*
		一般做项目开发，返回的都用error，因为使用 true和 false不能够充分说明错误的原因是什么，
		比如：登录时，用户名错误 ，密码错误，或者是 网络错误，这些错误必须使用error来描述
	*/

	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return // 如果连拨号连接服务器都连接不上，那么进行下一步代码的处理也没有必要，职业 return 退出
	}
	// **   延时关闭 conn
	defer conn.Close()

	// 2.如果拨号没有报错，则通过 conn 发送消息给服务器
	var mes message.Message         // 定义一个 Message结构体的实例 mes
	mes.Type = message.LoginMesType // 登录的消息		// 这里的 Type 是string，因为常量是string常量

	// 3.创建一个 LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId   // 这里的 UserId 是从上面接收客户端输入的 UserId
	loginMes.UserPwd = userPwd // 这里的 UserPwd 是从上面接收客户端输入的 UserPwd

	// 4.将 loginMes 序列化 （也就是把要发送的 UserId、UserPwd 消息进行序列化后才能发送到服务器）
	data, err := json.Marshal(loginMes) // 此时的 data是一个 byte 切片
	if err != nil {
		fmt.Println("json.Marshal loginMes err=", err)
		return // 客户端数据发送到服务器失败则直接退出
	}

	// 我们需要将序列化后的 byte 类型的 data 进行数据转换成 string类型
	// 5、把 data 赋给 mes.Data 字段
	mes.Data = string(data)
	// 转换成功后，此时的 mes（Message结构体的实例）既有 Type 又有 Data

	// 6、 将 mes 数据进行序列化
	data, err = json.Marshal(mes) // 此时的 data是一个 byte 切片
	if err != nil {
		fmt.Println("json.Marshal mes err", err)
		return
	}

	// 7、到这个时候， data 就是我们要发送的消息了
	// 7.1 先把 data 的长度发送给服务器
	//conn.Write(len(data))  这样是不行的，因为 Write 写入的数据是 byte 类型 ，而 len(data) 是一个int类型，因此会报错

	// 正确做法： 先获取到 data 的长度 len(data)为 int 类型 --> 然后转成一个表示长度的 byte切片，使用(binary.BigEndian)来进行转换
	var pkgLen uint32          // pkgLen 表示这个包的长度（消息的长度）
	pkgLen = uint32(len(data)) // pkgLen为 uint32 类型， 因为 len(byte) 是 int类型，所以需要转换成 uint32
	var buf [4]byte            // 定义一个 byte 切片 ： bytes
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 发送消息长度
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err) // 消息长度没有发送成功
		return                                     // 如果发送消息的长度失败，则直接退出
	}
	fmt.Printf("客户端发送消息的长度成功！！！发送的长度为：%d\n发送的内容为：%s\n", len(data), string(data))
	return
	
}
