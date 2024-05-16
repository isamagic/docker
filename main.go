package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	randID := uuid.NewString()
	fmt.Fprintf(w, "Hello %v", randID)
}

func main() {
	http.HandleFunc("/", sayHello)           // 设置访问的路由
	err := http.ListenAndServe(":1111", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
