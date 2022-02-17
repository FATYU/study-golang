package channel

import "sync"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	mutex sync.Mutex // guards cache
	cache map[string]*entry
}

type Func func(string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

//使用多个 goroutine 进行数据的获取和缓存的管理

func (memo *Memo) Get(key string) (interface{}, error) {
	// 首先进行加锁 尝试获取cache
	memo.mutex.Lock()
	e := memo.cache[key]
	if e == nil {
		//缓存不存在 通过entry进行数据保存，相当于加了个异步的共享内存，可以并发进行数据获取
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mutex.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		//已经存在缓存， 进行解锁
		memo.mutex.Unlock()
		// 等待缓存的值
		<-e.ready
	}
	return e.res.value, e.res.err
}
