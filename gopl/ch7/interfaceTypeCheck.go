package main

import "fmt"

// RPCError
//定义 RPCError类型
type RPCError struct {
	Code    int64
	Message string
}

// 实现了error的Error方法

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func NewRPCError(code int64, message string) error {
	return &RPCError{code, message} //返回类型和对象进行检查
}

func asError(e error) error {
	return e
}

func main() {
	var rpcError error = NewRPCError(500, "Internel Error") //类型检查，运行时，对创建的对象和声明类型进行比较
	e := asError(rpcError)                                  //入参和对象类型进行检查
	fmt.Println(e)

}
