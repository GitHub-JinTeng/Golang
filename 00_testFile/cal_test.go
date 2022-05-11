package cal

import (
	"testing"
)

func TestAddUpper(t *testing.T) {

	//调用 AddUpper函数并传入实参 ，然后创建一个变量接收该函数的返回值
	result := AddUpper(10)
	//获取该函数运行的结果后，对结果进行判断，是否和预期的结果（正确的结果）一致
	if result != 55 {
		t.Fatalf("AddUpper(10)函数执行错误,执行的结果=%v,预期的结果=%v", result, 55) // 这是函数的结果不等于 55 时返回的错误信息
	}
	//这是函数的结果等于预期结果： 55 时返回的正确信息
	t.Logf("AddUpper(10)函数执行正确....")
}
func TestSub(t *testing.T) {
	result := Sub(10, 3)
	if result != 7 {
		t.Fatalf("Sub(10,3)函数用例执行错误！期望值%v结果为%v", 7, result)
	}
	t.Logf("Sub(10,3)函数用例执行正确！")
}

// 测试用例；用于测试 Store方法
func TestStore(t *testing.T) {
	// 测试 Store()方法 创建的 JSON 格式数据 是否成功
	monster := Monster{
		// 创建 Monster 结构体实例
		Name:  "孙悟空",
		Age:   500,
		Skill: "火眼金睛",
	}
	// 测试函数，通过 结构体( Monster )的 实例对象( monster )进行调用它对应的方法 Store() ,并且创建一个变量 result来接收该方法的返回值
	result := monster.Store()
	if !result {
		t.Fatalf("monster.Store()方法错误，期望值为：%v结果为：%v", true, result)
	}
	// 该测试正确；为了证明该方法测试用例 和 该测试的文件( 内容 )是否一致，我们进一步进行判断
	t.Logf("Store()方法执行正确！")
	//该测试用例一执行，就会重新将Store()序列化的JSON数据格式的内容重新 “ 清空并写入”到指定的文件夹中
}

// 测试用例；测试 ReStore方法
func TestReStore(t *testing.T) {
	// 测试 ReStore()方法 将 Store()方法生成的 JSON格式数据 反序列化到文件中 是否成功
	// 1. 创建 结构体( Monster ) 的空实例对象( monster )来接收反序列化后的结果
	var monster Monster

	// 测试函数，通过 结构体( Monster )的 实例对象( monster )进行调用它对应的方法 ReStore() ,并且创建一个变量 result来接收该方法的返回值
	result := monster.ReStore()
	// 因为 ReStore()方法的返回值是bool类型
	if result != true {
		t.Fatalf("monster.ReStore()方法错误，期望值为：%v结果为：%v", true, result)
	}
	// 为了证明该测试的文件是否和 测试用例的JSON文件一致，所以我们做进一步判断
	if monster.Name != "孙悟空" {
		t.Fatalf("monster.ReStore()方法错误，期望值为：%v结果为：%v", "孙悟空", monster.Name)
	}
	t.Logf("monster.ReStore()方法测试成功！")
}

/*为了证明该测试的文件是否和 测试用例的JSON文件一致，所以我们做进一步判断
执行这个if monster.Name != "孙悟空" 判断时，必须是先执行Store()的测试，将Store()方法生成的JSON的数据格式保存到文件中，
如果我们需要将JSON中的某个字段的值给修改后 进行错误的测试时，我们必须单独执行ReStore()的测试，
如果Store()的测试和ReStore()的测试一起执行的话，在Store()测试后会重新将Store()序列化的JSON数据格式的内容重新 “ 清空并写入”到指定的文件夹中，
你就无法获取到错误的值来进行判断！*/
//代码带解析如下：
/*
//调用 AddUpper函数并传入实参 ，然后创建一个变量接收该函数的返回值
	result := AddUpper(10)

//获取该函数运行的结果后，对结果进行判断，是否和预期的结果（正确的结果）一致
	if result != 55 {

//原始的输出方式为：
    fmt.Printf("AddUpper(10)函数执行错误,执行的结果=%v,预期的结果=%v",result,55) 这种方式输出，耗时更长，性能更低。

//单元测试中输出方式：	（输出结果需要使用该函数的参数 t 来调用该参数的方法）
	1. t.Fatalf( ) 执行错误 使用 Fatalf ，Fatal：致命的（表示错误） ，f 表示格式化输出的意思
	2. t.Fatalf( ) 执行正确 使用 Logf ， Log ： 日志 （表示正确） ， f 表示 格式化输出的意思

    // 这是函数的结果不等于 55 时返回的错误信息
        t.Fatalf("AddUpper(10)函数执行错误,执行的结果=%v,预期的结果=%v", result, 55)
	}

	//这是函数的结果等于预期结果： 55 时返回的正确信息
        t.Logf("AddUpper(10)函数执行正确....")
/*

使用 testing 框架进行单元测试的流程
1.	要测试的文件命名规范为：要测试的文件名 + _test.go结尾
2.	声明一个函数，该函数为要测试的函数，函数命名规范是：Test+函数名(首字母必须是大写) ,
	该函数的参数和参数类型必须为 参数名：t  参数类型： *testing.T
3.	在单元测试中输出结果需要使用该函数的参数 t 来调用该参数的方法，
	1.执行错误 使用 Fatalf ，Fatal：致命的（表示错误） ，f 表示格式化输出的意思
	2.执行正确 使用 Logf ， Log ： 日志 （表示正确） ， f 表示 格式化输出的意思

//在终端运行单元测试的代码时，必须使用
1. 	go test - v 						输出 执行正确的函数 和 执行错误的函数
2.	go test								只输出 执行错误的函数
3.	go test -v cal_test.go 				只测试单个测试文件中的所有函数，并输出 执行正确的函数  										 和执行错误的函数
4.	go test -v -test.run  TestAddUpper	只测试某一个函数 ，并且输出 执行正确的函数 和 执行错 										 误的函数
}*/
