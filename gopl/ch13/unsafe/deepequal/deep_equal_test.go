package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	strs := strings.Split("a:b:c", ":")
	arr := []string{"a", "b", "c"}
	if !reflect.DeepEqual(strs, arr) {
		t.Error("not equal")
	}

	nilSlice := []string(nil)
	emptySlices := []string{}

	if !reflect.DeepEqual(nilSlice, emptySlices) {
		t.Error("not equal")
	}

}
