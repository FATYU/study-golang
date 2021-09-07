package main

import "fmt"

func main() {
	fmt.Println(comma("123456789"))
}
func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}
