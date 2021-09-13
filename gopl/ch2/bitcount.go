package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	bitCount(103213)
}

func bitCount(x uint64) int {
	// fmt.Println(x >> 0)
	// fmt.Println(byte(x >> 0))

	fmt.Println(byte(x >> (0 * 8)))
	fmt.Println(byte(x >> (1 * 8)))
	fmt.Println(byte(x >> (2 * 8)))
	fmt.Println(byte(x >> (3 * 8)))
	fmt.Println(byte(x >> (4 * 8)))
	fmt.Println(byte(x >> (5 * 8)))
	fmt.Println(byte(x >> (6 * 8)))
	fmt.Println(byte(x >> (7 * 8)))

	// fmt.Println(pc[byte(x>>(0*8))])

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
