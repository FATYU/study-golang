package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(200 * time.Millisecond)
	const n = 45
	result := fib(n)
	fmt.Printf("\r") //清除打印的字符
	fmt.Println("the fib result of ", n, " is ", result)
}

func fib(n int) int {

	if n > 1 {
		return fib(n-1) + fib(n-2)
	}
	return n

}

func spinner(delay time.Duration) {

	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
