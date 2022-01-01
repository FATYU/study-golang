package main

import (
	"fmt"
	"study-golang/gopl/ch9/bank1"
)

func main() {
	bank1.SafeDeposit(100)
	fmt.Println(bank1.Balance())

	go bank1.SafeDeposit(300)

	go func() {
		bank1.SafeDeposit(200)
		fmt.Println("-------------------->", bank1.Balance())
	}()
	for i := 0; i < 100; i++ {
		fmt.Println(bank1.Balance())
	}

}
