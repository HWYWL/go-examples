package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 0; i <= 35; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("消耗时间为: %s\n", delta)
}

//递归运算
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
