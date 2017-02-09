package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{}
	fmt.Printf("%p", &a)
	fmt.Println(a)

	a.Age = 100
	a.Name = "zhangyu"
	fmt.Println(a)
	fmt.Printf("%p\n", &a)
	//-----------------演示值传递,原对象的值不会变化
	renameValueTrans(a)
	fmt.Println("演示值传递,原对象的值不会变化", a)
	//-----------------演示引用传递,原对象会发生变化
	renamePointerTrans(&a)
	fmt.Println("演示引用传递,原对象会发生变化", a)

}

func renameValueTrans(person Person) {
	person.Name = "rename"
}

func renamePointerTrans(person *Person) {
	person.Name = "rename"
}
