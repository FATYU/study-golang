package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3
	fmt.Println("before set : ", x)
	//通过 reflect.ValueOf() 获取的值，是不能取地址的。

	fmt.Println("reflect.ValueOf(3) can get address: ", reflect.ValueOf(x).CanAddr()) //查看是否可以取地址

	fmt.Println("reflect.ValueOf(3)  can set new value :", reflect.ValueOf(x).CanSet()) //查看是否可以重新设置值

	// 获取x的地址
	// 获取地址的值 Elem returns the value that the interface v contains
	// or that the pointer v points to.
	reflect.ValueOf(&x).Elem().SetInt(4)

	// runtime err :  call of reflect.Value.SetFloat on int Value 不能将float64设置到int类型上
	//reflect.ValueOf(&x).Elem().SetFloat(65535.00)

	fmt.Println("after set : ", x)
}
