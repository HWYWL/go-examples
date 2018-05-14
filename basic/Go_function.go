package main

import "fmt"

func sum1(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func hello(name string) string  {
	return "你好呀low逼的少年 " + name
}

/*
	当两个或多个连续的函数命名参数类型相同的，除了最后一个外，其他的可以省略
 */
func main() {
	fmt.Println(sum1(10, 20))
	fmt.Println(sum2(10, 20))
	fmt.Println(hello("gay冰"))
}
