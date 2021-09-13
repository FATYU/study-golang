package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", " "

	//使用range方法进行遍历分片
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// }

	//使用strings join方法进行字符串拼接操作

	s = strings.Join(os.Args[1:], sep)
	fmt.Print(s)
}
