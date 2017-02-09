package main

import (
	"fmt"
)

type human struct {
	Sex  int
	Name string
}
type man struct {
	human
	Name string
	oi   int
}
type woman struct {
	human
	oi int
}

func main() {

	//m := man{Sex: 1, Name: "g", oi: 1}
	//w := woman{Sex: 0, Name: "w", oi: 0}

	m := man{human: human{Sex: 1, Name: "g"}, Name: "zy", oi: 1}
	w := woman{human: human{Sex: 0, Name: "w"}, oi: 0}
	m.Name = "re-zy"
	m.human.Name = "re-human"
	fmt.Println(m)
	fmt.Println(w)

}
