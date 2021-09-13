package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counters := make(map[string]int)

	files := os.Args[1:]

	if len(files) <= 0 {
		countLines(os.Stdin, counters)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup_file.go run err : %v\n", err)
				continue //继续遍历所有文件
			}
			countLines(f, counters)
			f.Close()
		}
	}

	for text, countNum := range counters {
		if countNum > 1 {
			fmt.Println(text, "count:", countNum)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f) // CHROM
	for input.Scan() {
		counts[input.Text()]++
		//出现重复数据时,打印文件内容
		if counts[input.Text()] > 1 {
			fmt.Println("file ", f.Name(), " contains exist text,text:[", input.Text(), "] count:[", counts[input.Text()], "]")
		}

		if input.Text() == "bye" {
			break
		}

	}
}
