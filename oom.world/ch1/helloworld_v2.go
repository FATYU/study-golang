package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("  你好，请访问 - http://127.0.0.1:8888/ ")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf("你好，world！Now is %s", time.Now().String())
		fmt.Fprintf(rw, " %v\n", s)
		log.Printf("the original data is %v\n", s)
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("http服务启动失败，错误信息：", err)
	}
}
