package main

import "fmt"

type Values map[string][]string

func (v Values) Get(key string) string {

	if len(v[key]) > 0 {
		return v[key][0]
	}
	return "[no values]"
}

func (v Values) Add(key string, vals string) {
	v[key] = append(v[key], vals)

}

func main() {
	v := Values{}
	v.Add("z", "z")
	v.Add("y", "y")

	fmt.Println("use type get method,value is ", v.Get("z"))
	fmt.Println("Direct map access,value is ", v["z"][0])

	v = nil
	fmt.Println(v.Get("z"))
	//fmt.Println(nil.Get("z"))

}
