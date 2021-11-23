package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// 使用 http 包中的 handler 来自定义一个从数据库中获取相应的数据映射的handler
func main() {
	db := database{"shoes": 50, "socks": 5}
	// 使用NewServerMux 进行单一路径的单一处理函数
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	//mux.Handle("/list", http.HandlerFunc(db.price)) // runtime error: http: multiple registrations for /list

	// 使用默认的httpserver
	//log.Fatal(http.ListenAndServe("localhost:8080", db)) //使用db默认的ServerHTTP函数
	// 使用ServerMux绑定httpserver
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%0.2f", d)
}

//
type database map[string]dollars

//database 实现http的ServerHTTP接口
func (db database) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("-                   http server handler process                           -")
	//根据url地址进行判断，然后返回不同的处理结果

	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(writer, "%s: %s\n", item, price)
		}
	case "/query":
		item := req.URL.Query().Get("item")
		if item != "" {
			price, ok := db[item]
			if !ok {
				writer.WriteHeader(http.StatusNotFound) // 404
				fmt.Fprintf(writer, "no such item: %q\n", item)
				return
			}
			fmt.Fprintf(writer, "%s\n", price)
		}

		price := req.URL.Query().Get("price")
		if price != "" {

			priceValue, _ := strconv.ParseFloat(price, 32)

			itemName, ok := mapkey(db, dollars(priceValue))
			if !ok {
				writer.WriteHeader(http.StatusNotFound) // 404
				fmt.Fprintf(writer, "no such price: %q\n", item)
				return
			}
			fmt.Fprintf(writer, "%s\n", itemName)
		}

	default:
		writer.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(writer, "no such page: %s\n", req.URL)
	}
}

func mapkey(m map[string]dollars, value dollars) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
