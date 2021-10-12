package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	if s == "" {
		return ""
	}
	var buf bytes.Buffer
	for start := 0; start < len(s); start++ {
		buf.WriteByte(s[start])
		if (len(s)-start-1)%3 == 0 && (start < len(s)-1) {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
