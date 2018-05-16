package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

/*
	类型通过实现那些方法来实现接口，没有显示声明的必要，所以也就没有关键字“implements”
	隐式接口解耦了实现接口的包和定义接口的包，互不依赖
	因此也就无需在每个实现上新增的接口名称，这样同时也鼓励了明确接口的定义
*/
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	//os.Stdout 实现了writer
	w = os.Stdout

	fmt.Println(w, "Hello writer\n")
}
