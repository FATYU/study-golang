package main

import (
	"bytes"
	"fmt"
)

// IntSet
// https://studygolang.com/articles/24576
// 定义int集合
// 数据示例 每行的范围是 0+64(n-1) ～（64*n-1）
// 063 062 061 '''''''''' 000
// 127 126 125 '''''''''' 064
// 191 190 189 '''''''''' 128
// ''''''''''''''''''''''''''
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value data
func (s *IntSet) Has(data int) bool {

	word, bit := data/64, uint(data%64) //计算数字的行和列

	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value data to the set
func (s *IntSet) Add(data int) {
	word, bit := data/64, uint(data%64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit

}

// UnionWith 合并集合 sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// 打印字符串
// 首尾{}拼接，中间以「 空格 」进行拼接
func (s *IntSet) String(data *IntSet) string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range data.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				//补充空格
				if buf.Len() > len("{") { //判断字符串长度已经大于 "{"
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}

	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	intSet := IntSet{}
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(64)
	intSet.Add(127)
	intSet.Add(129)
	fmt.Println(intSet.String(&intSet))

	newIntSet := IntSet{}
	newIntSet.Add(800)
	newIntSet.Add(65535)
	intSet.UnionWith(&newIntSet)
	fmt.Println((intSet.words))

	fmt.Println(intSet.String(&intSet))
	fmt.Println("the set data len is : ", len(intSet.words))
	fmt.Println("the set has a num 100? :::==  ", intSet.Has(100))
}
