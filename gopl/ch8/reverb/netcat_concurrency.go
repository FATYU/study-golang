package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("err:", err)
		return
	}
	defer dial.Close()
	go mustCopy(os.Stdout, dial)
	mustCopy(dial, os.Stdin)
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal("err:", err)
	}

}
