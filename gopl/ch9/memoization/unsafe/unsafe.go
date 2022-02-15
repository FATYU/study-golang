//提供并发安全的缓存机制

package unsafe

//定义一个缓存结构体，缓存函数和其返回值

type Memo struct {
	f     Func              //缓存的函数
	cache map[string]result //缓存的结果
}

type Func func(string) (interface{}, error) //函数类型，接受一个字符串参数，返回一个结果和错误

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// 与Memo结构体绑定的函数，实现缓存的调用,注意这个方式不是并发安全的

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok { //判断cache中是否存在
		res.value, res.err = memo.f(key) //如果不存在，执行对应的方法，获取方法执行的结果，并缓存对应的数据
		memo.cache[key] = res
	}
	return res.value, res.err
}
