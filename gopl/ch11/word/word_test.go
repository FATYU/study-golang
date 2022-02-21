package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("zhangyu") {
		t.Error(`IsPalindrome("zhangyu") = false`)
	}

	if !IsPalindrome("zyxyz") {
		t.Error(`IsPalindrome("zyxyz") = false`)
	}
}

func TestIsNotPalindrome(t *testing.T) {
	if IsPalindrome("zyxyz") {
		t.Error(`IsPalindrome("zyxyz") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
