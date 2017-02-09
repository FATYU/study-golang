package main

import "fmt"

func main() {
start:
	for i := 1; i < 100; i++ {
		for {
			fmt.Print(i)
			fmt.Println("...........dead loop......")

			continue start
		}
	}
}
