package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	messages = make(chan string) //所有客户端监听的信息channel

	entering = make(chan client)

	leaving = make(chan client)
)

func broadcast() {

	//定义全部客户端连接集合
	//是否存在某个客户端
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli) //关闭channel
		}
	}

}

func handleClient(conn net.Conn) {
	ch := make(chan string)
	go write2Client(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "you are [ + " + who + "]"
	messages <- who + " connected !" //发送给所有客户端
	entering <- ch
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() { //迭代打印客户端输入
		messages <- who + " : " + scanner.Text()
	}
	leaving <- ch // 客户端断开连接,发送退出消息
	messages <- who + " left"
	err := conn.Close()
	if err != nil {
		return
	}
}

func write2Client(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//功能介绍
//1.实现广播
//2.接收客户端连接，并向客户端发送欢迎信息
//3.转发客户端的消息给所有客户端
//4.关闭链接，删除客户端
func main() {
	//1.创建服务端连接
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	go broadcast() //处理广播信息
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Fatal(err) //出现错误，继续监听
			continue
		}
		go handleClient(accept)
	}
}
