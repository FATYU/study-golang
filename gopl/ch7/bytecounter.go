package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Writer(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var byteCounter ByteCounter
	p := []byte{'1', '2', '3', '0'}
	p1 := []byte("123456789123456789")
	length, _ := byteCounter.Writer(p)
	length2, _ := byteCounter.Writer(p1)
	fmt.Println("'1', '2', '3', '0' data len is ", length, "\n'123456789123456789' data len is ", length2, "\nthe total data len is ", byteCounter)

}
