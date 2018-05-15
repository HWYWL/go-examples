package main

import (
	"net/http"
	"fmt"
	"log"
)

//w表示response对象，返回给客户端的内容都在对象里处理
//r表示客户端请求的对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request)  {
	//往w里写入浏览器输出的内容
	fmt.Fprintf(w, "Hello golang http!")
}

func main() {
	//设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)

	//启动web服务器，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}