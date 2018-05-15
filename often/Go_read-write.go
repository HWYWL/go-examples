package main

import (
	"os"
	"fmt"
)

//文件读写

var userFile = "test.txt"

//创建文件并写入字符串
func CreateFile()  {
	fout, err := os.Create(userFile)
	//延迟执行，在最后执行
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	fout.WriteString("Just a Test!\r\n")
	fout.Write([]byte("我是一行可爱的测试文本！\r\n"))
}

//文件读取
func ReadFile()  {
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

// 删除文件
func DeleteFile()  {
	os.Remove(userFile)
}

func main() {
	CreateFile()
	ReadFile()
	DeleteFile()
}
