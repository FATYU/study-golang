package main

import (
	"fmt"
)

func main() {
	map_int_string := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(map_int_string)
	map_string_int := make(map[string]int)
	fmt.Println(map_string_int)

	for k, v := range map_int_string {
		map_string_int[v] = k
	}
	fmt.Println(map_string_int)

}
