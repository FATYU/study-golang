package main

import "fmt"

//在golang中 ，go关键词修饰，可以创建多个goroutine。
//不同的goroutine如果需要进行通信，就需要使用channel
//创建channel：
//每个channel有一种数据类型，例如 chan int 表示传递int类型的channel
func main() {

	//The make built-in function allocates and initializes an object of type slice, map, or chan (only).
	s := make(chan string) //如果不指定buffer大小就是unbuffered
	//一个channel有发送和接受两个主要操作，都是通信行为
	fmt.Println(s)

	//语句报错 fatal error: all goroutines are asleep - deadlock!
	//使用匿名goroutine进行发送数据
	//s <- "goodbye.eth" //向 channel 发送数据

	go func() {
		s <- "goodbye.eth"
	}()

	str := <-s // 接收语句，将channel数据赋值给str

	fmt.Println(str) //

	//<-s //接收语句， 丢弃通道数据
}
