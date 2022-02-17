package doublelock

import (
	"log"
	"sync"
)

type Memo struct {
	f     Func
	mutex sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	memo := &Memo{f: f, cache: make(map[string]result)}
	return memo
}

// 针对safe.go的代码，进行优化
//拆分加锁范围。
// 但是程序性能问题如果函数执行了相同的输入，在并发情况下，可能出现重复执行f()的情况。
//为了避免这种情况出现，（此方案叫做「重复抑制」Duplicate Suppression）可以通过channel进行消息广播， 从而避免重复执行f() 参照channel文件夹

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mutex.Lock()
	res, ok := memo.cache[key]
	memo.mutex.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		if res.err != nil {
			log.Print(res.err)
		}
		memo.mutex.Lock()
		memo.cache[key] = res
		memo.mutex.Unlock()
	}
	return res.value, res.err

}
