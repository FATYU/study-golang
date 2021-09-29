package main

import "fmt"

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

func main() {

	var movies = []Movie{
		{Title: "1", Year: 1988, Color: false, Actors: []string{"BBB", "AAA"}},
		{Title: "2", Year: 1989, Color: false, Actors: []string{"BBB", "AAA"}},
		{Title: "3", Year: 1990, Color: false, Actors: []string{"BBB", "AAA"}},
	}

	fmt.Printf("%#v\n", movies)

}
