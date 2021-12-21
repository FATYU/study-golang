package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Countdown started! ")

	tick := time.Tick(time.Second) //<-chan Time 返回是 时间类型的channel

	for countdown := 10; countdown > 0; countdown-- {

		fmt.Print(countdown)
		data := <-tick
		fmt.Println(data)
	}
	launch()
}

func launch() {
	fmt.Println("launch.....")
}
