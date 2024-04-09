package main

import "fmt"

// A function is a small portion of code that surrounds some action you want to perform and
// return one or more values (or nothing)
// This is the main tool for we developers to maintain structure, encapsulation and code readability

func hello(message string) error {
	fmt.Printf("message: %v\n", message)
	return nil
}

func main() {
	hello("Shivam Singh")
}
