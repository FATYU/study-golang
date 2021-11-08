package main

//一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的
import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {

	var writerBuf io.Writer

	if debug {
		writerBuf = new(bytes.Buffer)
	}
	f(writerBuf)

	var buf *bytes.Buffer

	if debug {
		buf = new(bytes.Buffer) // 创建对象，分配内存空间
	}
	f(buf)
	//if debug {
	//	// use buf do something
	//}
}

func f(out io.Writer) {
	//传入的对象是空，接口的nil和 原生nil是不相同的。
	// out  ==  {io.Writer} nil
	// nil == {}nil

	//可以使用 reflect.ValueOf(out).IsNil() 进行判断是否为 nil

	if out != nil {
		out.Write([]byte("done ! \n"))
		fmt.Println("--------------------------")
	}

}
