package main

import (
	"fmt"
)

type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}
	fmt.Printf("%#v\n", p.X)
	p.X = 2222
	fmt.Printf("%#v\n", p.X)
	// --------------------
	zzz := make(map[address]int)
	zzz[address{"tokensnake.com", 80}]++
	fmt.Printf("%#v\n", zzz)

	// --------------------

	var w Wheel
	w.Circle.Center.X = 100
	w.Circle.Center.Y = 200
	w.Circle.Radius = 10
	w.Spokes = 100

	fmt.Printf("%#v\n", w)
	// -------------------
	var aw AnonymousWheel

	aw.Center.X = 100
	aw.Center.Y = 200
	aw.Radius = 1
	aw.Spokes = 100
	fmt.Printf("%#v\n", aw)

}

// Block .
type Block struct {
	Nonce       string
	Hash        string
	Difficultly int
	Gas         float64
	Gaslimit    float64
}

// Points .
type address struct {
	hostname string
	port     int32
}

// 结构体嵌套
// Circle .
type Circle struct {
	Center Point
	Radius int
}

// Wheel .
type Wheel struct {
	Circle Circle
	Spokes int
}

/// 在Golang中, 如果嵌套结构体,使用匿名引用时,可以直接进行结构体中的属性的引用.示例如下:

// AnonymousWheel .
type AnonymousWheel struct {
	Circle
	Spokes int
}
