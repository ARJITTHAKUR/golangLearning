package main

import (
	"fmt"
)

func main() {

	var found bool
	carToLookFor := "Blazer"

	cars := []string{"Accord", "random", "Blazer"}

	for _, car := range cars {
		if car == carToLookFor {
			found = true
		}
	}

	fmt.Println("Found ? ", found)

	// functional way

	// fmt.Println("Found ? ", cars.contains("Blazer"))
}
