package main

import "fmt"

//应用闭包：将函数作为返回值

func main() {
	p2 := Add()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))

	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
