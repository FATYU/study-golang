package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

type ReaderWriter interface {
	io.Reader
	io.Writer
}

func main() {
	var writer io.Writer
	writer = os.Stdout //面向接口编程
	writer.Write([]byte("hello"))
	writer = new(bytes.Buffer)
	writer = time.Second

	var rwc io.ReadWriteCloser
	writer = rwc // 面线接口编程， 引用类型包含了声明类型，是可以赋值的。
	// rwc = writer //缺少了read方法和close方法，编译不通过

}


