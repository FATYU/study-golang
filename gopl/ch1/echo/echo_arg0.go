package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[0]
	fmt.Println(s)
	s1 := os.Args[1]
	fmt.Println(s1)
	s2 := os.Args[2]
	fmt.Println(s2)
}
