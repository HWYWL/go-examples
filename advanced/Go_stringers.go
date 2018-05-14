package main

import "fmt"

type person struct {
	Name string
	Age int
}

//类似于Java中的toString
func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"张三", 18}
	z := Person{"李四", 19}

	fmt.Println(a, z)
}
