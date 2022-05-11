package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 通过 go 连接 Redis 实现 Hash数据的写入和读取

	// 1、 通过 go 连接 Redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println(" redis Dial err", err)
		return // 拨号连接不成功就直接退出
	}
	defer conn.Close() //及时设计延时关闭机制

	//————————————————————  一次 一次 的读写数据太慢了 ——————————————————————
	// 2、 go 操作 Redis 写入 Hash 数据
	_, err = conn.Do("HSet", "User01", "name", "Jin")
	if err != nil {
		fmt.Println("HGet err", err)
		return
	}

	// 3、 go 操作 Redis 读取 Hash 数据

	r2, err := redis.String(conn.Do("HGet", "User01", "name"))
	if err != nil {
		fmt.Println("HGet err", err)
		return
	}
	fmt.Println("HGet succeed ...", r2)

	//——————————————————————————————  使用同时读写  ——————————————————————————————————————

	//	使用 Hmset 一次写入多个 Hash 数据
	_, err = conn.Do("Hmset", "User02", "name", "Mary", "age", 18, "address", "HuanJiang")
	if err != nil {
		fmt.Println("Hmset err", err)
		return
	}

	// 一次性 “读取” 多个数据 必须使用 Strings ，否则会报错
	r3, err := redis.Strings(conn.Do("Hmget", "User02", "name", "age", "address"))
	if err != nil {
		fmt.Println("Hmget err", err)
		return
	}
	fmt.Println("Hmget succeed ...", r3) // Hmget succeed ... [Mary 18 HuanJiang]  输出多个数据则以集合的形式输出

}
