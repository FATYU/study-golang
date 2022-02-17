package multichannel

/**
使用channel方式进行并发处理
*/

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//调用函数
func (entry *entry) call(f Func, key string) {

	//执行耗时函数
	entry.res.value, entry.res.err = f(key)

	close(entry.ready) //就绪后，关闭channel

}

//广播结果
func (entry *entry) deliver(result chan<- result) {
	<-entry.ready       //判断是否就绪，执行完成，close（）函数调用
	result <- entry.res //向通道发送结果数据

}

type result struct {
	value interface{}
	err   error
}

type Func func(key string) (interface{}, error)

//  对url请求进行并发处理 类似阻塞队列，接收要处理的url，并发处理，返回结果
//  call函数 调用http get请求，并返回结果
//  deliver进行结果数据的分发，通过channel提交到server
//  最终进行channel的关闭

// 定义结构体

// 1. request

type request struct {
	key      string        // url地址
	response chan<- result //定义channel ，接收http的返回结果
}

// 2. Memo cache
type Memo struct {
	requests chan request //缓存，请求的地址以及对应的响应结果
}

// 构造函数

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.handle(f) //异步启动服务端，并指定处理函数
	return memo
}

// 实现接口Get函数， 程序异步的入口,
// 可以理解为将数据添加到队列中， 等待被处理

func (memo *Memo) Get(key string) (interface{}, error) {
	// 获取结果后将数据推送到channel中
	response := make(chan result)
	memo.requests <- request{key: key, response: response} //将 request 推送到 channel 中
	res := <-response                                      //从channel中获取结果
	return res.value, res.err
}

//

func (memo *Memo) handle(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		c := cache[req.key]
		if c == nil {
			c = &entry{ready: make(chan struct{})}
			cache[req.key] = c
			go c.call(f, req.key) //异步调用函数
		}
		go c.deliver(req.response) // 启动deliver，进行数据处理
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}
