package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	image := make(map[int]int)
	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(3)
		go func(i int, mu *sync.Mutex) {
			mu.Lock()
			// fmt.Println("write", i)
			image[i] = rand.Intn(200)
			defer mu.Unlock()
			wg.Done()
		}(i, &mu)

		go func(i int, mu *sync.Mutex) {
			mu.Lock()
			// fmt.Println("read", i)
			fmt.Printf("data ===> %d : %d\n", i, image[i])
			defer mu.Unlock()
			wg.Done()

		}(i, &mu)

		go func(i int, mu *sync.Mutex) {
			mu.Lock()
			// fmt.Println("write", i)
			image[i] = rand.Intn(100)
			defer mu.Unlock()
			wg.Done()

		}(i, &mu)
	}
	wg.Wait()
	// time.Sleep(1 * time.Second)
	fmt.Printf("%+v", image)
}
