package main

import "fmt"

//常量定义
const PI float64 = 3.1415926

/*
	var语句定义了一个变量的列表，跟函数的参数列表一样，在类型的后面。
	就像在例子中看到的一样，var语句可以定义在包或者函数级别
 */
func main() {
	fmt.Println("Hello World,少年你low逼的一天开始了")
	fmt.Println("圆周率：", PI)
	var a, b int = 1, 2
	fmt.Println(a, b)
	var num = 250
	fmt.Println(num)

	//声明并赋值
	gay冰 := "gaygay的"
	fmt.Println(gay冰)

	var p1 int
	var p2 float32
	var p3 string
	var p4 bool

	fmt.Println(p1, p2, p3, p4)

	fmt.Println("3 + 2 = ", 3 + 2)
}
