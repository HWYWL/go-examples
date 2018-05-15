package main

import (
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	input := []byte("Hello golang base64 我是字符串")

	//base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	//对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		logrus.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))
	fmt.Println()

	//如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(uDec))
}
