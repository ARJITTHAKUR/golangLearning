package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Define the backend servers
	server1 := "http://localhost:8081"
	server2 := "http://localhost:8082"

	// Parse the backend server URLs
	url1, err := url.Parse(server1)
	if err != nil {
		panic(err)
	}

	url2, err := url.Parse(server2)
	if err != nil {
		panic(err)
	}

	// Create a reverse proxy for each backend server
	proxy1 := httputil.NewSingleHostReverseProxy(url1)
	proxy2 := httputil.NewSingleHostReverseProxy(url2)

	// Define a handler function that will choose the backend server based on some criteria
	handler := func(w http.ResponseWriter, r *http.Request) {
		// You can add your own logic here to determine which backend server to use
		// For example, based on the request path, headers, or any other criteria

		// In this example, we're using a simple round-robin strategy
		if r.URL.Path == "/server1" {
			proxy1.ServeHTTP(w, r)
		} else {
			proxy2.ServeHTTP(w, r)
		}
	}

	// Create a new HTTP server with the custom handler
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	// Start the server
	fmt.Println("Reverse proxy server listening on :8080")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
