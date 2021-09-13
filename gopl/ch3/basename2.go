package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename2("/a/b/c.go"))
	fmt.Println(basename2("/a/b/c测试.go"))
	fmt.Println(basename2(",,.llop..go"))
	fmt.Println(basename2("c...d.go"))
	fmt.Println(basename2("..go"))

}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dotIndex := strings.LastIndex(s, "."); dotIndex >= 0 {
		s = s[:dotIndex]
	}
	return s
}
