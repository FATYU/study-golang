package main

import (
	"errors"
	"fmt"
	"syscall"
)

//package errors
//
//func New(text string) error { return &errorString{text} }
//
//type errorString struct { text string }
//
//func (e *errorString) Error() string { return e.text }

func main() {

	if errors.New("error:oom") == errors.New("error:oom") {
		fmt.Println("error new equals")
	} else {
		errorMsg := fmt.Errorf("ERROR")
		fmt.Println(errorMsg.Error())
		fmt.Println("error not equals")
		//通常使用fmt.Errorf代替errors。New

	}

	//使用syscall来获取系统定义好的错误信息，有限集合
	var err error = syscall.Errno(2)
	fmt.Println(err)

	err = syscall.Errno(3)
	fmt.Println(err)

}
