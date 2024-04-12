package main

import (
	"errors"
	"fmt"
)

// to create an error simply make a call to error.New(string).
func main() {
	err := errors.New("Error created")
	fmt.Println(err)
}
