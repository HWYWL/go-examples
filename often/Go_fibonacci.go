package main

import (
	"fmt"
	"time"
)

const LIM = 51

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		//普通版 耗时3m20s
		//result = fibonacci(i)
		//缓存版 0s
		result = fibonaccc(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("消耗时间为: %s\n", delta)
}

//递归运算
func fibonacci(n int) (res uint64) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

//缓存版递归运算
func fibonaccc(n int) (res uint64) {
	// 检查数组内存中是否存在
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonaccc(n-1) + fibonaccc(n-2)
	}
	fibs[n] = res
	return
}
