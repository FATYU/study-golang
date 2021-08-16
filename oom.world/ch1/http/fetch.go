package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			fmt.Print("input arg is wrong,only support 'http://' ")
			os.Exit(1)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body) //读取resp中body数据到临时变量b
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "response err :%v\n", err)
		// 	os.Exit(1)
		// }
		//fmt.Printf("http response body data is: \n------------------------\n%s", b)

		_, copyErr := io.Copy(os.Stdout, resp.Body) //直接拷贝response-body 到标准输出流
		if copyErr != nil {
			fmt.Fprintf(os.Stderr, "io copy  err :%v\n", err)
			os.Exit(100)
		}
		fmt.Println("http response status code is :", resp.StatusCode)

	}
}
