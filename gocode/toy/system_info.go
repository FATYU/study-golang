package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
	out1, err1 := exec.Command("ps").Output()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("the ps is :", out1)
}
