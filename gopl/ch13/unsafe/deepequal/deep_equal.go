package main

import (
	"reflect"
	"unsafe"
)

func DeepEqual(x, y interface{}) bool {
	seen := make(map[Comparison]bool)
	return deepEqual(reflect.ValueOf(x), reflect.ValueOf(y), seen)

}

func deepEqual(x, y reflect.Value, seen map[Comparison]bool) bool {
	//判断x，y是否有效

	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	if x.Type() != y.Type() {
		return false
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return x.Int() == y.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint()
	case reflect.Float32, reflect.Float64:
		return x.Float() == y.Float()
	case reflect.Complex64, reflect.Complex128:
		return x.Complex() == y.Complex()
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return deepEqual(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !deepEqual(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !deepEqual(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !deepEqual(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true

	}

	panic("未正常处理")
}

type Comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
