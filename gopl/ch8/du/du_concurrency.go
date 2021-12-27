package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	//创建countdown
	countdown := sync.WaitGroup{}
	fileSize := make(chan int64)
	paths := []string{"/Users/zhangyu/workspace/labs/N", "/Users/zhangyu/workspace/labs/n3", "/Users/zhangyu/workspace/labs/blockchain"}

	for _, path := range paths {
		countdown.Add(1) //计数器+1
		go concurrencyRecursDirAndFileSize(path, &countdown, fileSize)
	}

	//等待并发递归结束
	go func() {
		countdown.Wait()
		close(fileSize)
	}()

	var fileCount, fileByteSize int64
	for size := range fileSize {
		fileCount++
		fileByteSize += size
	}

	printInfo(fileCount, fileByteSize)
}

func printInfo(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func concurrencyRecursDirAndFileSize(dir string, n *sync.WaitGroup, size chan<- int64) {
	defer n.Done() //函数结束后，进行计数器-1
	entries := concurrencyDirEntries(dir)
	for _, fileInfo := range entries {
		if fileInfo.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, fileInfo.Name())       //拼接子路径
			go concurrencyRecursDirAndFileSize(subDir, n, size) // 并发计算，注意goroutine数量会创建很多
		}
		size <- fileInfo.Size()
	}
}

func concurrencyDirEntries(dir string) []fs.FileInfo {
	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : ", err)
		return nil
	}
	return readDir
}
