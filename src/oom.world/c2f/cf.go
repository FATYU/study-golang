package c2f

//注意,即使是两个内存类型相同的type ,不可以用==、>=、 <=等操作符号进行比较
type C float64
type F float64

const (
	AbsoluteZeroC C = -273.15
	FreezingC     C = 0
	BoilingC      C = 100
)

func f2C(f float64) C {
	return C((f - 32) * 5 / 9)
}

func F2C(f float64) C {
	return C((f - 32) * 5 / 9)
}
func c2F(c float64) F {

	return F(c*9/5 + 32)
}
