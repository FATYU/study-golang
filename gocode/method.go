package main

import (
	"fmt"
)

type A struct {
}

type B struct {
}

func (b B) print() {
	fmt.Println("this is B")
}

func (a A) print() {
	fmt.Println("this is A")
}

func main() {
	a := A{} //方法绑定到对象上
	a.print()

	b := B{}
	b.print()
}
