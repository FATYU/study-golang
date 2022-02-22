package word_bugfix

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	// 忽略大小写
	s = strings.ToLower(s)

	var runes []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			runes = append(runes, r)
		}
	}

	for i := range runes {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func SoPut(s string) {
	fmt.Println("0000::::::::::::::::::::0")
}
