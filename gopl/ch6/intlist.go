package main

import "fmt"

func main() {

	var tail = IntList{3, nil}
	var mid = IntList{2, &tail}
	var il = IntList{1, &mid}
	fmt.Println(il.Sum())

}

type IntList struct {
	Value int
	Tail  *IntList
}

func (il *IntList) Sum() int {
	if nil == il {
		return 0
	}
	return il.Value + il.Tail.Sum()
}
