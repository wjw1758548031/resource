package main

import (
	"fmt"
	"log"
	"net/http"
)

// 处理根路径请求
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP server!")
}

// 处理 `/hello` 路径请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// 设置路由处理函数
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	// 启动 HTTP 服务器
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
