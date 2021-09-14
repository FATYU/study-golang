package main

import "fmt"

func main() {

	fmt.Printf("the result is %v\n", filterEmpty([]string{"", "1", "", "2", "3"}))

	// filterEmpty 方法存在一个问题
	data := []string{"1", "", "3"}
	fmt.Printf("filter result is %v\n", filterEmpty(data)) //[1,3]
	fmt.Printf("data is %v\n", data)                       //[1,3,3]
	//解决上上面的问题 使用 data = filterEmpty(data)

}

func filterEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	//比较微妙的地方是，输入的slice和输出的slice共享一个底层数组
	return strings[:i]
}
