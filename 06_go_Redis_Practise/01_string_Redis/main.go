package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 需求：通过 go 向 Redis “写入数据 ” 和 “ 读取数据 ”

	//———————————————————————— 一次 一次的 写入数据 ————————————————————————————————
	// 1、 连接到 Redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err", err)
		return // 如果连接不上就直接退出！
	}
	defer conn.Close() //必须及时设置延时关闭机制，否则服务器瞬时报红

	// 2、连接成功后，通过 go 向 Redis 写入 string [ key - value ]数据
	_, err = conn.Do("Set", "name1", "Jack")
	if err != nil {
		fmt.Println("Set err", err)
	}

	// 3、通过 go 向 Redis 读取 string [ key - value ] 数据
	r, err := redis.String(conn.Do("Get", "name1"))
	if err != nil {
		fmt.Println("Get err = ", err)
		return
	}
	/*
		因为返回 r 是 interface{ } 类型
		因为 name 对应的值是 “Jack” 是 string ，因此我们需要转换
		namestring := r.(string)	虽然 r 是 interface{}类型，从理论上可以使用类型断言转换成string类型，
		但在这里使用类型断言转换会报错，因为它本身已经提供了转换方法了，就是上面的 redis.String(conn.Do("Get","name"))
		如果是多个string类型，就写 Strings ，如果是其他类型，就写其他类型 例如： redis.Int 等
	*/
	fmt.Println("conn succeed...", r)

	//———————————————————————— 一次性写入过个数据 ————————————————————————————————
	// 一次性 “写入” 多个数据
	_, err = conn.Do("Mset", "name", "Tom", "age", 20, "address", "HeChi")
	if err != nil {
		fmt.Println("Mset err", err)
		return
	}

	// 一次性 “读取” 多个数据	必须使用 Strings
	r1, err := redis.Strings(conn.Do("Mget", "name", "age", "address"))
	if err != nil {
		fmt.Println("Mget err ", err)
		return
	}
	fmt.Println("Mget succeed....", r1) //	Mget succeed.... [Tom 20 HeChi]

	// 查看 r1 的数据类型
	fmt.Printf("r1=%v\nr1=%T\n", r1, r1) //	r1=[Tom 20 HeChi]   r1=[]string 类型为 string切片

	// 将切片遍历出来
	for i, v := range r1 {
		fmt.Printf("r1[%v]=%s\n", i, v)
	}
	/*
		r1[0]=Tom
		r1[1]=20
		r1[2]=HeChi
	*/

}
