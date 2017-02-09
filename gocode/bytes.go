package main

import "fmt"

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
