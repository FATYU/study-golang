package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["zhangyu"] = 33
	ages["zy"] = 1988

	classes := map[string]string{
		"math": "zhangyu",
	}

	fmt.Print("classes index ", classes["math"])
}
