package main

import (
	"crypto/md5"
	"io"
	"fmt"
	"os"
)

//计算字符串md5
func CalcStringMd5()  {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker")
	io.WriteString(h, "Add leon's getting hahaha")
	// d9b028d11a9f6377a8f8544812fabebf
	fmt.Printf("%x\n", h.Sum(nil))

	h2 := md5.New()
	h2.Write([]byte("The fog is getting thicker"))
	h2.Write([]byte("Add leon's getting hahaha"))
	//d9b028d11a9f6377a8f8544812fabebf
	fmt.Printf("%x\n", h2.Sum(nil))
}

//计算文件md5
func ClacFileMd5()  {
	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return
	}
	fmt.Printf("%x\n", h.Sum(nil))
}

func main() {
	CalcStringMd5()
	ClacFileMd5()
}
