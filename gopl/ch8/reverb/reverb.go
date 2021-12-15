package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			continue
		}
		go handleConn(conn)
	}

}

func echo(c net.Conn, shout string, duration time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(duration)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(duration)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(duration)

}

func handleConn(c net.Conn) {

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		echo(c, scanner.Text(), 1*time.Second)    //接收到客户端信息， 然后回响
		go echo(c, scanner.Text(), 1*time.Second) //开启 goroutine 接收到客户端信息， 然后回响
	}
	c.Close()
}
