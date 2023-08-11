package main

import "fmt"

func newMap(arr []int, fn func(a int, i int, ar []int) int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = fn(arr[i], i, arr)
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4}
	newArr := newMap(arr, func(e int, i int, ar []int) int {
		return e * 2
	})
	fmt.Println(newArr)
}
