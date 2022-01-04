package main

import (
	"sync"
	"time"
)

type RW interface {
	Read()
	Write()
}

const cost = time.Millisecond * 100 //暂停时间

type Lock struct {
	count int
	mutex sync.Mutex
}

func (l *Lock) Read() {
	l.mutex.Lock()
	_ = l.count //读取count数据
	time.Sleep(cost)
	l.mutex.Unlock()
}

func (l *Lock) Write() {
	l.mutex.Lock()
	l.count++ //进行count计数处理
	time.Sleep(cost)
	l.mutex.Unlock()
}

type RWLock struct {
	count int
	mutex sync.RWMutex
}

func (l *RWLock) Read() {
	l.mutex.RLock()
	_ = l.count
	time.Sleep(cost)
	l.mutex.RUnlock()
}

func (l *RWLock) Write() {
	l.mutex.Lock()
	l.count++
	time.Sleep(cost)
	l.mutex.Unlock()
}
