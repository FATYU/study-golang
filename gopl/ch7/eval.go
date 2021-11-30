package main

// 拓展Expr进行表达式计算

// A Var identifies a variable, e.g., x.
//元素
type Var string

// A literal is a numeric constant, e.g., 3.141.
//常量
type literal float64

// A unary represents a unary operator expression, e.g., -x.
//一元操作
type unary struct {
	op rune // one of '+', '-' 例如 -x +x
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
//二元操作
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
//函数调用
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

type Env map[Var]float64

type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
}
