package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

}

func findlink(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("http get error :" + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println("http get result :", resp.StatusCode)
	}

	// html.Parse(resp.Body)

	// return visit(body), nil
	return nil, nil
}

func visit(html string) []string {

	return nil
}
