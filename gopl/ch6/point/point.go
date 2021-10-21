package point

import "math"

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
