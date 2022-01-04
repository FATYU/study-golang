package main

import (
	"sync"
	"testing"
)

// run with ` go test -bench . `
func benchmark(tb *testing.B, rw RW, read, write int) {
	for i := 0; i < tb.N; i++ {
		var wait = sync.WaitGroup{}
		for k := 0; k < read*100; k++ {
			wait.Add(1)
			go func() {
				rw.Read()
				wait.Done()
			}()

		}

		for k := 0; k < write*100; k++ {
			wait.Add(1)
			go func() {
				rw.Write()
				wait.Done()
			}()
		}
		wait.Wait()
	}

}

func BenchmarkRWReadMore(b *testing.B) {
	benchmark(b, &RWLock{}, 9, 1)
}

func BenchmarkRWWriteMore(b *testing.B) {
	benchmark(b, &RWLock{}, 1, 9)
}

func BenchmarkRW(b *testing.B) {
	benchmark(b, &RWLock{}, 5, 5)
}

func BenchmarkReadMore(b *testing.B) {
	benchmark(b, &Lock{}, 9, 1)
}

func BenchmarkWriteMore(b *testing.B) {
	benchmark(b, &Lock{}, 1, 9)
}

func Benchmark(b *testing.B) {
	benchmark(b, &Lock{}, 5, 5)
}
