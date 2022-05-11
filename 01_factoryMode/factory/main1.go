package main

import (
	"fmt"
	"learn/01_factoryMode/jack"
)

func main() {
	//p := model.NewPerson("jack")
	//p.SetAge(24) // 传参的时候是要使用 SetXxx()方法传参
	//p.SetSal(10000)
	//
	//fmt.Println(p)  // 因为 p是一个指针，所以不使用取值符号（*）时会出现地址符号 &{jack 24 10000}
	//fmt.Println(*p) // 使用取值符号（*）时就直接访问到结构体的值 {jack 24 10000}
	//
	//fmt.Println(p.Name, "age=", p.GetAge())                     // 调用的时候要用 GetXxx()方法调用
	//fmt.Println(p.Name, "age=", p.GetAge(), "sal=", p.Getsal()) //jack age= 24 sal= 10000

	p := jack.NewPerson("", 10, 30000)
	p.SetSal(50000)
	p.SetAge(150)
	p.SetName("")
	fmt.Println(*p)
}
