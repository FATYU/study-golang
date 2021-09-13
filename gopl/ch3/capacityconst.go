package main

import (
	"fmt"
	"math"
)

const (
	KB = 1024
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Printf("T%\n", x)
	fmt.Printf("T%\n", y)
	fmt.Printf("T%\n", z)
}
