package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 使用 ioutil 读取文件
func main() {

	counters := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup_file_ioutil : %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counters[line]++
		}
	}
	for text, countNum := range counters {
		if countNum > 1 {
			fmt.Println(text, "count:", countNum)
		}
	}

}
