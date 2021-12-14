package main

import (
	"fmt"
	"io"
	"os"
)

// 通过类型断言，对对象进行正确的赋值，如果正确，赋值成功，如果类型断言错误，赋值为nil
//断言类型如 x.(T)
// x 表示一个接口的类型
// T 表示一个类型
//

func main() {
	var w io.Writer
	w = os.Stdout
	o := w.(*os.File) //赋值成功
	fmt.Println(o)

	//c := w.(*bytes.Buffer) //Panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	//fmt.Println(c)

	// ---------------------------

	i := interface{}(10)
	fmt.Println("the original data is :", i)

	//断言是整数
	assertIisInt := i.(int)
	fmt.Println("type assert int :", assertIisInt)

	assertIisStr := i.(string)
	fmt.Println("type assert string :", assertIisStr) //panic  interface {} is int, not string

}
