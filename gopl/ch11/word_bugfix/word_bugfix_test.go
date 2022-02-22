package word_bugfix

import (
	"fmt"
	"testing"
)

//「表格」测试 准备响应的结构体和测试数据，进行批量测试
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

//覆盖率测试
func TestCoverage(t *testing.T) {
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

//基准测试
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("zhangyuzhangyuzhangyuyu")
	}
}

// 示例测试
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("zzzzzzzzzzzzzzzz"))

	//Output:
	//true
}
