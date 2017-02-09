package main

import (
	"fmt"
)

func main() {
	var (
		fs = [4]func(){} //匿名函数数组
	)

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer closure i = ", i)
		}() //defer 闭包调用

		fs[i] = func() {
			fmt.Println(".closure..", i)
		}

	}

	for _, f := range fs {
		f() //闭包的调用
	}

}
