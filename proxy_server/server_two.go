package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	fmt.Println("starting server two")
	http.HandleFunc("/two", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("server two responding")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}
