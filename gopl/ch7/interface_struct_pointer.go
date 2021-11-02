package main

import "fmt"

// 指针和结构体两种实现方式不能同时存在

type Duck interface {
	Walk()
	Quack()
}
type Cat struct{}

//// Walk //-------------------------指针实现类型
//func (c *Cat) Walk() {
//	fmt.Println("WWWWalk")
//}
//
//// Quack //-------------------------指针实现类型
//func (c *Cat) Quack() {
//	fmt.Println("Quack")
//}

//Walk  //-------------------------结构体实现类型
func (c Cat) Walk() {
	fmt.Println("WWWWalk")
}

//Quack  //-------------------------结构体实现类型
func (c Cat) Quack() {
	fmt.Println("Quack")
}

//-----------------------------------------------------
//                      结构体实现接口   ｜ 指针实现接口
//结构体初始化变量           通过         ｜ 不通过
//指针初始化变量             通过         ｜ 通过

func main() {
	var pointerDuck Duck = &Cat{} //使用指针初始化变量，指针类型隐式的获取到真实的对象
	var structDuck Duck = Cat{}   //使用结构体初始化变量
	pointerDuck.Walk()
	structDuck.Walk()

}
