package main

import a "fmt"
import (
	os "os"
)

func main() {
	a.Println("hw")
	pid := os.Getpid()
	a.Println(pid)
	var b int = 66
	a.Println(b)
	a.Println("----------------------------")
	bb := string(b)
	a.Println(bb)
}
