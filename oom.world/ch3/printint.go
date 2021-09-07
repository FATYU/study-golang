package main

import (
	"bytes"
	"fmt"
)

func intToString(ints []int) string {
	//使用 bytes.Buffer 进行字节的存储,可以动态拼接
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range ints {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	return buf.String()
}

func main() {
	fmt.Println(intToString([]int{1, 2, 3, 3, 45, 33}))
}
