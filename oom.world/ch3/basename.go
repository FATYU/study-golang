package main

import "fmt"

func main() {
	fmt.Println(basename("/a/b/c.go"))
	fmt.Println(basename("/a/b/c测试.go"))
	fmt.Println(basename(",,.llop..go"))
	fmt.Println(basename("c...d.go"))

}

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
