package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollar float32

func (d dollar) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollar

func (d database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// for item, price := range d {
	// 	fmt.Fprintf(w, "%s\n", item, price)
	// }

	switch req.URL.Path {
	case "/list":
		for item, price := range d {
			fmt.Fprintf(w, "%s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := d[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item : %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}

	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
