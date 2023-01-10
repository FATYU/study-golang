package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix() // 一个无效的码点
	fmt.Println(now)
	str := string(now)
	fmt.Println(string(2251799813685248))
	fmt.Println(str)
	fmt.Printf("%X\n", []byte(str))

}
