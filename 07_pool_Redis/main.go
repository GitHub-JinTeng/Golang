package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的 pool
var pool *redis.Pool

// 创建一个初始化链接池的函数；当启动程序时，必须先初始化链接池后才能使用链接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 链接的最大空闲数
		MaxActive:   0,   // 表示和数据库的最大链接数，0 表示没有限制（但也必须是在最大链接数的范围之内）
		IdleTimeout: 100, // 表示最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码，链接那个ip地址
			return redis.Dial("tcp", "localhost:6379") // localhost：本地主机
		},
	}
}
func main() {
	// 先从 pool（链接池）取出一个链接
	conn := pool.Get()
	defer conn.Close() // 设置延时关闭机制
	//  写入数据
	_, err := conn.Do("Set", "name", "凯迪拉克CT5")
	if err != nil {
		fmt.Println("Set err", err)
		return // 如果写入数据失败，则直接退出，不再往下执行
	}

	// 取出数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Printf("Get err", err)
		return // 如果读取数据失败，则直接退出，不再往下执行
	}
	fmt.Println("Get succeed...", r) //	Get succeed... 凯迪拉克CT5

	// 注意：当我们关闭了 pool 链接池后，就不能够再次取出链接了
	//pool.Close() //关闭链接池

	conn1 := pool.Get()
	fmt.Println(conn1)
	//  {0xc0000ac020} 这个地址是无效的	注释掉 pool.Close() //关闭链接池后的有效地址为 ：&{0xc000066500 0xc000050410 0}

	// ——————————————————————  证明：——————————————————————-

	// 1、关闭链接池后的 写入数据
	_, err = conn1.Do("Set", "name1", "梅赛德斯")
	if err != nil {
		fmt.Printf("conn1 Set err", err)
		return // 写入失败，直接退出
	}

	// 2、关闭链接池后的 读取数据
	r1, err := redis.String(conn1.Do("Get", "name1"))
	if err != nil {
		fmt.Println("conn1 Get err", err)
		return // 读取失败，直接退出
	}
	fmt.Println("coon1 Get  succeed .....", r1)
	/*
		没有注释掉 pool.Close( ) 语句时的结果为：conn1 Set err%!(EXTRA *errors.errorString=redigo: get on closed pool)

		注释 pool.Close( ) 语句后 的结果为的结果为：coon1 Get  succeed ..... 梅赛德斯
	*/

}
