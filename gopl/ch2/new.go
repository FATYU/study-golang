package main

import "fmt"

func main() {
	age := new(int8) // 使用内置new函数进行对象、基本类型创建
	fmt.Println(age)
	fmt.Println(*age)
	address := &age
	fmt.Println(&address)
	fmt.Println(*(*address))
	addressOfAddress := &address
	fmt.Println(&addressOfAddress)
	fmt.Println(*(*addressOfAddress))

}
