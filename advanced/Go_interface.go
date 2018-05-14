package main

import "fmt"

//接口
type Map interface {
	put(key, value string)
	get(key string) string
}

type HashMap struct {

}

var pool = make(map[string]string)

func (self *HashMap) put(key, value string) {
	pool[key] = value
}

func (self * HashMap) get(key string) string {
	return pool[key]
}

func main() {
	var m Map
	hm := HashMap{}
	hm.put("yi", "Hello YI")

	m = &hm

	fmt.Println(m.get("yi"))
}
