package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 3, 6)
	fmt.Printf("%p\n", slice1) //第一次分配alisce内存值,整个slice容量是6

	slice1 = append(slice1, 2, 3, 4)
	fmt.Printf("%p\n", slice1) //append后内存地址
	slice1 = append(slice1, 21, 31, 41)

	fmt.Printf("%p\n", slice1) //第二次分append后内存地址
	fmt.Println(slice1)

}
