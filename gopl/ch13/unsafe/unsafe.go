package main

import (
	"fmt"
	"unsafe"
)

//People的SizeOf是24字节，AlignOf是8字节

type People struct {
	age      int    // 占用一个机器字 8字节 64位
	name     string // 占用一个机器字 8字节 64位
	birthday string // 占用一个机器字 8字节 64位
}

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	// 使用sizeof打印类型内存占用字节数
	fmt.Println("-------------------SizeOf--------------------")
	fmt.Println(unsafe.Sizeof(int(1)))                                              //8 1word
	fmt.Println(unsafe.Sizeof(float64(1)))                                          //8 1word
	fmt.Println(unsafe.Sizeof("zy"))                                                //16 2words 两个机器字
	fmt.Println(unsafe.Sizeof(People{age: 1, name: "zy", birthday: "1990-01-01"}))  //（ 1+2+2 ）*8 = 40 字节
	fmt.Println(unsafe.Sizeof(&People{age: 1, name: "zy", birthday: "1990-01-01"})) //地址 1word
	fmt.Println("-------------------AlignOf--------------------")
	fmt.Println(unsafe.Alignof(int(1)))
	fmt.Println(unsafe.Alignof(float64(1)))
	fmt.Println(unsafe.Alignof("zy"))
	fmt.Println(unsafe.Alignof(People{age: 1, name: "zy", birthday: "1990-01-01"}))
	fmt.Println(unsafe.Alignof(&People{age: 1, name: "zy", birthday: "1990-01-01"}))
	fmt.Println("-------------------OffsetOf--------------------")
	fmt.Println(unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a)) //0 , 内存补位 补到2
	fmt.Println(unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b)) //2 ，内存补位 补到8（如果是32位机器，大小是4位，不用进行内存补位）
	fmt.Println(unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c)) //8
}
