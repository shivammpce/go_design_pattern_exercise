package main

import "fmt"

// An anonymous function is a function without a name. This is useful when we want to return
// a function from another function that doesn't need a context or when we want to pass a function
// to a different function

func main() {
	add := func(m int) int {
		return m + 1
	}
	result := add(6)
	fmt.Printf("result: %v\n", result)
}
