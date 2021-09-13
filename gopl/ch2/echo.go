package main

import (
	"flag"
	"fmt"
	"strings"
)

var blankline = flag.Bool("n", false, "忽略行尾换行符")
var sep = flag.String("sep", "/", "连接符")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))

	if !*blankline {
		fmt.Println() //打印空行
	}

	fmt.Println("=====================")
}
