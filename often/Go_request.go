package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net/url"
)

const BaseUrl = "http://httpbin.org"

//get请求
func httpGet() {
	resp, err := http.Get(BaseUrl + "/get")
	if err != nil {
		fmt.Println("handle error")
	}

	//延迟释放网络资源
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	fmt.Println(string(body))
}

//post方式请求
func httpPost()  {
	resp, err := http.Post(BaseUrl + "/post?hello=world",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=yi"))

	if err != nil {
		fmt.Println(err)
	}

	//延迟释放网络资源
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	fmt.Println(string(body))
}

//post表单提交
func httpPostForm()  {
	resp, err := http.PostForm(BaseUrl + "/post",
		url.Values{"key":{"Value"}, "id":{"123"}})
	if err != nil {
		fmt.Printf("err was %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("handle error %v", err)
	}

	fmt.Println(string(body))
}

func httpDo()  {
	client := &http.Client{}

	req, err := http.NewRequest("Post", BaseUrl + "/post", strings.NewReader("name=xxxxxxx"))
	if err != nil {
		fmt.Printf("handle error %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("handle error %v", err)
	}

	fmt.Println(string(body))
}

func main() {
	httpGet()
	fmt.Println("==========可爱的分割线==========")

	httpPost()
	fmt.Println("==========可爱的分割线==========")

	httpPostForm()
	fmt.Println("==========可爱的分割线==========")

	httpDo()
}
