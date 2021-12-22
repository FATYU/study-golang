package main

import "fmt"

func main() {
	channel := make(chan int, 1) //当channel中不存在缓冲

	// 这种写法，select是随机的分配资源进行case执行， 如果channel没有缓冲，会导致死锁问题
	for i := 0; i < 10; i++ {
		select {
		case channel <- i:
		case x := <-channel:
			fmt.Println(x)
		}

	}

}
