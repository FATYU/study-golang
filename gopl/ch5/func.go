package main

import "fmt"

// 函数在golang中是一等公民
//可以有类型、赋值给变量、作为值传递给别的函数、返回一个函数
func main() {
	test(test(2))

	var a = test(1)
	fmt.Println(a)
}

func test(a int) int {
	return 1
}
