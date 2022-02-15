package safe

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

// 使用mutex进行加锁，但是降低了并发性能

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mutex.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		if res.err != nil {
			log.Print(res.err)
		}
		memo.cache[key] = res
	}

	memo.mutex.Unlock()
	return res.value, res.err

}
