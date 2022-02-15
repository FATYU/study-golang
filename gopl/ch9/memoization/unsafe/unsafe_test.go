package unsafe

import (
	"study-golang/gopl/ch9/memoization/test"
	"testing"
)

var httpGetBody = test.HttpGetBody

func Test(t *testing.T) {
	m := New(httpGetBody) //自身包下的测试，不需要进行 import 操作
	test.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	test.Concurrent(t, m)
}
