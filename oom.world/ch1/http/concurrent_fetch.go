package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) //创建 channel 接收字符串

	for _, url := range os.Args[1:] {
		go fetch(url, ch) //每一个参数都会开启goroutine,异步执行http.Get请求
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds()) // 使用since方法进行耗时计算

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("[ERROR] while reading %s: %v", url, err)
		return
	}
	secondCost := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs 7%d %s", secondCost, nbytes, url)

}
