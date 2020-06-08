package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认不解析
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url.long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello zy !") // 写入w的是输出到客户端的信息
}

func main() {
	http.HandleFunc("/", sayHelloName) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的接口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
