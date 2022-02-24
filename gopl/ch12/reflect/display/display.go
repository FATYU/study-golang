package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func format(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Slice, reflect.Map, reflect.Ptr, reflect.Func:
		return v.Type().String() + " 's value is " + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " 's value is " + v.String()
	}
}

func Display(name string, v interface{}) {
	fmt.Printf("Display name :%s, type: (%T),value  :\n", name, v)
	displayValue(name, reflect.ValueOf(v))

}

func displayValue(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = Invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			displayValue(fmt.Sprintf("%s[%d]", path, i), v.Index(i)) //递归调用displayValue方法
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldName := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			displayValue(fieldName, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			displayValue(fmt.Sprintf("%s[%s]", path, format(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s == nil\n", path)
		} else {
			displayValue(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s == nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			displayValue(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, format(v))
	}

}
