package main

import (
	"strings"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	strs := strings.Split("a:b:c", ":")
	arr := []string{"a", "b", "c"}
	if !DeepEqual(strs, arr) {
		t.Error("not equal")
	}

}
