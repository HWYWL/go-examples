package main

import (
	"net/http"
	"fmt"
	"log"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello golang http!")
}

func main() {
	http.HandleFunc("/", index)

	//设置静态目录
	fsh := http.FileServer(http.Dir("E:/gitCode/go-examples/often"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))

	//设置服务器端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
