package main

import "fmt"

func main() {
	s := "oom.world"
	t := s
	s = "hi," + s
	fmt.Printf("string s data is [%s]\n", s)
	fmt.Printf("substring:5,data is [%s]\n", s[:5])
	fmt.Printf("t data is [%s]\n", t)
	//对字符串中的字符进行修改
	//s[1] = 'p' //compile error: cannot assign to s[0]

	//原生字符串面值
	zhangyu := ` 年龄 :33岁
	-----------
	性别 : 男
	`
	fmt.Println(zhangyu)
}
