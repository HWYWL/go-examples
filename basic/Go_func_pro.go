package main

import (
	"fmt"
	"math"
)

// 函数作为返回值,根据输入的函数，进行不同的运算规则
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/*
	函数的闭包
	Go 函数可以是一个闭包，闭包是一个函数值，它引用了函数体之外的变量。
	这个函数可以对这个引用的变量进行访问和赋值，换句话说这个函数被“绑定”在这个变量上
*/
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i

		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		//开平方根
		return math.Sqrt(x*x + y*y)
	}
	//函数作为参数传递
	fmt.Println(hypot(5, 12))

	//传入函数，作为运算规则
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("=====邪恶的分割线=====")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
