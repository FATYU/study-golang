package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	links := "/Users/zhangyu/Desktop/links.txt"

	file, err := os.OpenFile(links, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("File open error!", err)
		return
	}

	defer file.Close()

	//按行读取：
	start, err := file.Stat()

	if err != nil {
		fmt.Println("read start error!", err)
		panic(err)
	}

	var size = start.Size()
	fmt.Println("file bytes :", size)

	buf := bufio.NewReader(file)

	var i = 1
	var split = ""
	for {
		if i%20 == 1 {

			istr := strconv.Itoa(i)
			split = strings.Join([]string{"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<opml version=\"2.0\">\n<head>\n<title>\U0001F9E9 Feeds</title>\n</head>\n<body><outline text=\"独立博客", istr, "\\\" title=\"独立博客", istr, "\\\">"}, "")
		}
		line, err := buf.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				split = strings.Join([]string{split, "\n</outline>\n</body>\n</opml>"}, "")
				mkFile(i, split)
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		line = strings.TrimSpace(line)
		//fmt.Println(line)

		split = strings.Join([]string{split, line}, "")

		if i%20 == 0 {

			// writer to file
			split = strings.Join([]string{split, "\n</outline>\n</body>\n</opml>"}, "")
			mkFile(i, split)

		}
		i = i + 1

	}

}

func mkFile(i int, data string) {
	var f *os.File
	fileName := strconv.Itoa(i)
	fileName = "/Users/zhangyu/Desktop/opml/独立博客" + fileName + ".opml"
	f, err1 := os.Create(fileName) //创建文件

	if err1 != nil {
		panic(err1)
	}

	defer f.Close()

	io.WriteString(f, data)

}
