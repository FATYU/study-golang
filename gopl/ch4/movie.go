package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int
	Color  bool `json:"json-color,omitempty"` // 当类型的对象被 json 格式化的时候, 此字段使用「json-color」最为 key
	Actors []string
}

func main() {

	var movies = []Movie{
		{Title: "1", Year: 1988, Color: false, Actors: []string{"BBB", "AAA"}},
		{Title: "2", Year: 1989, Color: true, Actors: []string{"BBB", "AAA"}},
		{Title: "3", Year: 1990, Color: false, Actors: []string{"BBB", "AAA"}},
	}

	fmt.Printf("%#v\n", movies)

	data, err := json.Marshal(movies) // json 编组(marchalling)
	if err != nil {
		log.Fatalf("Json marchalling error : %s\n", err)
	}
	fmt.Printf("%s\n", data)                              //打印格式比较紧凑,可以使用json.MarshalIndent来代替json.Marshal方法
	data1, err1 := json.MarshalIndent(movies, "", "    ") // json 编组(marchalling)
	if err1 != nil {
		log.Fatalf("Json marchalling error : %s\n", err1)
	}
	fmt.Printf("%s\n", data1) //打印格式比较紧凑,可以使用json.MarshalIndent来代替json.Marshal方法
}
