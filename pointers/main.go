package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var ptr *int = &x
	// anotherPtr := &x
	fmt.Println(x)
	fmt.Println(ptr)
	// fmt.Println(*ptr)
	// fmt.Println(anotherPtr)

	arr := [3]int{1, 2, 3}

	arrPtr := &arr[0]

	fmt.Println(*arrPtr)
	// arrPtr++
	// arrPtr = (*int)(unsafe.Pointer(
	// 	uintptr(unsafe.Pointer(ptr))
	// 	 + uintptr(unsafe.Sizeof(arr[0])))
	// 	 )

	fmt.Println(*arrPtr)
}
