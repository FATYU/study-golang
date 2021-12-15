package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// 创建http server
	server, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Print(err)
			continue //等待新的客户端链接
		}
		handleConn(conn) // 间隔 1s 打印当前日期
	}

	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Now().Format("2007-01-02 15:04:05"))
	//fmt.Println(time.Now().Format("2006-02-02 15:04:05"))
	//fmt.Println(time.Now().Format("2006-01-03 15:04:05"))
	//fmt.Println(time.Now().Format("2006-01-02 16:04:05"))
	//fmt.Println(time.Now().Format("2006-01-02 15:05:05"))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:06"))

}

func handleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("server conn error：", err)
		}
	}(conn)

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))

		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}

}
