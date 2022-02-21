package word_bugfix

import "testing"

//func TestIsPalindrome(t *testing.T) {
//	if !IsPalindrome("hi,ih") {
//		t.Errorf("IsPalindrome(%q) = false", "hi,ih")
//	}
//
//	if !IsPalindrome("été") {
//		t.Errorf("IsPalindrome(%q) = false", "été")
//	}
//}

func TestIsPalindrome(t *testing.T) {
	var data = []struct {
		input string
		want  bool
	}{
		{"zhangyu", false},
		{"hello", false},
		{"122221", true},
		{"hi,ih", true},
		{"zzzzzz", true},
	}

	for _, test := range data {
		if result := IsPalindrome(test.input); result != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, result)
		}
	}

}
