package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

}

//有向无环图
func toposort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool) //是否已经存在

	var visitAll func(items []string) //函数当作类型

	visitAll = func(items []string) {
		for _, item := range items { //遍历所有的元素
			if !seen[item] {
				seen[item] = true
				visitAll(m[item]) //获取当前元素依赖的所有元素
				order = append(order, item)
			}
		}
	}

	var keys []string

	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	visitAll(keys)
	return order
}
