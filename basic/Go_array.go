package main

import "fmt"

/*
	数组的长度是其类型的一部分，因此数组不能改变大小。
	这看起来是一个制约，但是请不要担心，Go提供了更加方便的方式使用数组
*/
func main() {

	//普通数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//进化...,高阶数组
	b := [...]int{1, 2, 3}
	fmt.Println("数组长度：", len(b), "========>", b)

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
}
