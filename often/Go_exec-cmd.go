package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

//执行系统命令

func Exec1() {
	//执行系统命令
	//第一个参数是命令名称
	//后面参数可以有多个，命令参数
	cmd := exec.Command("cmd")
	//获取输出对象，可以从改对象中获取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	//延迟关闭资源，保证关闭输出流
	defer stdout.Close()

	//运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	//读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(opBytes))
}

func main() {
	Exec1()
}
