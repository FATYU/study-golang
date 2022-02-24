package main

import (
	"testing"
)

//func Example_int() {
//	Display("int 3", 3)
//	//Output:
//	//Display name :int 3, type: (int),value  :
//	//int 3 = 3
//}

func TestDisplay(t *testing.T) {
	Display("int 3", 3)

	//Output:
	//Display name :int 3, type: (int),value  :
	//int 3 = 3
}

func TestDisplayStruct(t *testing.T) {
	type Movie struct {
		Title    string
		Subtitle string
		Year     int
		Color    bool
		Actor    map[string]string
		Oscars   []string
		Sequel   *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", strangelove)

}
