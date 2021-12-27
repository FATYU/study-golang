package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {

	fmt.Println("-------------------------------------------------------", time.Now().Format("2016-01-02 15:04:05"))

	fileSize := make(chan int64)

	paths := []string{"/Users/zhangyu/workspace/labs/N", "/Users/zhangyu/workspace/labs/n3", "/Users/zhangyu/workspace/labs/blockchain"}

	go func() { //开启一个goroutine递归统计
		for _, path := range paths {
			recursDirAndFileSize(path, fileSize)
		}
		close(fileSize)
	}()

	var fileCount, fileByteSize int64
	for size := range fileSize {
		fileCount++
		fileByteSize += size
	}

	printDiskUsage(fileCount, fileByteSize)
	fmt.Println("-------------------------------------------------------", time.Now().Format("2016-01-02 15:04:05"))
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func recursDirAndFileSize(dir string, size chan<- int64) {
	entries := dirEntries(dir)
	for _, fileInfo := range entries {
		if fileInfo.IsDir() {
			subDir := filepath.Join(dir, fileInfo.Name()) //拼接子路径
			recursDirAndFileSize(subDir, size)
		}
		size <- fileInfo.Size()
	}
}

func dirEntries(dir string) []fs.FileInfo {
	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : ", err)
		return nil
	}
	return readDir
}
