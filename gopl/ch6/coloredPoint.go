package main

import (
	"fmt"
	color "image/color"
	point "study-golang/gopl/ch6/point"
)

type ColoredPoint struct {
	point.Point
	Color color.RGBA
}

func main() {

	var pointX = ColoredPoint{point.Point{1, 1}, color.RGBA{255, 0, 0, 255}}
	var pointY = ColoredPoint{point.Point{1, 2}, color.RGBA{1, 1, 1, 1}}

	fmt.Println(pointX)
	fmt.Println(pointY)
}
