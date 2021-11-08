package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer

	fmt.Printf("origin type is %T\n", w)
	//w.Write([]byte("Writer:::>hello\n")) //空指针异常
	w = os.Stdout
	fmt.Printf("origin type is %T\n", w)
	w.Write([]byte("Stdout:::>hello\n"))
	//fmt.Println(w)
	w = new(bytes.Buffer)
	fmt.Printf("origin type is %T\n", w)
	w.Write([]byte("Buffer:::>hello\n"))
	fmt.Println(w)
	w = nil
	fmt.Printf("origin type is %T\n", w)
	//w.Write([]byte("nil:::>hello\n")) //空指针异常
	//fmt.Println(w)

	var z io.Writer
	if z == w {
		fmt.Println("declare object is equal")
	}
}
