package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// reflect type
	fmt.Println("reflect typeof {zhangyu}: ", reflect.TypeOf("zhangyu"))
	//Output: string

	var w io.Writer = os.Stdout

	fmt.Println("reflect typeof {os.Stdout}: ", reflect.TypeOf(w))
	fmt.Printf("reflect typeof {os.Stdout} by %%T:  %T\n", w)

	// reflect value
	v := reflect.ValueOf(3)
	fmt.Println("reflect valueof {reflect.ValueOf(3)}: ", v)
	fmt.Println("reflect valueof {reflect.ValueOf(3).Interface()}: ", reflect.ValueOf(3).Interface())
	fmt.Printf("reflect valueof 3 by %V", 3)

}
