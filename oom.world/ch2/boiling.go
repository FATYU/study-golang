package main

import "fmt"

const boilingF = 212.0 //包级别的常量声明,整个包可以访问

func main() {
	var f = boilingF //函数级别的变量声明
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
