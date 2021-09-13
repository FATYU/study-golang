package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	Sha256x := sha256.Sum256([]byte("x"))
	Sha256X := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n", Sha256x)
	fmt.Printf("%x\n", Sha256X)            // x% 以十六进制的格式打印数组或slice全部的元素
	fmt.Printf("%t\n", Sha256X == Sha256x) //t% 打印 boolean 类型
	fmt.Printf("%T\n", Sha256X)            //T% 打印数据类型
}
