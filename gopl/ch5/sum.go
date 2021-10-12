package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 23, 32, 1, 321, 3, 21))

}

func sum(vals ...int) int {
	var total = 0
	for i, val := range vals {
		fmt.Println("current index is --->", i)
		total += val
	}
	return total
}
