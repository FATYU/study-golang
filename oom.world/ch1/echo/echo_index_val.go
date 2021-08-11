package main

import (
	"fmt"
	"os"
)

func main() {
	for index, val := range os.Args[1:] {
		fmt.Println("index : ", index, ", val : ", val)
	}
	// 运行效果如下:

	// go run echo_index_val.go aaa bbb ccc
	// index :  0 , val :  aaa
	// index :  1 , val :  bbb
	// index :  2 , val :  ccc

}
