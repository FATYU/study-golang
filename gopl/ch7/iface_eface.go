package main

import (
	"bytes"
	"fmt"
)

func main() {
	var eface interface{} //interface{}值持有一个boolean，float，string，map，pointer，或者任意其它的类型
	eface = true
	fmt.Println(eface)
	eface = 12.34
	fmt.Println(eface)
	eface = "hello"
	fmt.Println(eface)
	eface = map[string]int{"one": 1}
	fmt.Println(eface)
	eface = new(bytes.Buffer)
	fmt.Println(eface)

}

type iface struct {
	//一个具体的类型可能实现了很多不相关的接口
}

func newIface() *iface {
	return &iface{}
}
