package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x)

	fmt.Println(real(x + y))
	fmt.Println(real(x - y)) //-5 (real1+real2) - (image1+image2)
	fmt.Println(real(x * y)) //
	fmt.Println(real(x / y))

	z := 1 + 2i
	fmt.Println(z)
}
