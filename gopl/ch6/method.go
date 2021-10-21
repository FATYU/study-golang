package main

import (
	"fmt"
	"study-golang/gopl/ch6/point"
)

func main() {
	p := point.Point{1, 2}
	q := point.Point{2, 4}

	fmt.Println(point.Point.Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(q.Distance(p))

	path := Path{point.Point{1, 2}, point.Point{1, 3}, point.Point{2, 2}, point.Point{2, 3}, point.Point{3, 4}}

	fmt.Println(path.Distance())

	//方法值
	distanceP := p.Distance //将一个方法赋值给一个变量

	fmt.Println(distanceP(q))

	pointDistance := point.Distance
	pointInstanceDistance := p.Distance

	fmt.Println("%T\n", pointDistance)
	fmt.Println("%T\n", pointInstanceDistance)

}

type Path []point.Point

func (p Path) Distance() float64 {
	sum := 0.0
	for pointIndex := range p {
		if pointIndex > 0 {
			sum += p[pointIndex-1].Distance(p[pointIndex])
		}
	}
	return sum

}
