package main

import (
	"bufio"
	"fmt"
	"os"
)

// 找出系统输入中,重复输入的文本数据,并打印文本内容
func main() {

	//使用 map 数据结构进行存储
	counters := make(map[string]int)
	scannerIn := bufio.NewScanner(os.Stdin)
	for scannerIn.Scan() { //遍历标准输入
		//代码等价于
		// text = scannerIn.Text()
		// counters[text] = counters[text]+1
		counters[scannerIn.Text()]++
		if scannerIn.Text() == "bye" {
			fmt.Println("the dulac line  is ⬇️:")
			break
		}
	}

	for text, countNum := range counters {
		if countNum > 1 {
			fmt.Println(text, "count:", countNum)
		}
	}

}
