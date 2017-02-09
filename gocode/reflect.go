package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   string
	Name string
	Age  int
}

type Zhangyu struct {
	User
	Gf string
}

func main() {
	zy := Zhangyu{User: User{"id-zy", "zy", 28}, Gf: "gff"}
	u := reflect.TypeOf(zy)
	fmt.Println("reflect typeof is:" + u.FieldByName("Gf"))
}
