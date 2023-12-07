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
	for item, price := range d {
		fmt.Fprintf(w, "%s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}

	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
