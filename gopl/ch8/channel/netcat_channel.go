package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	dial, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("err:", err)
		return
	}
	//doneLock := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, dial)
		log.Println("done")
		fmt.Println("ddddone")
		//doneLock <- struct{}{} //写进一个空结构
	}()

	//mustCopy(os.Stdout, dial)
	mustCopy(dial, os.Stdin)

	dial.Close()
	//<-doneLock //等待后台goroutine结束

	// 关注代码中的donelock，会让主goroutine等待后台goroutine执行完毕才停止
	/**doneLock := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, dial)
		log.Println("done")
		fmt.Println("ddddone")
		//doneLock <- struct{}{} //写进一个空结构
	}()

	//mustCopy(os.Stdout, dial)
	mustCopy(dial, os.Stdin)

	dial.Close()
	<-doneLock //等待后台goroutine结束
	*/

}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal("err:", err)
	}

}
