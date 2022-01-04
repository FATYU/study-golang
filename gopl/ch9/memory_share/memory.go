package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	var x, y int

	go func() {
		wg.Add(1)
		x = 1
		fmt.Print("y:", y, " ")
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		y = 1
		fmt.Print("x:", x, " ")
		wg.Done()
	}()

	wg.Wait()
}
