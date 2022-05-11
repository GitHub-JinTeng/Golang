package cal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func AddUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func Sub(num1 int, num2 int) int {
	return num1 - num2
}

// 定义 Monster结构体
type Monster struct {
	Name  string
	Age   int
	Skill string
}

/*
1.	编写一个Monster结构体，字段 Name、Age、Skill
2.	给Monster绑定方法Store ，可以将一个Monster 变量（对象）序列化后保存到文件中
*/
func (m *Monster) Store() bool {
	// 序列化为JSON格式数据
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化失败！", err)
		return false
	}
	//序列化成功，保存到文件中
	// 1.先定义一个文件路径,如果定义的文件路径不存在则创建该文件
	filePath := "G:/ABCD.txt"

	// 2.通过为 ioutil.WriteFile(要写入的文件路径，将什么写到文件，写入的模式)一次性写入文件夹中
	err = ioutil.WriteFile(filePath, data, 0666) //注意这里还又一个模式（这个模式只有在Linux、Unix系统下有效）
	if err != nil {
		fmt.Println("open file err", err)
	}
	// 写入成功
	return true
}

/*
1.	给 Monster绑定方法 ReStore ，可以将一个序列化的 Monster 从文件中读取，并反序列化 Monster 对象，将Store()方法生成的 JSON格式数据 反序列化到文件中
2.	变成测试用例文件 store_test.go ，编写测试用例函数 TestStore 和 TestReStore
*/
func (m *Monster) ReStore() bool {
	// 将Store()方法生成的 JSON格式数据 反序列化到文件中
	// 定义读取的文件路径
	fielPath := "G:/ABCD.txt"
	// 2.通过为 ioutil.ReadFile(读取的文件路径)	一次性读取到文件夹中
	data, err := ioutil.ReadFile(fielPath)
	if err != nil {
		fmt.Println(" read file err", err)
		return false
	}
	// 读取成功
	// 读取成功后，将读取的内容反序列化
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Unmarshal", err)
		return false
	}
	return true
}
