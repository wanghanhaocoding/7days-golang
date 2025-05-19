package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{}

// 实现了 net/http 包中的 ServeHTTP 方法，从而满足了 http.Handler 接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
func main() {
	engine := new(Engine)
	//任何实现了 ServeHTTP 方法的类型都可以作为 http.Handler 使用。
	//Engine 结构体实现了 ServeHTTP 方法
	log.Fatal(http.ListenAndServe(":9999", engine))
}
