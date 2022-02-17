package channel

import (
	"study-golang/gopl/ch9/memoization/test"
	"testing"
)

var httpGetBody = test.HttpGetBody

func TestSequential(t *testing.T) {
	m := New(httpGetBody)
	test.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	test.Concurrent(t, m)
}
