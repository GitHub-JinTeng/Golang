package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//数组
func TestSuzu(t *testing.T) {
	// 1，定义一个数组
	var hens [6]float64 // 声明 6个float类型的数
	// 2， 给数组的每个元素赋值，元素的下标是从0开始的
	hens[0] = 5.0
	hens[1] = 5.0
	hens[2] = 5.0
	hens[3] = 5.0
	hens[4] = 5.0
	hens[5] = 5.0
	//3，遍历数组求和
	sum := 0.0 // 0.0 表示float64
	for i := 0; i < len(hens); i++ {
		sum += hens[i] // 把每个数加到sum中
	}
	fmt.Printf("len(hens)的数据类型%T\n", len(hens)) //len(hens)的数据类型 int
	//4，求平均值
	pj := fmt.Sprintf("%.2f", sum/float64(len(hens))) //"%.2f"是直接将值格式化成保留两位小数
	// 因为[6]float的个数为6,是int类型，因此需要转成float64；
	fmt.Printf("总和为=%v\n平均值=%v", sum, pj)

}

// 数组的几种定义方式
func Test6(t *testing.T) {

	var a [3]int = [3]int{1, 2, 3}
	fmt.Printf("a数组的元素有:%d\n", a) //a数组的元素有:[1 2 3]

	var b = [3]int{4, 5, 6}
	fmt.Printf("b数组的元素有：%d\n", b) //b数组的元素有：[4 5 6]

	var c = [...]int{6, 7, 8, 9}
	fmt.Printf("c数组的元素有：%d\n", c) //c数组的元素有：[6 7 8 9]

	// 可以指定元素值对应的下标
	var names = [3]string{1: "tom", 0: "jack", 2: "marry"}
	fmt.Printf("names数组的元素有：%v\n", names) //names数组的元素有：[jack tom marry]

	// 使用不定长度赋值
	namesArr := [...]string{1: "yoy", 0: "gol", 2: "opwe"}
	fmt.Printf("names数组的元素有：%v", namesArr) //names数组的元素有：[gol yoy opwe]

}

// for - range 遍历数组
func Test7(t *testing.T) {

	name := [...]string{"宋江", "吴用", "卢俊义"}
	for i, v := range name {
		fmt.Printf("[%v]-------%v\n", i, v)
	}
	// 不获取下标
	for _, v := range name {
		fmt.Println(v)
	}

}

func Test4(t *testing.T) {

	//当我们定义数组后，它们此时都有了内存空间，数组的元素都为默认值 0
	var intArr [3]int
	fmt.Println(intArr) // 运行结果：[0 0 0]

}

func Test5(t *testing.T) {
	// 从终端循环输入5个成绩，保存到float64数组中，并输出
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值:\n", i+1) // 因为i从0开始，所以在i后面+1更美观些
		fmt.Scanln(&score[i])
	}
	//打印数组变量
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}
}

//for - range 应用案例：打印输出26个字母
func Test8(t *testing.T) {
	var myChars [26]byte
	for i := 0; i < 26; i++ { //这里的 i 是将 ‘A’字母开始往后循环遍历多少次，26个字母，显然需要遍历26次
		myChars[i] = 'A' + byte(i) // ‘A’ 是byte类型 ，i 是 int类型，所以需要转换成byte类型
	}
	for i := 0; i < 1; i++ { //这里的 i 是将上面生成的26个字母循环输出多少次，所以为 i<1
		fmt.Printf("%c", myChars)
	}
	//运行结果：
	//	[A B C D E F G H I J K L M N O P Q R S T U V W X Y Z]

}

// 求出最大值，并且求出最大值的下标
func Test9(t *testing.T) {
	var cb [9]int = [9]int{1, 3, 4, 6, 7, 8, 0, 5, 32}
	var max = cb[0] // 假设第一个元素就是最大值
	var xia = 0     // 声明一个变量来保存下标
	for i := 0; i < len(cb); i++ {
		if cb[i] > max {
			max = cb[i] // 将大的值赋给 max
			xia = i     // 将最大值的下标赋给 xia
		}
	}
	fmt.Printf("cb的最大值为:%v\n最大值的索引为：%v\n", max, xia)
	//运行结果：
	//cb的最大值为:32
	//		最大值的索引为：8
}

// 求数组的平均值
func Test10(t *testing.T) {

	var a = [5]int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	sj := sum / int(len(a))
	fmt.Println("平均值为：", sj)

}

//随机生成数
func Test11(t *testing.T) {

	var intArr3 [6]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100) //0 <=n<100
	}
}


