package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	for _, connUrl := range os.Args[1:] {
		go conn(connUrl)
	}
	for {
		time.Sleep(1000 * time.Microsecond)
	}

}

func conn(url string) {
	dial, err := net.Dial("tcp", url)

	if err != nil {
		log.Fatal("err:", err)
		return
	}
	fmt.Println("------------------------------- server : ", url, " connected -----------------------------")
	defer dial.Close()
	mustCopy(os.Stdout, dial)
}

func mustCopy(stdout *os.File, dial net.Conn) {
	if _, err := io.Copy(stdout, dial); err != nil {
		log.Fatal("err:", err)
	}

}
