package main

import (
	"fmt"
)

/*
不定长参数必须是变量参数中的最后一个
*/
func sysout(args ...string) (results string) {
	results = args[0]
	return results
}

func main() {
	args := sysout("", "", "b", "c")
	fmt.Println("args is :")
	fmt.Println(args)
}
