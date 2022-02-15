package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HttpGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{

			"https://www.baidu.com",
			"https://www.sohu.com",
			"https://www.qq.com",
			"https://www.sina.com",
			"https://www.qq.com",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string) (interface{}, error)
}

// 单线程排队执行方式

func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s,%s,%d bytes\n", url, time.Since(start), len(value.([]byte))) //woo, golang is so f****king awesome!
	}
}

func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup //组等待 == Java中deCountdownLatch
	for url := range incomingURLs() {
		n.Add(1) //增加一个等待
		go func(url string) {
			defer n.Done()      //	等待完成
			start := time.Now() //开始时间
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s,%s,%d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()

}
