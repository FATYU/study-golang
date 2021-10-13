package main

import (
	"fmt"
	"os"
)

func main() {

	var s = os.Args[1]
	switch s {
	case "z":
		fmt.Println("z")
	case "y":
		fmt.Println("y")
	default:
		panic(fmt.Sprintf("invalid arg! PANIC!!!"))
	}

}
