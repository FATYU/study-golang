package main

import (
	"fmt"
	"unsafe"
)

func main() {
	SetXXB()
	OneLineSetXXB()
}

//⚠️ 不要使用此方法进行指针操作，因为这种方式是不安全的
// 比较微妙的问题，当在执行 unsafe.Pointer(data) 的时候， 如果xx的地址发生了变化。
// data的数据是固定的， 这个时候data指针的地址其实指向了异常数据

func SetXXB() {
	// 计算出&xx对象内存地址，并加上b的偏移量，获取b的地址
	data := uintptr(unsafe.Pointer(&xx)) + unsafe.Offsetof(xx.b)
	//等价于 unsafe.Pointer(&xx.b) 将 pointer 转换成 int16 的基础类型指针
	ptrB := (*int16)(unsafe.Pointer(data))
	*ptrB = 100
	fmt.Println(xx.b)
}

func OneLineSetXXB() {
	// 计算出 &xx 对象内存地址，并计算出b的偏移量
	ptrB := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&xx)) + unsafe.Offsetof(xx.b)))
	*ptrB = 101
	fmt.Println(xx.b)
}

func ErrorUsageuintptr() {
	//	ptr := uintptr(unsafe.Pointer(new(T))) //错误的使用方式
	//	fmt.Println(ptr)
}

var xx struct {
	a bool
	b int16
	c []int
}
