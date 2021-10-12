package main

import "fmt"

func main() {

	f := square()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

//返回一个函数
//函数值不仅仅是一串代码(行为)，还记录了状态
func square() func() int {
	var x int
	return func() int {
		//匿名函数访问了其他函数中的变量
		x++
		return x * x
	}

}
