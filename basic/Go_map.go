package main

import "fmt"

type User struct {
	name string
	agg int
}

var m map[string]User

func main() {
	m = make(map[string]User)
	//key		value
	m["yi"] = User{"懿", 25}
	m["rose"] = User{"肉丝", 22}

	fmt.Println(m["yi"])
	fmt.Println(m["rose"])

	fmt.Println("=====华丽丽的分割线=====")

	var m2 = map[string]User{
		"yi":{"懿", 25},
		"rose":{"肉丝", 22},
	}
	fmt.Println(m2)

	fmt.Println("=====邪恶的分割线=====")

	delete(m2, "yi")
	fmt.Println(m2)

	fmt.Println("=====可爱的分割线=====")

	m2["yi"] = User{"懿(#^.^#)", 25}
	fmt.Println(m2)
}
