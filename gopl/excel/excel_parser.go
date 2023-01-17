package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

func main() {
	f, err := excelize.OpenFile("/Users/zhangyu/Desktop/话术.xlsx")
	// 读取文件错误
	if err != nil {
		fmt.Println(err)
		return
	}
	// 清理 IO
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	rows, err := f.GetRows("自定义话术")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		huashu := ""
		fileurl := ""
		for index, column := range row {
			if index == 3 {
				huashu = column
			}
			if index == 7 {
				fileurl = column
				fileurl = "y:/video/ALvideo/Daisy/" + strings.TrimSuffix(fileurl, ".wav") + ".3gp"
			}
		}

		fmt.Printf("daisy.put(%q,%q);\n", huashu, fileurl)
	}
}
