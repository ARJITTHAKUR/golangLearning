package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	server := http.Server{
		Addr: port,
	}

	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hey there<h1/>"))
	})
	serverMux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hey New<h1/>"))
	})
	server.Handler = serverMux
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Running on port : %s", port))
}
