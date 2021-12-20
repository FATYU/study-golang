package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//获取现在三大crypto currency cex中，btc-usdt报价最快的数据作为第一结果

func main() {
	dataChannel := make(chan string, 3)

	go func() {
		//curl http://google.com
		dataChannel <- fetch("http://www.baidu.com")
	}()

	go func() {
		dataChannel <- fetch("http://www.sohu.com")
	}()

	go func() {
		dataChannel <- fetch("http://www.sina.com")
	}()

	for data := range dataChannel {
		fmt.Println(data)
		fmt.Println("----------------------------------------------------------end")
		break
	}
}

func fetch(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return " http get error "
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "res err"
	}
	return string(data)
}
