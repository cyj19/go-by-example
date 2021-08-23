package main

import (
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
)

var addr = "127.0.0.1:8080"

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post_form", postByFormHandler)
	http.HandleFunc("/post_json", postByJsonHandler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("服务异常退出：%v \n", err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	path := r.URL.Path
	data := r.URL.Query()
	result := path + " form " + data.Get("name")
	w.Write([]byte(result))
}

// 处理表单数据
func postByFormHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	path := r.URL.Path
	// 请求类型是application/x-www-urlencoded时解析from数据
	r.ParseForm()
	result := path + " name=" + r.PostForm.Get("name") + " age=" + r.PostForm.Get("age")
	w.Write([]byte(result))
}

func postByJsonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	path := r.URL.Path
	param, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read req fail err: %v \n", err)
		return
	}
	result := fmt.Sprintf("%s %s", path, string(param))
	w.Write([]byte(result))
}
