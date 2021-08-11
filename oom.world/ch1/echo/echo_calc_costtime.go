package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	var s string
	var sep = " "
	var start = nowNanosecond()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	var end = nowNanosecond()
	fmt.Println(s)
	fmt.Println("foreach cost ", end-start, " nano seconds")
	start = nowNanosecond()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end = nowNanosecond()

	fmt.Println("strings.join cost ", end-start, " nano seconds")
}

func nowMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
func nowNanosecond() int64 {
	return time.Now().UnixNano()
}
