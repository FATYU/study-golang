package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(" Countdown started ! Press return to abort... ")

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) //获取输入数据
		abort <- struct{}{}
	}()

	//

	select {
	case <-time.After(10 * time.Second):
		// do nothing
	case <-abort:
		fmt.Println("Countdown aborted!")
		return
	}
	fire()
}

func fire() {
	fmt.Println("F F Fire.............")
}
