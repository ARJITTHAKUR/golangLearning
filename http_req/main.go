package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	authRequest, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	authRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "token"))
	authRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// make request here
	res, err := client.Do(authRequest)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%+v", body)
}
