package main

import (
	"fmt"
	"math/rand"
)

// 随机密码生成

//盐
var letters = []rune("hfislfhasfhfdjskfhisufhifusifusgy")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func main() {
	//idusjkahfhshfffhifsfsh
	fmt.Println(randSeq(22))
}
