package main

import "fmt"

func main() {
	// var array [2]int

var test := {}

	//var array_2 [1]int
	brray := [...]int{3: 4}
	var a = [...]int{10: 2, 3: 3}

	//array = array_2

	fmt.Println(brray)
	fmt.Println(a)

	//数组指针 指针数组
	a1 := [...]int{99: 1}

	var p *[100]int = &a1

	fmt.Println(p)

	x, y := 1, 2
	a2 := [...]*int{&x, &y}
	fmt.Println(a2)
	s1, s2 := "s1", "s2"
	string_pointer_array := [...]*string{&s1, &s2}
	fmt.Println(string_pointer_array)
	//⬇️error cannot use &s1 (type *string) as type *int in array
	// string_fake_int_pointer_array := [...]*int{&s1, &s2}

}
