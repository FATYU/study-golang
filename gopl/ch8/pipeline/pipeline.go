package main

import "fmt"

// counter --> squarer -->printer
//排序
//平方
//打印

// 设置两个channel : natural, squarer
func main() {

	natural := make(chan int)
	squarer := make(chan int)
	//natural func
	go func() {
		for x := 0; x < 5; x++ {
			natural <- x
		}
		//当循环结束，停止goroutine, 如果还有代码从当前channel获取数据， 将获取0值，同时channel返回的ok值是false
		//这时可以根据ok的值进行判断是否可以继续从当前channel获取数据
		close(natural)

	}()

	//squarer
	go func() {
		for {

			//x := <-natural
			x, ok := <-natural
			if !ok {
				break
			}
			squarer <- x * x
		}
		close(squarer)
	}()

	//main goroutine

	// Printer
	//for {
	//	fmt.Println(<-squarer)
	//}

	// for循环也可以直接迭代 channel

	for x := range squarer {
		fmt.Println(x)
	}
}

// ⚠️ 其实你不用关心是否需要关闭channel，因为当channel没有再被引用的时候， 会被golang的GC自动回收
