package main

import (
	"fmt"
	"time"
)

func main() {
	// chan index --> |] |] |] 创建了有三个缓存的队列
	//当向队列赋值时，是从末尾向队列进行数据插入。
	//当从队列中获取数据时，是从队列头部开始获取并删除数据。

	channel := make(chan string, 3)
	fmt.Println("channel data cap:", cap(channel)) //返回当前channel容量
	fmt.Println("channel data len:", len(channel)) //返回当前channel 存在数据的量
	go func() {
		for x := 1000; x < 1012; x++ {
			channel <- "num:" + string(x)
			fmt.Println("add+++++++++++++++++", len(channel)) //返回当前channel 存在数据的量
			//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("remove---------------", len(channel)) //返回当前channel 存在数据的量
		fmt.Println(data)
		time.Sleep(1000 * time.Microsecond)
	}

	//⚠️
	//新手容易使用 channel 在单个 goroutine 中，当作队列使用，这个会有隐藏的风险， 如果处理不当
	//会造成 goroutine 阻塞。如果只是简单的使用队列，推荐使用slice
}
