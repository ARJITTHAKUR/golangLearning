package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.DefaultServeMux
	server := &http.Server{
		Addr:              "localhost:8080",
		Handler:           mux,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})

	server.ListenAndServe()
}
