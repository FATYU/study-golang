package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var fmtString = ""

func main() {

	port := os.Args[1]
	// 创建http server
	server, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("+++++++++++++++++++++++++ server : ", port, " started success +++++++++++++++++++++++++++++")
	for {
		conn, err := server.Accept()
		fmt.Println("------------------------------- client : ", port, " connected -----------------------------")
		if err != nil {
			log.Print(err)
			continue //等待新的客户端链接
		}

		if port == "7000" {
			fmtString = "15:04:05 ------------- > server7 \n"
		}

		if port == "8000" {
			fmtString = "15:04:05 ------------- > server8 \n"
		}

		if port == "9000" {
			fmtString = "15:04:05 ------------- > server9 \n"
		}
		if port == "10000" {
			fmtString = "15:04:05 ------------- > server10 \n"
		}

		go handleConCurrencyConn(conn, fmtString) // 间隔 1s 打印当前日期
	}
}

func handleConCurrencyConn(conn net.Conn, str string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("server conn error：", err)
		}
	}(conn)

	for {
		_, err := io.WriteString(conn, time.Now().Format(str))

		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}

}
