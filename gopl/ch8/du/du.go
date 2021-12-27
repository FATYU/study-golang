package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	RecursDir("/Users/zhangyu/workspace/labs/N")
}
func RecursDir(dir string) {

	entries := DirEntries(dir)

	for _, fileInfo := range entries {
		if fileInfo.IsDir() {
			subDir := filepath.Join(dir, fileInfo.Name()) //拼接子路径
			RecursDir(subDir)
		}
		fmt.Print(fileInfo.Size())

		fmt.Println("\t" + dir + "/" + fileInfo.Name())
	}

}

func DirEntries(dir string) []fs.FileInfo {

	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : ", err)
		return nil
	}
	return readDir
}
