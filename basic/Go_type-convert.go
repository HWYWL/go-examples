package main

import (
	"fmt"
	"math"
)

// 类型转换
func main() {
	var x, y int = 3, 4
	// f = 5.291502622129181
	var f float64 = math.Sqrt(float64(x*y + y*y))
	var z int = int(f)

	// z = 5
	fmt.Println(z)
}
