package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	mapset "github.com/deckarep/golang-set"
)

func main() {
	d := downloadedCourses()
	all := courseSet()

	diff := all.Difference(d).Iterator()
	for value := range diff.C {
		fmt.Println(value)
	}
}

func downloadedCourses() mapset.Set {
	courses := mapset.NewSet()
	dirs, err := os.ReadDir("/Users/zhangyu/geektime-downloader/18757155952")
	if err != nil {
		panic(err)
	}
	for _, file := range dirs {
		// fmt.Println(file.Name())
		courses.Add(file.Name())
	}
	return courses
}
func courseSet() mapset.Set {
	courses := mapset.NewSet()
	filePath := "/Users/zhangyu/Desktop/1.html"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("读取文件失败", err)
		panic(err)
	}
	defer file.Close() // 关闭文件流

	reader := bufio.NewReader(file)
	html := ""
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		html += str
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("解析字符串到docuemnt失败")
		panic(err)
	}
	doc.Find("div.course-title").Each(func(i int, selection *goquery.Selection) {
		ss := strings.Replace(selection.Text(), " ", "", -1)
		courses.Add(ss)
	})

	return courses
}
