package main

import (
	"bytes"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	//strs := strings.Split("a:b:c", ":")
	//arr := []string{"a", "b", "c"}
	//if !DeepEqual(strs, arr) {
	//	t.Error("not equal")
	//}

	///  准备数据

	var one, againOne, two = 1, 1, 2
	//
	//type CyclePtr *CyclePtr //循环指针
	//var cyclePtr1, cyclePtr2 CyclePtr
	//
	//cyclePtr1 = &cyclePtr1
	//cyclePtr2 = &cyclePtr2
	//
	//type CycleSlice []CycleSlice
	//var cycleSlice = make(CycleSlice, 1)
	//cycleSlice[0] = cycleSlice
	//
	var ch1, ch2 = make(chan int), make(chan int)
	var ch1ro <-chan int = ch1
	//
	type mystring string

	var interface1, interfaceAgain1, interface2 interface{} = &one, &againOne, &two

	//创建 for 循环测试数据集
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{ // basic types
		{1, 1, true},
		{1, 2, false},   // different values
		{1, 1.0, false}, // different types
		{"foo", "foo", true},
		{"foo", "bar", false},
		{mystring("foo"), "foo", false}, // different types
		//// slices
		{[]string{"foo"}, []string{"foo"}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string{}, []string(nil), true},
		// slice cycles
		//{cycleSlice, cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3}},
			true,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3, 4}},
			false,
		},
		{
			map[string][]int{},
			map[string][]int(nil),
			true,
		},
		// pointers
		{&one, &one, true},
		{&one, &two, false},
		{&one, &againOne, true},
		{new(bytes.Buffer), new(bytes.Buffer), true},
		//// pointer cycles
		//{cyclePtr1, cyclePtr1, true},
		//{cyclePtr2, cyclePtr2, true},
		//{cyclePtr1, cyclePtr2, true}, // they're deeply equal
		// functions
		{(func())(nil), (func())(nil), true},
		{(func())(nil), func() {}, false},
		{func() {}, func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}, true},
		{[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
		// channels
		{ch1, ch1, true},
		{ch1, ch2, false},
		{ch1ro, ch1, false}, // NOTE: not equal
		// interfaces
		{&interface1, &interface1, true},
		{&interface1, &interface2, false},
		{&interfaceAgain1, &interface1, true},
	} {
		if DeepEqual(test.x, test.y) != test.want {
			t.Errorf("DeepEqual(%v, %v) = %t", test.x, test.y, !test.want)
		}
	}

	//遍历执行测试

}
