package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap(y):%d \t%v\n", i, cap(y), y)
		x = y
	}
}

//自动进行slice扩容,对slice容量进行倍数扩容
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	fmt.Println("zlen:", zlen, "cap(x):", cap(x))
	if zlen <= cap(x) { //当x的容量足够时,进行共享底层数组
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) //拷贝x到z
	}
	z[len(x)] = y
	return z
}
