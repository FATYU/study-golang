package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("zhangyu"))
	//Output: string

	var w io.Writer = os.Stdout

	fmt.Println(reflect.TypeOf(w))
	fmt.Printf("%T\n", w)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Println(reflect.ValueOf(3))
	fmt.Printf("%V", 3)

}
