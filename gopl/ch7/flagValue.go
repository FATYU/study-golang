package main

import (
	"flag"
	"fmt"
	"time"
)

// *time.Duration
var period = flag.Duration("p", time.Second, " sleep period ")

func main() {

	fmt.Println(">>>>>>>>>>", *period)
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println("===================")

}
