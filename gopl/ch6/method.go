package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 2}
	q := Point{2, 4}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(q.Distance(p))

	path := Path{{1, 2}, {1, 3}, {2, 2}, {2, 3}, {3, 4}}

	fmt.Println(path.Distance())

}

type Point struct {
	X, Y float64
}

/**
 * 第一个Distance的调用实际上用的是包级别的函数geometry.Distance
 * 第二个则是使用刚刚声明的Point，调用的是Point类下声明的Point.Distance方法
 */

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)

}

type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for pointIndex := range p {
		if pointIndex > 0 {
			sum += p[pointIndex-1].Distance(p[pointIndex])
		}
	}
	return sum

}
