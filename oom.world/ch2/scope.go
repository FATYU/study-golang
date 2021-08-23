package main

import (
	"fmt"
)

func f() {
	fmt.Println("func f", p)
}

var p = " file scope p"

func main() {
	f()
	f := "f"
	fmt.Println(f)
	fmt.Println(p)
	// f() //error
	// fmt.Println(h) //error not define

}
