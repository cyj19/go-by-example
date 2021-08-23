package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	post_json()
}

func get() {
	addr := "http://127.0.0.1:8080/get"
	// 设置请求参数
	data := url.Values{}
	data.Set("name", "zs")
	u, err := url.ParseRequestURI(addr)
	if err != nil {
		log.Fatalf("parse request uri err: %v \n", err)
	}
	u.RawQuery = data.Encode()
	// 发送请求
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get fail err: %v \n", err)
		return
	}
	// 处理完关闭回复体
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp fail err: %v \n", err)
	}
	fmt.Println(string(result))
}

func post_form() {
	addr := "http://127.0.0.1:8080/post_form"
	// 表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "name=zs&age=20"
	resp, err := http.Post(addr, contentType, strings.NewReader(data))
	if err != nil {
		log.Fatalf("post fail err: %v \n", err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("post resp err: %v \n", err)
		return
	}
	fmt.Println(string(result))
}

func post_json() {
	addr := "http://127.0.0.1:8080/post_json"
	// json格式发送
	contentType := "application/json"
	data := `{"name":"zs", "age":20}`
	resp, err := http.Post(addr, contentType, strings.NewReader(data))
	if err != nil {
		log.Fatalf("post fail err: %v \n", err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("post resp err: %v \n", err)
		return
	}
	fmt.Println(string(result))
}
