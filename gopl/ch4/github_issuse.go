package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
	User      *User
}

type User struct {
	HTMLURL string
	Login   string
}

func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		fmt.Errorf("search Issues error : %s", err)
	}
	if result == nil {
		return
	}
	fmt.Printf("issuse result count is : %d\n", result.TotalCount)

	for _, data := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", data.Number, data.User.Login, data.Title)
	}

}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("http response error is : %s", resp.StatusCode)
	}
	// 正确数据
	var result IssuesSearchResult
	//基于流式的解码器json.Decoder，它可以从一个输入流解码JSON数据
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, err
}
