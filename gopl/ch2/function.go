package main

import (
	"fmt"
)

func main() {
	var f1 = 212.0
	var f2 = 222.0
	var address = &f1

	fmt.Println("f1 memory address: ", address)
	fmt.Println("print f1 value, by memory pointer :", *address)
	fmt.Printf("boiling point = %g°F or %g°C\n", f1, fc(f1))
	fmt.Printf("boiling point = %g°F or %g°C\n", f2, fc(f2))
}

func fc(f float64) float64 {
	return (f - 32) * 5 / 9
}
