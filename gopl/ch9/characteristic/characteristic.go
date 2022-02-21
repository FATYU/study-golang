package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 打印机器的处理器数量

	fmt.Println(getProcessorCount())

	// 打印GOMAXPROCS的值

	fmt.Println(getGOMAXPROCS())

	//验证 gomaxprocs

	//执行 `GOMAXPROCS=1 go run characteristic.go` 效果如下
	//11111111000000001111000011001011110000是等量、交替打印

	//执行 `GOMAXPROCS=2 go run characteristic.go` 效果如下
	//0101001010101010101010101001010101010010111111 无序打印

	for {
		go fmt.Print("0")
		fmt.Print("1")
	}
}

func getProcessorCount() int {
	return runtime.NumCPU()
}

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}
