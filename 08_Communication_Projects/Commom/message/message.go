package message

//定义常量
const (
	LoginMesType    = "LoginMes"    // 登录消息的类型常量	string 类型
	LoginResMesType = "LoginResMes" // 登录返回消息的类型常量
)

// 定义消息类型
type Message struct {
	Type string `json:"type"` // 消息类型	,因为结构体数据序列化后才能发送出去，所以直接序列化
	Data string `json:"data"` // 消息的数据类型
}

// 先定义以下消息..后面需要的时候再增加
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户 Id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名

}

// 定义客户端返回的数据
type LoginResMes struct {
	Code  int    `json:"code"`  // 客户端返回的动态码 ， 500 表示该用户未注册，200表示登录成功
	Error string `json:"error"` // 返回错误信息
}
