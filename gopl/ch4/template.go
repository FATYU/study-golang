package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

const templateData = `
{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}
`

func main() {

	//生成模版需要执行两步:
	//1.分析模版结构转换为内部表示
	//2.根据指定的输入,执行模版
	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templateData)
	if err != nil {
		fmt.Printf("error:%s", err)
	}
	report.Execute(os.Stdout, "")
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
