package main

import (
	"fmt"
	"net/http"
)

var count int

func main() {

	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/counter", counter)
	http.ListenAndServe("localhost:8888", nil)
}

func helloHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi,you request the service : hello\n")
	count++

}

func counter(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello service called  %d times\n", count)
}
