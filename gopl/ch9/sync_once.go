package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {

	go onceFunc()
	go onceFunc()
	go onceFunc()
	go onceFunc()
	go onceFunc()
	time.Sleep(time.Second)
}

func onceFunc() {
	once.Do(func() {
		fmt.Println("initilizing")
	})
}
