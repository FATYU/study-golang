package main

import "fmt"

// 大函数进行拆分小函数时， channel可能从原来的双向（in & out）转变成了单向（in ｜ out）
// <-chan : one-way in channel表示只发送的channel
// chan<- : one-way out channel 表示只接收的channel
//

func main() {
	natural := make(chan int)
	squarer := make(chan int)
	go naturals(natural)
	go squarers(squarer, natural)
	printer(squarer)
}

func naturals(out chan<- int) { //往out赋值，只接收

	for i := 0; i < 5; i++ {
		out <- i
	}
	close(out)

}

func squarers(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {

	for data := range in {
		fmt.Println(data)
	}
}
