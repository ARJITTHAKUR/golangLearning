package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	// fmt.Println("starting server 1")
	http.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("server one responding")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
